# Golang Interviews

Here we study Go from simple fundamential choices to the enterprise decisions.

## Flags:

```bash
# Common UNIX/Linux flags:
-f, --force       # force operation, skip confirmation
-y, --yes         # automatically answer "yes" to all prompts
-la, --list-all   # list all items (including hidden)
-v, --verbose     # verbose output, show details
-q, --quiet       # quiet mode, minimal output
-r, --recursive   # recursive operation
-d, --debug       # enable debug mode
-h, --help        # show help message
--version         # show version information

# Go-specific flags:
-race             # enable race detector
-cover            # enable code coverage
-tags             # build tags
-ldflags          # linker flags
-gcflags          # compiler flags
```

## Git commands:

Before writing a commit message, check this table:
```bash
# Основные типы коммитов (types):

# Type	    # Когда использовать	                        Пример
feat	    # Новая функциональность	                    feat: add user authentication
fix	        # Исправление бага	                            fix: resolve memory leak in cache
docs	    # Изменения в документации	                    docs: update API documentation
style	    # Форматирование, пробелы	                    style: format code with gofmt
refactor	# Рефакторинг без изменения функциональности	refactor: simplify database layer
test	    # Добавление тестов	                            test: add unit tests for service
chore	    # Вспомогательные задачи	                    chore: update dependencies
build	    # Сборка системы	                            build: update Dockerfile
ci	        # CI конфигурация	                            ci: add GitHub Actions workflow
```

```bash
# Daily commands:
git add -u
git commit -m "feat(basetypes): implement integer operations with tests"
git commit -m "feat(leetcode): implement palindrome with two pointers method"
git push origin dev # only admin is allowed to merge with main

### Basic workflow ###

# Статус и информация
git status                      # показать состояние рабочей директории
git log                         # история коммитов
git log --oneline --graph       # компактная история с графиком
git diff                        # показать изменения в файлах
git show <commit>               # показать изменения в коммите

# Добавление и коммиты
git add .                       # добавить все изменения (⚠️ осторожно!)
git add -u                      # добавить только измененные файлы
git add -p                      # интерактивное добавление изменений
git commit -m "message"         # коммит с сообщением
git commit --amend              # исправить последний коммит
git commit --allow-empty        # пустой коммит (для триггеров)

# Ветки
git branch                      # список веток
git branch <name>               # создать новую ветку
git checkout <branch>           # переключиться на ветку
git checkout -b <branch>        # создать и переключиться
git merge <branch>              # слить ветку в текущую
git rebase <branch>             # перебазировать текущую ветку

# Push/Pull
git push origin dev             # отправить в удаленную ветку
git push -u origin dev          # отправить и установить upstream
git push --force-with-lease     # безопасный force push
git pull origin dev             # получить изменения
git fetch --all                 # получить все изменения без мерджа

# Undo и Reset
git reset --soft HEAD~1         # отменить коммит, сохранить изменения
git reset --hard HEAD~1         # полностью отменить коммит
git restore <file>              # отменить изменения в файле
git clean -fd                   # удалить неотслеживаемые файлы

# Stash
git stash                       # временно сохранить изменения
git stash pop                   # восстановить изменения
git stash list                  # список stash'ей

# Теги
git tag v1.0.0                  # создать тег
git push origin --tags          # отправить теги на remote

### Professional Workflow ###

# Convention Commits
git commit -m "feat: add user authentication"       # for adding new files
git commit -m "fix: resolve memory leak"            # for changes and upgrades
git commit -m "docs: update API documentation"      # for README.md update 
git commit -m "refactor: simplify database layer"   # for directories decomposition
git commit -m "test: add unit tests for service"    # for tests only

# Interactive Rebase
git rebase -i HEAD~5            # интерактивный rebase последних 5 коммитов

# Cherry-pick
git cherry-pick <commit-hash>   # применить конкретный коммит

# Bisect для поиска багов
git bisect start
git bisect bad
git bisect good <commit>
```

## Docker:

```bash
docker ps                       # список running контейнеров
docker ps -a                    # список всех контейнеров (включая stopped)
docker start <container>        # запустить контейнер
docker stop <container>         # остановить контейнер
docker restart <container>      # перезапустить контейнер
docker rm <container>           # удалить контейнер
docker rm -f <container>        # принудительно удалить running контейнер

# Информация
docker logs <container>         # логи контейнера
docker logs -f <container>      # логи в реальном времени
docker inspect <container>      # детальная информация о контейнере
docker stats                    # live usage statistics
docker top <container>          # процессы внутри контейнера

# Images
docker images                   # список образов
docker rmi <image>              # удалить образ
docker pull <image>             # скачать образ из registry
docker push <image>             # загрузить образ в registry

# Build
docker build -t myapp:latest .  # собрать образ с тегом
docker build --no-cache .       # собрать без кэша

# Execution
docker run -it <image> bash         # запустить контейнер с интерактивным shell
docker run -p 8080:80 <image>       # пробросить порты
docker run -v $(pwd):/app <image>   # монтировать volume
docker run --env VAR=value <image>  # установить environment variable

# Cleanup
docker system prune             # удалить все остановленные контейнеры, неиспользуемые образы
docker system prune -a          # полная очистка (⚠️ осторожно!)
```

## Docker-Compose:

```bash
### Basic Operations Basic Operations ###
docker-compose up                   # запустить все сервисы
docker-compose up -d                # запустить в фоновом режиме
docker-compose down                 # остановить и удалить контейнеры
docker-compose down -v              # остановить и удалить volumes
docker-compose build                # собрать образы
docker-compose build --no-cache     # собрать без кэша

# Service Management:
docker-compose start            # запустить сервисы
docker-compose stop             # остановить сервисы
docker-compose restart          # перезапустить сервисы
docker-compose pause            # приостановить сервисы
docker-compose unpause          # возобновить сервисы

# Information:
docker-compose ps               # статус сервисов
docker-compose logs             # логи всех сервисов
docker-compose logs -f          # логи в реальном времени
docker-compose logs <service>   # логи конкретного сервиса
docker-compose top              # процессы в контейнерах

# Execution:
docker-compose exec <service> bash      # shell в running контейнере
docker-compose run <service> <command>  # запустить команду в новом контейнере

# Advanced:
docker-compose up --scale web=3     # запустить несколько инстансов сервиса
docker-compose config               # проверить конфигурацию
docker-compose pull                 # скачать последние образы

# Development Workflow:
docker-compose up --build                           # собрать и запустить
docker-compose up --build -d                        # собрать и запустить в фоне
docker-compose down && docker-compose up --build    # полный перезапуск
```

## Delve debugger commands:

```bash
# Basic Debugging:
dlv debug ./main.go              # компилирует и запускает
dlv exec ./myapp                 # отладка уже скомпилированного бинарника
dlv test                         # отладка тестов
dlv attach <pid>                 # присоединиться к работающему процессу

# Main commands inside Delve:
(dlv) break main.main            # брейкпоинт на функцию main.main
(dlv) break main.go:15           # брейкпоинт на строку 15
(dlv) continue                   # продолжить выполнение (c)
(dlv) next                       # следующая строка (n)
(dlv) step                       # шаг с заходом в функцию (s)
(dlv) stepout                    # шаг из текущей функции (so)
(dlv) restart                    # перезапуск программы
(dlv) quit                       # выход из отладчика

# Inspection Commands:
(dlv) print variableName         # напечатать значение переменной (p)
(dlv) locals                     # показать все локальные переменные
(dlv) args                       # показать аргументы функции
(dlv) goroutines                 # список всех горутин
(dlv) goroutine <id>             # переключиться на горутину
(dlv) stack                      # показать стек вызовов
(dlv) threads                    # список потоков

# Breakpoint Management:
(dlv) breakpoints                # список всех брейкпоинтов
(dlv) clear <breakpoint-id>      # удалить брейкпоинт
(dlv) clearall                   # удалить все брейкпоинты
(dlv) condition <id> expr        # установить условие для брейкпоинта
(dlv) on <id> cmd                # выполнить команду при срабатывании брейкпоинта

# Execution Control:
(dlv) continue                   # продолжить до следующего брейкпоинта
(dlv) next                       # шаг через (не заходя в функции)
(dlv) step                       # шаг с заходом в функции
(dlv) step-instruction           # шаг на одну инструкцию
(dlv) reverse-step               # шаг назад (требуется запись execution)
(dlv) reverse-continue           # продолжить назад

# Variable Manipulation:
(dlv) set variable = value       # изменить значение переменной
(dlv) whatis variable            # показать тип переменной
(dlv) types regexp               # поиск типов по regexp

# Goroutine Debugging:
(dlv) goroutine <id> stack       # стек конкретной горутины
(dlv) goroutine <id> bp          # установить брейкпоинт для конкретной горутины
(dlv) goroutines -with user      # показать горутины с пользовательским кодом

# Advanced Features:
(dlv) source list                # показать исходный код
(dlv) disassemble                # дизассемблировать текущую функцию
(dlv) regs                       # показать регистры процессора
(dlv) config                     # показать конфигурацию отладки
(dlv) trace functionName         # трассировка вызовов функции

# Useful Aliases:
(dlv) b = break                  # алиас для break
(dlv) c = continue               # алиас для continue  
(dlv) n = next                   # алиас для next
(dlv) s = step                   # алиас для step
(dlv) p = print                  # алиас для print

### DEBUGGING SESSION AS IS ###

# Запуск
dlv debug ./main.go

# Установка брейкпоинтов
(dlv) b main.main
(dlv) b main.go:23

# Запуск
(dlv) c

# Когда программа остановится на main.main
(dlv) n
(dlv) p someVariable
(dlv) s
(dlv) goroutines
(dlv) c
```

## Project tree (actual for 20.09.2025):
```text
Use this command: tree -L 9
.
├── codereview
│   ├── 10.tcp&udp
│   ├── 11.http
│   ├── 12.rest_api
│   ├── 13.rpc
│   ├── 14.grpc
│   ├── 15.system_design
│   ├── 1.basetypes
│   ├── 2.cruds
│   ├── 3.sync
│   ├── 4.concurrency
│   ├── 5.runtime
│   ├── 6.profiling
│   ├── 7.oop
│   ├── 8.patterns
│   └── 9.algos
│       └── leetcode
├── companies
│   ├── avito
│   ├── mts
│   ├── ozon
│   ├── samokat
│   ├── wildberries
│   └── yandex
├── golang
│   └── Dennis
│       ├── part1_fundamentials
│       │   ├── 1.basetypes
│       │   │   ├── task1_ints_uints
│       │   │   │   ├── homework
│       │   │   │   │   ├── main.go
│       │   │   │   │   └── main_test.go
│       │   │   │   └── lections
│       │   │   │       ├── complex_nums_128
│       │   │   │       │   └── main.go
│       │   │   │       ├── complex_nums_64
│       │   │   │       │   └── main.go
│       │   │   │       ├── int
│       │   │   │       │   └── main.go
│       │   │   │       ├── int16_uint16
│       │   │   │       │   └── main.go
│       │   │   │       ├── int32_uint32
│       │   │   │       │   └── main.go
│       │   │   │       ├── int64_uint64
│       │   │   │       │   └── main.go
│       │   │   │       ├── int8_uint8
│       │   │   │       │   └── main.go
│       │   │   │       └── uintptr
│       │   │   │           └── main.go
│       │   │   ├── task2_floats
│       │   │   │   ├── homework
│       │   │   │   │   ├── main.go
│       │   │   │   │   └── main_test.go
│       │   │   │   └── lections
│       │   │   │       ├── float32
│       │   │   │       │   └── main.go
│       │   │   │       └── float64
│       │   │   │           └── main.go
│       │   │   ├── task3_strings
│       │   │   │   ├── homework
│       │   │   │   │   ├── case_aboba
│       │   │   │   │   │   ├── main.go
│       │   │   │   │   │   └── main_test.go
│       │   │   │   │   ├── main.go
│       │   │   │   │   └── main_test.go
│       │   │   │   └── lections
│       │   │   │       └── main.go
│       │   │   ├── task4_arrays
│       │   │   │   ├── homework
│       │   │   │   │   ├── main.go
│       │   │   │   │   └── main_test.go
│       │   │   │   └── lections
│       │   │   │       └── main.go
│       │   │   ├── task5_slices
│       │   │   │   ├── homework
│       │   │   │   │   ├── main.go
│       │   │   │   │   └── main_test.go
│       │   │   │   └── lections
│       │   │   │       └── main.go
│       │   │   └── task6_maps
│       │   │       ├── homework
│       │   │       │   ├── main.go
│       │   │       │   └── main_test.go
│       │   │       └── lections
│       │   │           ├── old_map
│       │   │           │   └── main.go
│       │   │           └── swiss_map
│       │   │               └── main.go
│       │   ├── 2.composites
│       │   │   ├── task1_struct
│       │   │   │   └── main.go
│       │   │   ├── task2_interface
│       │   │   │   └── main.go
│       │   │   ├── task3_constructor
│       │   │   │   └── main.go
│       │   │   ├── task4_method
│       │   │   │   └── main.go
│       │   │   └── task5_crud
│       │   │       ├── example_refactoring
│       │   │       │   ├── cringe.go
│       │   │       │   │   └── main.go
│       │   │       │   └── main.go
│       │   │       └── main.go
│       │   ├── 3.sync
│       │   │   ├── task1_goroutine
│       │   │   │   └── main.go
│       │   │   ├── task2_chan
│       │   │   │   └── main.go
│       │   │   ├── task3_mutex
│       │   │   │   └── main.go
│       │   │   ├── task4_wg
│       │   │   │   └── main.go
│       │   │   ├── task5_context
│       │   │   │   └── main.go
│       │   │   └── task6_sync_map
│       │   │       └── main.go
│       │   ├── 4.concurrency
│       │   │   ├── task1_generator
│       │   │   │   └── main.go
│       │   │   ├── task2_pipeline
│       │   │   │   └── main.go
│       │   │   ├── task3_fan_in_out
│       │   │   │   ├── in
│       │   │   │   │   └── main.go
│       │   │   │   └── out
│       │   │   │       └── main.go
│       │   │   ├── task4_worker_pool
│       │   │   │   └── main.go
│       │   │   ├── task5_semaphore
│       │   │   │   └── main.go
│       │   │   ├── task6_single_flight
│       │   │   │   └── main.go
│       │   │   └── task7_extras
│       │   │       ├── atomics
│       │   │       │   └── main.go
│       │   │       ├── barrier
│       │   │       │   └── main.go
│       │   │       ├── error_handling
│       │   │       │   └── main.go
│       │   │       ├── fan_in_out
│       │   │       │   └── main.go
│       │   │       ├── generics
│       │   │       │   └── main.go
│       │   │       ├── promise
│       │   │       │   └── main.go
│       │   │       ├── semaphore
│       │   │       │   └── main.go
│       │   │       └── worker_pool
│       │   │           └── main.go
│       │   ├── 5.runtime
│       │   │   ├── task1_scheduler
│       │   │   │   └── main.go
│       │   │   ├── task2_gc
│       │   │   │   └── main.go
│       │   │   ├── task3_memory
│       │   │   │   └── main.go
│       │   │   └── task4_gomaxprocs
│       │   │       └── main.go
│       │   └── 6.profiling
│       │       ├── pprof
│       │       │   └── main.go
│       │       └── trace
│       │           └── main.go
│       ├── part2_oop_patterns
│       │   ├── 1.oop
│       │   │   └── main.go
│       │   ├── 2.patterns
│       │   │   └── main.go
│       │   └── 3.algos
│       │       └── main.go
│       └── part3_servers
│           ├── 1.tcp_udp
│           │   ├── task1_tcp
│           │   │   ├── client
│           │   │   │   └── main.go
│           │   │   └── server
│           │   │       └── main.go
│           │   └── task2_udp
│           │       ├── client
│           │       │   └── main.go
│           │       └── server
│           │           └── main.go
│           ├── 2.http
│           │   ├── client
│           │   │   └── main.go
│           │   └── server
│           │       └── main.go
│           ├── 3.rest_api
│           │   ├── cmd
│           │   │   └── main.go
│           │   └── internal
│           │       └── main.go
│           ├── 4.rpc
│           │   └── main.go
│           ├── 5.grpc
│           │   └── main.go
│           ├── 5.system_design
│           │   └── main.go
│           └── README.md
├── go.mod
├── leetcode
│   └── Dennis
│       ├── 10.gas_station
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 1.palindrome
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 2.two_sum
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 3.valid_anagram
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 4.merge-intervals
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 5.sort-colors
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 6.reverse-linked-list
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 7.first_occurrence
│       │   ├── main.go
│       │   └── main_test.go
│       ├── 8.valid_sudoku
│       │   ├── main.go
│       │   └── main_test.go
│       └── 9.scramble-string
│           ├── main.go
│           └── main_test.go
├── main.go
├── README.md
└── sql
    └── Dennis
        ├── find_users.sql
        └── merge_tables.sql

136 directories, 103 files
```
