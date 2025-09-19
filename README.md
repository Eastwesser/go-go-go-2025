# Golang Interviews

Here we study Go from simple fundamential choices to the enterprize decisions.

## Flags:

```bash
-f # stands for "force"
-y # stands for "yes"
-la # list all
```

## Delve debugger commands:

```bash
dlv run
```

## Git commands:

```bash
git add .
git commit -m "#v.0.0.1 task commit: initial commit"
git push origin dev # only admin is allowed to merge with main
```

## Docker:

```bash
docker ps
```

## Docker-Compose:

```bash
docker-compose up && docker-compose build
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
