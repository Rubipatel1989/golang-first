# Project Categories Overview

This document provides a quick reference guide to all categories and their contents.

## 📂 Category Index

### 01-basics/ - Fundamentals
Essential Go programming basics for beginners.

| Program | Description |
|---------|-------------|
| `1_hello/` | Hello World program |
| `2_simple_value/` | Simple value assignments |
| `3_variables/` | Variable declarations |
| `4_constant/` | Constants usage |
| `print/` | Print statements |
| `hello-with-util/` | Using custom utility package |

### 02-control-flow/ - Control Structures
Conditional statements and loops.

| Program | Description |
|---------|-------------|
| `5_for/` | For loop basics |
| `6_ifelse/` | If-else statements |
| `7_switchcase/` | Switch-case statements |
| `forLoop/` | Advanced for loops |
| `ifElse/` | Conditional logic |
| `switch/` | Switch variations |

### 03-data-types/ - Data Structures
Working with arrays, maps, strings, structs, and pointers.

| Program | Description |
|---------|-------------|
| `array/` | Array operations |
| `array/arrayToSlice/` | Array to slice conversion |
| `array/largestElementInArray/` | Find largest element |
| `array/reverseAnArray/` | Reverse array |
| `array/sumOfAllElement/` | Sum array elements |
| `maps/` | Map data structure |
| `strings/` | String manipulation |
| `structure/` | Struct types |
| `pointer/` | Pointers |

### 04-functions/ - Functions
Function definitions and defer statements.

| Program | Description |
|---------|-------------|
| `function/` | Function basics |
| `deferKeyword/` | Defer statement |

### 05-oop/ - Object-Oriented Programming
Struct embedding and composition.

| Program | Description |
|---------|-------------|
| `3_inheritance/` | Basic inheritance |
| `3_inheritance_1/` | Advanced inheritance |

### 06-concurrency/ - Concurrency & Parallelism
Goroutines, channels, and synchronization.

| Program | Description |
|---------|-------------|
| `9_goroutine/` | Basic goroutines |
| `9_1_goroutine/` | Goroutine examples |
| `9_2_goroutineEvenOdd/` | Even/odd with goroutines |
| `10_channel/` | Channel communication |
| `11_goroutne-channel/` | Goroutines + channels |
| `12_sync_waitgroup_Signal/` | WaitGroup synchronization |
| `13_oddEvenNo/` | Odd/even separation |

### 07-algorithms/ - Algorithms
Common algorithms and problem-solving.

| Program | Description |
|---------|-------------|
| `7_1_remove_duplicate/` | Remove duplicates |
| `8_concatenate/` | Concatenation operations |
| `14_primeNo/` | Prime number detection |

### 08-error-handling/ - Error Handling
Error handling patterns.

| Program | Description |
|---------|-------------|
| `errorHadling/` | Error handling examples |

### 09-file-operations/ - File I/O
File reading and writing.

| Program | Description |
|---------|-------------|
| `fileHandle/` | File operations |
| `fileHandle/example.txt` | Sample text file |

### 10-web-apis/ - Web Development
HTTP requests and API interactions.

| Program | Description |
|---------|-------------|
| `crudApi/` | CRUD API operations |
| `handleRequest/` | HTTP request handling |
| `webRequest/` | Web requests |
| `post/` | POST requests |

### 11-utilities/ - Utilities
Helper functions and common tasks.

| Program | Description |
|---------|-------------|
| `json/` | JSON operations |
| `timePackage/` | Time/date operations |
| `userInput/` | User input reading |
| `dataConversion/` | Type conversions |
| `15_AssertionType/` | Type assertions |

### utils/ - Custom Packages
Reusable utility packages.

| Package | Description |
|---------|-------------|
| `myutil/` | Custom utility functions |

## 🎯 Quick Navigation

### By Topic

**Goroutines & Concurrency:**
- `06-concurrency/9_goroutine/`
- `06-concurrency/9_1_goroutine/`
- `06-concurrency/9_2_goroutineEvenOdd/`
- `06-concurrency/10_channel/`
- `06-concurrency/11_goroutne-channel/`
- `06-concurrency/12_sync_waitgroup_Signal/`
- `06-concurrency/13_oddEvenNo/`

**Strings:**
- `03-data-types/strings/`
- `07-algorithms/8_concatenate/`

**Arrays:**
- `03-data-types/array/`
- `03-data-types/array/arrayToSlice/`
- `03-data-types/array/largestElementInArray/`
- `03-data-types/array/reverseAnArray/`
- `03-data-types/array/sumOfAllElement/`

**Web & APIs:**
- `10-web-apis/crudApi/`
- `10-web-apis/handleRequest/`
- `10-web-apis/webRequest/`
- `10-web-apis/post/`

**Data Types:**
- `03-data-types/maps/`
- `03-data-types/structure/`
- `03-data-types/pointer/`

## ✅ All Programs Status

All programs are properly organized with:
- ✅ Correct `package main` declarations
- ✅ Proper folder structure
- ✅ Logical categorization
- ✅ Ready to run with `go run <path>/main.go`

## 📌 Notes

- Each program is self-contained
- All use `package main` for executables
- Utility package uses `package myutil`
- Import paths updated for new structure

