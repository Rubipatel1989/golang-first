# Test Server

A simple HTTP server for testing your Go applications on GCP.

## Usage

### Run Locally
```bash
go run main.go
```

### Build and Run
```bash
go build -o test-server main.go
./test-server
```

### Test Endpoints

1. **Root endpoint**
   ```bash
   curl http://localhost:8080/
   ```

2. **Health check**
   ```bash
   curl http://localhost:8080/health
   ```

3. **Get all users**
   ```bash
   curl http://localhost:8080/api/users
   ```

4. **Get user by ID**
   ```bash
   curl http://localhost:8080/api/user?id=1
   ```

### Test on GCP Server

1. **SSH into your server**
   ```bash
   gcloud compute ssh YOUR_INSTANCE_NAME --zone YOUR_ZONE
   ```

2. **Navigate to directory**
   ```bash
   cd /path/to/golang-first/10-web-apis/testServer
   ```

3. **Run the server**
   ```bash
   go run main.go
   # OR run in background
   nohup go run main.go > server.log 2>&1 &
   ```

4. **Test from server**
   ```bash
   curl http://localhost:8080/health
   ```

5. **Test from external machine**
   ```bash
   curl http://YOUR_GCP_EXTERNAL_IP:8080/health
   ```

## Notes

- Server runs on port 8080 by default
- Make sure firewall allows port 8080
- For production, use environment variables for configuration
- Add authentication/authorization as needed
