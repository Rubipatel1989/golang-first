#!/bin/bash

# GCP Testing Script for Go Learning Project
# This script helps test your Go applications on GCP server

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration (modify these for your setup)
SERVER_IP="YOUR_GCP_EXTERNAL_IP"  # Replace with your GCP external IP
PORT="8080"
PROJECT_DIR="/path/to/golang-first"  # Replace with your project path on server

echo -e "${YELLOW}========================================${NC}"
echo -e "${YELLOW}  GCP Go Application Testing Script${NC}"
echo -e "${YELLOW}========================================${NC}"

# Function to test HTTP endpoint
test_endpoint() {
    local endpoint=$1
    local expected_status=${2:-200}
    
    echo -e "\n${YELLOW}Testing: $endpoint${NC}"
    
    response=$(curl -s -o /dev/null -w "%{http_code}" "http://$SERVER_IP:$PORT$endpoint")
    
    if [ "$response" -eq "$expected_status" ]; then
        echo -e "${GREEN}✅ PASSED${NC} - HTTP Status: $response"
        # Get actual response
        curl -s "http://$SERVER_IP:$PORT$endpoint" | head -5
        return 0
    else
        echo -e "${RED}❌ FAILED${NC} - Expected: $expected_status, Got: $response"
        return 1
    fi
}

# Function to test if server is running
test_server_running() {
    echo -e "\n${YELLOW}Checking if server is running...${NC}"
    
    if curl -s --connect-timeout 5 "http://$SERVER_IP:$PORT/health" > /dev/null 2>&1; then
        echo -e "${GREEN}✅ Server is running${NC}"
        return 0
    else
        echo -e "${RED}❌ Server is not responding${NC}"
        echo -e "${YELLOW}Please make sure:${NC}"
        echo "  1. Server is running on port $PORT"
        echo "  2. Firewall allows port $PORT"
        echo "  3. Server IP is correct: $SERVER_IP"
        return 1
    fi
}

# Main testing function
main() {
    # Update configuration from command line arguments if provided
    if [ "$1" != "" ]; then
        SERVER_IP=$1
    fi
    
    if [ "$2" != "" ]; then
        PORT=$2
    fi
    
    echo -e "Server IP: ${GREEN}$SERVER_IP${NC}"
    echo -e "Port: ${GREEN}$PORT${NC}"
    
    # Test if server is running
    if ! test_server_running; then
        echo -e "\n${RED}Cannot proceed with tests. Server is not accessible.${NC}"
        exit 1
    fi
    
    # Run tests
    echo -e "\n${YELLOW}Starting endpoint tests...${NC}"
    
    passed=0
    failed=0
    
    # Test root endpoint
    if test_endpoint "/" 200; then
        ((passed++))
    else
        ((failed++))
    fi
    
    # Test health endpoint
    if test_endpoint "/health" 200; then
        ((passed++))
    else
        ((failed++))
    fi
    
    # Test API endpoints
    if test_endpoint "/api/users" 200; then
        ((passed++))
    else
        ((failed++))
    fi
    
    # Summary
    echo -e "\n${YELLOW}========================================${NC}"
    echo -e "${YELLOW}Test Summary${NC}"
    echo -e "${YELLOW}========================================${NC}"
    echo -e "${GREEN}Passed: $passed${NC}"
    echo -e "${RED}Failed: $failed${NC}"
    echo -e "${YELLOW}Total: $((passed + failed))${NC}"
    
    if [ $failed -eq 0 ]; then
        echo -e "\n${GREEN}🎉 All tests passed!${NC}"
        exit 0
    else
        echo -e "\n${RED}⚠️  Some tests failed${NC}"
        exit 1
    fi
}

# Help message
show_help() {
    echo "Usage: $0 [SERVER_IP] [PORT]"
    echo ""
    echo "Options:"
    echo "  SERVER_IP    GCP server external IP (default: $SERVER_IP)"
    echo "  PORT         Server port (default: $PORT)"
    echo ""
    echo "Examples:"
    echo "  $0                                    # Use default IP and port"
    echo "  $0 34.123.45.67                      # Specify IP"
    echo "  $0 34.123.45.67 8080                 # Specify IP and port"
    echo ""
    echo "Before running:"
    echo "  1. Update SERVER_IP in this script OR pass as argument"
    echo "  2. Make sure your server is running"
    echo "  3. Ensure firewall allows the port"
}

# Check for help flag
if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
    show_help
    exit 0
fi

# Run main function
main "$@"
