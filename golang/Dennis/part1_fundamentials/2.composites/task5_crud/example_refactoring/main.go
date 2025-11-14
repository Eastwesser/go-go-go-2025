package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Order struct {
	ID   string `json:"id"`
	User User   `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const (
	port         = ":8042"
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
	userService  = "http://user-service:8081/user"
	maxIDs       = 100
)

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ids := r.URL.Query().Get("ids")
	if ids == "" {
		http.Error(w, "Missing ids parameter", http.StatusBadRequest)
		return
	}

	orderIDs := strings.Split(ids, ",")
	if len(orderIDs) > maxIDs {
		http.Error(w, fmt.Sprintf("Too many IDs (max %d)", maxIDs), http.StatusBadRequest)
		return
	}

	orders := fetchOrders(orderIDs)
	if err := enrichOrdersWithUsers(orders); err != nil {
		log.Printf("Failed to enrich orders: %v", err)
		http.Error(w, "Failed to fetch user details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func fetchOrders(orderIDs []string) []Order {
	orders := make([]Order, 0, len(orderIDs))
	for _, id := range orderIDs {
		if id == "" {
			continue
		}
		orders = append(orders, Order{
			ID: id,
			User: User{
				ID: fmt.Sprintf("user-%s", id),
			},
		})
	}
	return orders
}

func enrichOrdersWithUsers(orders []Order) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(orders))
	var mu sync.Mutex

	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  false,
			DisableKeepAlives:   false,
			MaxIdleConnsPerHost: 10,
		},
	}

	for i := range orders {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			userID := orders[idx].User.ID
			user, err := fetchUserDetails(client, userID)
			if err != nil {
				errCh <- fmt.Errorf("failed to fetch user %s: %w", userID, err)
				return
			}

			mu.Lock()
			orders[idx].User = user
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	close(errCh)

	// Проверяем все ошибки
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}

func fetchUserDetails(client *http.Client, userID string) (User, error) {
	url := fmt.Sprintf("%s?id=%s", userService, userID)
	
	resp, err := client.Get(url)
	if err != nil {
		return User{}, fmt.Errorf("failed to fetch user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("user service returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, fmt.Errorf("failed to read response: %w", err)
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return user, nil
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", orderHandler)

	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("Server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

/*
    Параллельные запросы - используем goroutines и WaitGroup
    Безопасный доступ к данным - sync.Mutex для конкурентного доступа
    Обработка ошибок - правильная валидация и обработка
    Таймауты - защита от зависаний
    Лимиты - ограничение на количество ID
    Пул HTTP-клиентов - переиспользование соединений
    Правильная конкатенация - через fmt.Sprintf()
    Буферизованные слайсы - предварительное выделение памяти
    Stream JSON encoding - вместо Marshal + Write
    Proper error handling - никаких игнорирований ошибок
*/
