# Go Learning Project

A comprehensive collection of Go programming examples organized by category for learning and reference.

## 📁 Project Structure

This project is organized into logical categories to help you navigate and learn Go concepts systematically.

### 01-basics/
Fundamental Go programming concepts for beginners.

- **1_hello/** - Basic "Hello World" program
- **2_simple_value/** - Working with simple values and variables
- **3_variables/** - Variable declarations and types
- **4_constant/** - Constants in Go
- **print/** - Print statements and formatting
- **hello-with-util/** - Hello world using custom utility package

### 02-control-flow/
Control flow statements and loops.

- **5_for/** - For loops
- **6_ifelse/** - If-else statements
- **7_switchcase/** - Switch-case statements
- **forLoop/** - Advanced for loop examples
- **ifElse/** - Conditional statements
- **switch/** - Switch statement variations

### 03-data-types/
Working with different data types and data structures.

- **array/** - Array operations
  - **arrayToSlice/** - Converting arrays to slices
  - **largestElementInArray/** - Finding largest element
  - **reverseAnArray/** - Reversing arrays
  - **sumOfAllElement/** - Sum of array elements
- **maps/** - Map data structure
- **strings/** - String manipulation and operations
- **structure/** - Struct types and nested structures
- **pointer/** - Pointers and memory addresses

### 04-functions/
Functions and function-related concepts.

- **function/** - Function definitions and calls
- **deferKeyword/** - Defer statement for cleanup

### 05-oop/
Object-oriented programming concepts in Go.

- **3_inheritance/** - Struct embedding and composition
- **3_inheritance_1/** - Advanced inheritance patterns

### 06-concurrency/
Concurrency, goroutines, channels, and synchronization.

- **9_goroutine/** - Basic goroutines
- **9_1_goroutine/** - Goroutine examples
- **9_2_goroutineEvenOdd/** - Goroutines for even/odd number processing
- **10_channel/** - Channel communication
- **11_goroutne-channel/** - Combining goroutines and channels
- **12_sync_waitgroup_Signal/** - WaitGroup for synchronization
- **13_oddEvenNo/** - Odd/even number separation using goroutines

### 07-algorithms/
Common algorithms and problem-solving examples.

- **7_1_remove_duplicate/** - Removing duplicates from collections
- **8_concatenate/** - String/array concatenation
- **14_primeNo/** - Prime number detection

### 08-error-handling/
Error handling patterns and best practices.

- **errorHadling/** - Error handling examples

### 09-file-operations/
File I/O operations.

- **fileHandle/** - Reading and writing files
  - **example.txt** - Sample text file

### 10-web-apis/
Web development and API interactions.

- **crudApi/** - CRUD operations with REST APIs
- **handleRequest/** - HTTP request handling
- **webRequest/** - Making web requests
- **post/** - POST request examples

### 11-utilities/
Utility functions and helper examples.

- **json/** - JSON encoding and decoding
- **timePackage/** - Time and date operations
- **userInput/** - Reading user input
- **dataConversion/** - Type conversions
- **15_AssertionType/** - Type assertions

### utils/
Custom utility packages.

- **myutil/** - Custom utility functions package

## 🚀 Getting Started

### Prerequisites

- Go 1.23.4 or later
- Basic understanding of command-line interface

### Installation

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd golang-first
   ```

3. Initialize Go modules (if not already done):
   ```bash
   go mod init goLangFirst
   ```

4. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running Examples

Each folder contains a `main.go` file that can be executed independently:

```bash
# Example: Run a basic program
go run 01-basics/1_hello/main.go

# Example: Run a concurrency example
go run 06-concurrency/9_goroutine/main.go

# Example: Run a web API example
go run 10-web-apis/crudApi/main.go
```

## 📚 Learning Path

### Beginner Path
1. Start with **01-basics/** to understand fundamentals
2. Move to **02-control-flow/** for control structures
3. Learn **03-data-types/** for data structures
4. Practice with **04-functions/** for reusable code

### Intermediate Path
5. Explore **05-oop/** for object-oriented concepts
6. Study **06-concurrency/** for parallel programming
7. Practice **07-algorithms/** for problem-solving
8. Learn **08-error-handling/** for robust code

### Advanced Path
9. Work with **09-file-operations/** for I/O
10. Build **10-web-apis/** for web development
11. Use **11-utilities/** for common tasks

## 🔧 Setup Instructions (Ubuntu)

If you're using Ubuntu, follow these steps:

1. Download Go:
   ```bash
   sudo wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
   ```

2. Extract to `/usr/local`:
   ```bash
   sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
   ```

3. Add to PATH by editing `~/.bashrc`:
   ```bash
   nano ~/.bashrc
   ```
   
   Add these lines:
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=$HOME/go 
   export PATH=$PATH:$GOPATH/bin
   ```

4. Reload configuration:
   ```bash
   source ~/.bashrc
   ```

5. Verify installation:
   ```bash
   go version
   ```

## 📖 Tutorial Resources

- [Go Tutorial Series](https://www.youtube.com/watch?v=_utEjVtOwFE&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=13)
- [Go Programming Course](https://www.youtube.com/watch?v=yZgwW6Yuc_E&t=1635s)

## 📝 Notes

- All programs use `package main` for executables
- The `utils/myutil` package provides reusable utility functions
- Each example is self-contained and can be run independently
- Some examples may require internet connection (web API examples)

## 🤝 Contributing

Feel free to add more examples or improve existing ones. Make sure to:
- Follow Go naming conventions
- Add comments explaining complex logic
- Keep examples simple and educational
- Test all programs before committing

## 📄 License

This is a learning project. Feel free to use and modify as needed.

---

**Happy Learning! 🎓**
