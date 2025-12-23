# Quick Start Guide

## 🚀 Running Programs

All programs can be run using the `go run` command:

```bash
# Basic syntax
go run <category>/<program>/main.go

# Examples
go run 01-basics/1_hello/main.go
go run 06-concurrency/9_goroutine/main.go
go run 10-web-apis/crudApi/main.go
```

## 📚 Category Quick Reference

### 01-basics/
**Start here if you're new to Go!**
```bash
go run 01-basics/1_hello/main.go
go run 01-basics/3_variables/main.go
go run 01-basics/4_constant/main.go
```

### 02-control-flow/
**Control structures and loops**
```bash
go run 02-control-flow/5_for/main.go
go run 02-control-flow/6_ifelse/main.go
go run 02-control-flow/7_switchcase/main.go
```

### 03-data-types/
**Arrays, maps, strings, structs, pointers**
```bash
go run 03-data-types/strings/main.go
go run 03-data-types/maps/main.go
go run 03-data-types/structure/main.go
go run 03-data-types/array/main.go
```

### 06-concurrency/
**Goroutines, channels, synchronization**
```bash
go run 06-concurrency/9_goroutine/main.go
go run 06-concurrency/10_channel/main.go
go run 06-concurrency/12_sync_waitgroup_Signal/main.go
```

### 10-web-apis/
**HTTP requests and API interactions**
```bash
go run 10-web-apis/crudApi/main.go
go run 10-web-apis/webRequest/main.go
```

## 🔍 Finding Programs by Topic

### Goroutines
All goroutine-related programs are in `06-concurrency/`:
- `9_goroutine/` - Basic goroutines
- `9_1_goroutine/` - More examples
- `9_2_goroutineEvenOdd/` - Even/odd processing
- `11_goroutne-channel/` - Goroutines with channels

### Strings
- `03-data-types/strings/` - String operations
- `07-algorithms/8_concatenate/` - String concatenation

### Arrays
- `03-data-types/array/` - Basic arrays
- `03-data-types/array/reverseAnArray/` - Reverse array
- `03-data-types/array/largestElementInArray/` - Find largest
- `03-data-types/array/sumOfAllElement/` - Sum elements

### Web & APIs
All in `10-web-apis/`:
- `crudApi/` - CRUD operations
- `webRequest/` - Web requests
- `handleRequest/` - Request handling
- `post/` - POST requests

## 📝 Common Commands

```bash
# List all programs in a category
ls 06-concurrency/

# Count total programs
find . -name "main.go" | wc -l

# Find all goroutine programs
find 06-concurrency -name "main.go"

# Run with output
go run 01-basics/1_hello/main.go

# Build executable
go build -o hello 01-basics/1_hello/main.go
./hello
```

## ⚠️ Important Notes

1. **Module Path**: The project uses `goLangFirst` as the module name
2. **Utility Package**: Located in `utils/myutil/`, imported as `goLangFirst/utils/myutil`
3. **Package Names**: All executables use `package main`
4. **Dependencies**: Run `go mod tidy` to ensure all dependencies are installed

## 🎯 Learning Path

1. **Week 1**: Basics (01-basics, 02-control-flow)
2. **Week 2**: Data Types (03-data-types)
3. **Week 3**: Functions & OOP (04-functions, 05-oop)
4. **Week 4**: Concurrency (06-concurrency)
5. **Week 5**: Advanced Topics (07-algorithms, 08-error-handling, 09-file-operations)
6. **Week 6**: Web Development (10-web-apis, 11-utilities)

## 🐛 Troubleshooting

**Import errors?**
- Make sure you're in the project root directory
- Run `go mod tidy` to update dependencies
- Check that the module path in `go.mod` matches imports

**Program not found?**
- Verify the path: `ls <category>/<program>/`
- Check that `main.go` exists in the directory

**Build errors?**
- Ensure Go is installed: `go version`
- Check for syntax errors in the file
- Verify all imports are correct

---

For detailed information, see [README.md](README.md) and [CATEGORIES.md](CATEGORIES.md)

