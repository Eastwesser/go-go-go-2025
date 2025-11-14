# 📦 Установка Go на Arch Linux

Подробное руководство по установке Go (включая последние версии) на Arch Linux.

## 🚀 Метод 1: Из официальных репозиториев Arch (Рекомендуется)

Самый простой способ:

```bash
# Обновляем систему
sudo pacman -Syu

# Устанавливаем Go
sudo pacman -S go

# Проверяем версию
go version
```

**Вывод должен показать что-то вроде:**
```
go version go1.23.x linux/amd64
```

### Настройка окружения (опционально)

Добавь в `~/.bashrc` или `~/.zshrc`:

```bash
# Go paths
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Затем перезагрузи shell:
```bash
source ~/.bashrc  # или source ~/.zshrc
```

---

## 🎯 Метод 2: Установка конкретной версии (например, Go 1.24)

### Через официальный бинарник

```bash
# 1. Скачай нужную версию с официального сайта
# Проверь актуальную версию на https://go.dev/dl/
wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz

# 2. Удали старую версию (если установлена)
sudo rm -rf /usr/local/go

# 3. Распакуй новую версию
sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz

# 4. Добавь в PATH (в ~/.bashrc или ~/.zshrc)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

# 5. Перезагрузи shell
source ~/.bashrc

# 6. Проверь версию
go version

# 7. Удали архив (опционально)
rm go1.24.0.linux-amd64.tar.gz
```

---

## 🔧 Метод 3: Через AUR (Arch User Repository)

### Установка yay (если ещё нет)

```bash
# Установи base-devel и git
sudo pacman -S --needed base-devel git

# Клонируй yay
git clone https://aur.archlinux.org/yay.git
cd yay

# Собери и установи
makepkg -si

# Вернись назад
cd ..
rm -rf yay
```

### Установка Go через yay

```bash
# Поиск доступных версий Go
yay -Ss go | grep "^aur/go"

# Установка Go
yay -S go

# Или конкретной версии (если есть в AUR)
yay -S go-1.24
```

---

## 🎮 Метод 4: goenv (Менеджер версий Go)

Если нужно переключаться между разными версиями Go:

```bash
# 1. Установи goenv
git clone https://github.com/go-nv/goenv.git ~/.goenv

# 2. Добавь в ~/.bashrc или ~/.zshrc
cat >> ~/.bashrc << 'EOF'

# goenv
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
eval "$(goenv init -)"
export PATH="$GOROOT/bin:$PATH"
export PATH="$GOPATH/bin:$PATH"
EOF

# 3. Перезагрузи shell
source ~/.bashrc

# 4. Посмотри доступные версии
goenv install -l

# 5. Установи нужную версию
goenv install 1.24.0

# 6. Установи глобально
goenv global 1.24.0

# 7. Проверь
go version
```

### Переключение между версиями с goenv

```bash
# Установить несколько версий
goenv install 1.22.0
goenv install 1.23.0
goenv install 1.24.0

# Посмотреть установленные
goenv versions

# Переключиться глобально
goenv global 1.24.0

# Переключиться локально (для текущей директории)
cd /path/to/project
goenv local 1.22.0

# Проверить текущую версию
goenv version
```

---

## ✅ Проверка установки

После установки любым методом, проверь:

```bash
# Версия Go
go version

# Переменные окружения
go env

# Создай тестовую программу
mkdir -p ~/test-go
cd ~/test-go

# Создай go.mod
go mod init test

# Создай main.go
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello from Go!")
}
EOF

# Запусти
go run main.go

# Должно вывести: Hello from Go!
```

---

## 🎯 Рекомендации для проекта Рейна Трекер

### Минимальные требования

- **Go 1.21+** (для generics и других фич)
- Go 1.23+ рекомендуется

### Проверка совместимости

```bash
cd /home/denismatveev/Desktop/treyna/reyna-train-tracker

# Проверить go.mod
cat go.mod

# Должно быть:
# module reyna-train-tracker
# go 1.21

# Обновить зависимости
go mod tidy

# Запустить проект
go run cmd/main.go
```

---

## 🐛 Решение проблем

### Проблема: "go: command not found"

**Решение:**
```bash
# Проверь PATH
echo $PATH

# Добавь Go в PATH
export PATH=$PATH:/usr/local/go/bin

# Или для постоянного эффекта добавь в ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Проблема: "permission denied" при установке

**Решение:**
```bash
# Используй sudo для системных директорий
sudo tar -C /usr/local -xzf go*.tar.gz

# Или устанавливай в home directory
tar -C $HOME -xzf go*.tar.gz
export PATH=$PATH:$HOME/go/bin
```

### Проблема: Конфликт версий

**Решение:**
```bash
# Удали все версии Go
sudo rm -rf /usr/local/go
sudo pacman -R go

# Выбери один метод установки и используй его
```

### Проблема: "cannot find package"

**Решение:**
```bash
# Обнови модули
go mod tidy

# Скачай зависимости
go mod download

# Очисти кэш (если нужно)
go clean -modcache
```

---

## 📝 Полезные команды

```bash
# Информация о Go окружении
go env

# Версия Go
go version

# Список установленных пакетов
go list ...

# Форматирование кода
go fmt ./...

# Проверка кода
go vet ./...

# Тестирование
go test ./...

# Сборка
go build

# Установка в $GOPATH/bin
go install

# Очистка
go clean

# Обновление зависимостей
go get -u ./...

# Показать зависимости
go mod graph

# Проверка зависимостей
go mod verify
```

---

## 🔄 Обновление Go

### Через pacman

```bash
sudo pacman -Syu go
```

### Через официальный бинарник

```bash
# Скачай новую версию
wget https://go.dev/dl/go1.XX.X.linux-amd64.tar.gz

# Удали старую
sudo rm -rf /usr/local/go

# Установи новую
sudo tar -C /usr/local -xzf go1.XX.X.linux-amd64.tar.gz

# Проверь
go version
```

### Через goenv

```bash
# Посмотри новые версии
goenv install -l | tail -10

# Установи
goenv install 1.XX.X

# Переключись
goenv global 1.XX.X
```

---

## 🎓 Дополнительные инструменты

### Полезные Go инструменты

```bash
# gopls (Language Server)
go install golang.org/x/tools/gopls@latest

# goimports (автоимпорты)
go install golang.org/x/tools/cmd/goimports@latest

# golangci-lint (линтер)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# delve (отладчик)
go install github.com/go-delve/delve/cmd/dlv@latest

# air (hot reload)
go install github.com/cosmtrek/air@latest
```

---

## 📚 Ресурсы

- **Официальный сайт Go**: https://go.dev/
- **Документация**: https://go.dev/doc/
- **Arch Wiki - Go**: https://wiki.archlinux.org/title/Go
- **Go Playground**: https://go.dev/play/
- **Go Tour**: https://go.dev/tour/

---

## ✨ Для проекта Рейна Трекер

После установки Go:

```bash
# Перейди в проект
cd /home/denismatveev/Desktop/treyna/reyna-train-tracker

# Проверь зависимости
go mod tidy

# Запусти
go run cmd/main.go

# Или скомпилируй
go build -o reyna-tracker cmd/main.go
./reyna-tracker
```

---

**Готово! Теперь у тебя установлен Go и ты можешь запускать Рейна Трекер! 🚂💨**

