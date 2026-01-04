# Testing Your Go Applications on GCP Server

This guide explains various methods to test your Go learning project after deploying it on Google Cloud Platform (GCP).

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Testing Methods](#testing-methods)
3. [SSH Testing](#ssh-testing)
4. [Testing Web APIs](#testing-web-apis)
5. [Automated Testing](#automated-testing)
6. [Monitoring & Logging](#monitoring--logging)

---

## Prerequisites

Before testing, ensure:
- ✅ Go is installed on your GCP server
- ✅ Your application is deployed and running
- ✅ Firewall rules allow necessary ports (8080, 443, etc.)
- ✅ SSH access is configured

---

## Testing Methods

### 1. SSH Testing (Command Line Programs)

For programs that run and exit (like your learning examples):

```bash
# SSH into your GCP server
gcloud compute ssh YOUR_INSTANCE_NAME --zone YOUR_ZONE

# Navigate to your project directory
cd /path/to/golang-first

# Test individual programs
go run 01-basics/hello/main.go
go run 03-data-types/array/main.go
go run 06-concurrency/goroutine/main.go
go run 07-algorithms/primeNo/main.go

# Build and run compiled binaries
go build -o myapp 01-basics/hello/main.go
./myapp

# Test with output redirection (save results)
go run 10-web-apis/crudApi/main.go > test_output.txt
cat test_output.txt
```

### 2. Testing Web APIs (HTTP Servers)

If you have HTTP servers running, use these methods:

#### Method A: Using `curl` from SSH

```bash
# SSH into server
gcloud compute ssh YOUR_INSTANCE_NAME --zone YOUR_ZONE

# Test GET request
curl http://localhost:8080/api/users

# Test POST request
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Pawan","email":"pawan@example.com"}'

# Test with verbose output
curl -v http://localhost:8080/api/users

# Save response to file
curl http://localhost:8080/api/users > response.json
```

#### Method B: Using `curl` from Your Local Machine

```bash
# Replace EXTERNAL_IP with your GCP server's external IP
curl http://EXTERNAL_IP:8080/api/users

# Test from browser
# Open: http://EXTERNAL_IP:8080/api/users
```

#### Method C: Using Postman/Insomnia

1. Open Postman/Insomnia
2. Create new request
3. Set URL: `http://YOUR_GCP_EXTERNAL_IP:8080/api/endpoint`
4. Set method (GET, POST, PUT, DELETE)
5. Add headers if needed
6. Send request

#### Method D: Using Browser

Simply open in your browser:
```
http://YOUR_GCP_EXTERNAL_IP:8080/your-endpoint
```

---

## SSH Testing

### Quick Test Script

Create a test script to run multiple tests:

```bash
#!/bin/bash
# test_all.sh

echo "Testing Go Programs on GCP Server..."
echo "====================================="

echo -e "\n1. Testing Hello World:"
go run 01-basics/hello/main.go

echo -e "\n2. Testing Variables:"
go run 01-basics/variables/main.go

echo -e "\n3. Testing Arrays:"
go run 03-data-types/array/main.go

echo -e "\n4. Testing Concurrency:"
go run 06-concurrency/goroutine/main.go

echo -e "\n5. Testing Web Request:"
go run 10-web-apis/webRequest/main.go

echo -e "\n✅ All tests completed!"
```

Make it executable and run:
```bash
chmod +x test_all.sh
./test_all.sh
```

### Testing with Timeout

For programs that might hang:
```bash
# Run with 10-second timeout
timeout 10 go run 06-concurrency/goroutine/main.go
```

---

## Testing Web APIs

### Creating a Simple Test Server

First, let's create a simple HTTP server for testing:

```go
// test_server/main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

type Response struct {
    Message string `json:"message"`
    Time    string `json:"time"`
    Status  string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{
        Message: "Server is running!",
        Time:    time.Now().Format(time.RFC3339),
        Status:  "healthy",
    }
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from GCP Server!")
    })
    
    port := ":8080"
    fmt.Printf("Server starting on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
```

### Testing the Server

```bash
# 1. Start the server (in background or new terminal)
go run test_server/main.go &

# 2. Test from server itself
curl http://localhost:8080/health
curl http://localhost:8080/

# 3. Test from external machine
curl http://YOUR_EXTERNAL_IP:8080/health

# 4. Check if server is running
ps aux | grep "go run"
netstat -tulpn | grep 8080
```

### Testing Different HTTP Methods

```bash
# GET request
curl http://YOUR_IP:8080/api/users

# POST request
curl -X POST http://YOUR_IP:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Pawan","age":25}'

# PUT request
curl -X PUT http://YOUR_IP:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Pawan Kumar","age":26}'

# DELETE request
curl -X DELETE http://YOUR_IP:8080/api/users/1

# With authentication
curl -H "Authorization: Bearer YOUR_TOKEN" http://YOUR_IP:8080/api/users
```

---

## Automated Testing

### Using Go Test Command

```bash
# Run all tests
go test ./...

# Run tests in verbose mode
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -v ./01-basics/hello

# Run tests with timeout
go test -timeout 30s ./...
```

### Creating Test Scripts

```bash
#!/bin/bash
# automated_test.sh

SERVER_IP="YOUR_GCP_IP"
PORT="8080"

echo "Starting automated tests..."
echo "============================"

# Test 1: Health check
echo -e "\n[Test 1] Health Check"
response=$(curl -s -o /dev/null -w "%{http_code}" http://$SERVER_IP:$PORT/health)
if [ $response -eq 200 ]; then
    echo "✅ Health check passed"
else
    echo "❌ Health check failed: HTTP $response"
fi

# Test 2: GET request
echo -e "\n[Test 2] GET Request"
response=$(curl -s http://$SERVER_IP:$PORT/api/users)
if [ ! -z "$response" ]; then
    echo "✅ GET request successful"
    echo "Response: $response"
else
    echo "❌ GET request failed"
fi

# Test 3: POST request
echo -e "\n[Test 3] POST Request"
response=$(curl -s -X POST http://$SERVER_IP:$PORT/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User"}')
if [ ! -z "$response" ]; then
    echo "✅ POST request successful"
else
    echo "❌ POST request failed"
fi

echo -e "\n✅ All automated tests completed!"
```

---

## Monitoring & Logging

### View Application Logs

```bash
# If running as service
sudo journalctl -u your-service-name -f

# If running with output redirection
tail -f /var/log/myapp.log

# If running in terminal, logs appear directly
```

### Check Server Status

```bash
# Check if Go processes are running
ps aux | grep go

# Check port usage
sudo netstat -tulpn | grep 8080
# OR
sudo ss -tulpn | grep 8080

# Check system resources
top
htop

# Check disk space
df -h

# Check Go version
go version
```

### Real-time Monitoring

```bash
# Watch logs in real-time
tail -f /var/log/application.log

# Monitor network connections
watch -n 1 'netstat -an | grep 8080'

# Monitor CPU/Memory
watch -n 1 'ps aux | grep go'
```

---

## GCP-Specific Testing

### Using Cloud Shell

```bash
# Open GCP Cloud Shell
# Test from Cloud Shell terminal
curl http://YOUR_INTERNAL_IP:8080/health
```

### Using gcloud Command

```bash
# SSH and run command in one line
gcloud compute ssh YOUR_INSTANCE_NAME --zone YOUR_ZONE --command "cd /path/to/project && go run main.go"

# Copy files to server
gcloud compute scp local_file.go YOUR_INSTANCE_NAME:/remote/path/ --zone YOUR_ZONE

# Copy files from server
gcloud compute scp YOUR_INSTANCE_NAME:/remote/path/file.go ./local_path/ --zone YOUR_ZONE
```

### Testing Load Balancer/Cloud Run

If using Cloud Run or Load Balancer:

```bash
# Cloud Run provides HTTPS URL
curl https://YOUR_SERVICE_URL.run.app/health

# Load Balancer
curl http://YOUR_LOAD_BALANCER_IP/health
```

---

## Testing Checklist

Use this checklist when testing:

- [ ] **Connection**: Can you SSH into the server?
- [ ] **Go Installation**: Is Go installed and in PATH?
- [ ] **Code Deployment**: Is your code deployed on the server?
- [ ] **Compilation**: Do programs compile without errors?
- [ ] **Basic Programs**: Do standalone programs run successfully?
- [ ] **Web Servers**: Are HTTP servers starting on correct ports?
- [ ] **Firewall**: Are ports open (8080, 443, etc.)?
- [ ] **External Access**: Can you access from outside the server?
- [ ] **Logs**: Are logs being generated and accessible?
- [ ] **Resources**: Are CPU/Memory usage normal?
- [ ] **Errors**: Check for any error messages in logs

---

## Common Issues & Solutions

### Issue 1: Connection Refused
```bash
# Check if server is running
ps aux | grep go

# Check firewall rules
sudo ufw status
# OR
gcloud compute firewall-rules list
```

### Issue 2: Port Already in Use
```bash
# Find process using port
sudo lsof -i :8080
# OR
sudo netstat -tulpn | grep 8080

# Kill the process
kill -9 PID
```

### Issue 3: Permission Denied
```bash
# Make file executable
chmod +x script.sh

# Run with sudo if needed
sudo go run main.go
```

### Issue 4: Module Not Found
```bash
# Download dependencies
go mod download
go mod tidy

# Verify go.mod
cat go.mod
```

---

## Best Practices

1. **Use Environment Variables**: Store config in environment variables
2. **Logging**: Implement proper logging (log levels: INFO, ERROR, DEBUG)
3. **Health Checks**: Create `/health` endpoint for monitoring
4. **Error Handling**: Proper error handling and reporting
5. **Testing Locally First**: Test on local machine before deploying
6. **Version Control**: Keep code in Git (GitHub, GitLab)
7. **Backup**: Regularly backup important data
8. **Monitoring**: Set up monitoring and alerts
9. **Security**: Use HTTPS in production, implement authentication
10. **Documentation**: Document your API endpoints

---

## Quick Reference

```bash
# SSH into GCP
gcloud compute ssh INSTANCE_NAME --zone ZONE

# Run Go program
go run path/to/main.go

# Build binary
go build -o app path/to/main.go

# Test HTTP endpoint
curl http://IP:PORT/endpoint

# Check logs
tail -f /var/log/app.log

# Check if running
ps aux | grep go

# Check port
netstat -tulpn | grep PORT
```

---

## Example: Complete Testing Workflow

```bash
# 1. SSH into server
gcloud compute ssh my-instance --zone us-central1-a

# 2. Navigate to project
cd /home/user/golang-first

# 3. Pull latest code (if using Git)
git pull origin main

# 4. Install/update dependencies
go mod tidy

# 5. Build application
go build -o myapp 10-web-apis/test_server/main.go

# 6. Run in background
nohup ./myapp > app.log 2>&1 &

# 7. Test locally
curl http://localhost:8080/health

# 8. Test externally (from your local machine)
curl http://EXTERNAL_IP:8080/health

# 9. Check logs
tail -f app.log

# 10. Monitor
ps aux | grep myapp
```

---

Happy Testing! 🚀

For more help:
- GCP Documentation: https://cloud.google.com/docs
- Go Documentation: https://golang.org/doc/
- HTTP Testing: https://httpie.io/ or https://www.postman.com/
