#!/bin/bash

# Mobile Forensics Toolkit Build Script
# This script builds the application for different platforms

set -e

echo "=== Mobile Forensics Toolkit Build Script ==="
echo "Building the mobile forensics analysis tool..."
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Wails is installed
print_status "Checking Wails CLI installation..."
if ! command -v wails &> /dev/null; then
    print_error "Wails CLI is not installed."
    echo "Please install it with: go install github.com/wailsapp/wails/v2/cmd/wails@latest"
    exit 1
fi
print_success "Wails CLI found: $(wails version)"

# Check if Go is installed
print_status "Checking Go installation..."
if ! command -v go &> /dev/null; then
    print_error "Go is not installed."
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi
print_success "Go found: $(go version)"

# Check if Node.js is installed
print_status "Checking Node.js installation..."
if ! command -v node &> /dev/null; then
    print_error "Node.js is not installed."
    echo "Please install Node.js from https://nodejs.org/"
    exit 1
fi
print_success "Node.js found: $(node --version)"

# Install Go dependencies
print_status "Installing Go dependencies..."
if ! go mod tidy; then
    print_error "Failed to install Go dependencies"
    exit 1
fi
print_success "Go dependencies installed"

# Install frontend dependencies
print_status "Installing frontend dependencies..."
cd frontend
if ! npm install; then
    print_error "Failed to install frontend dependencies"
    exit 1
fi
cd ..
print_success "Frontend dependencies installed"

# Parse command line arguments
PLATFORM=""
MODE="production"

while [[ $# -gt 0 ]]; do
    case $1 in
        -p|--platform)
            PLATFORM="$2"
            shift 2
            ;;
        -m|--mode)
            MODE="$2"
            shift 2
            ;;
        -h|--help)
            echo "Usage: $0 [OPTIONS]"
            echo "Options:"
            echo "  -p, --platform PLATFORM    Target platform (windows/amd64, darwin/universal, linux/amd64)"
            echo "  -m, --mode MODE            Build mode (dev, production) [default: production]"
            echo "  -h, --help                 Show this help message"
            echo ""
            echo "Examples:"
            echo "  $0                         # Build for current platform"
            echo "  $0 -p windows/amd64        # Build for Windows 64-bit"
            echo "  $0 -p darwin/universal     # Build universal macOS binary"
            echo "  $0 -p linux/amd64         # Build for Linux 64-bit"
            echo "  $0 -m dev                  # Development build"
            exit 0
            ;;
        *)
            print_error "Unknown option: $1"
            exit 1
            ;;
    esac
done

# Build based on mode
if [ "$MODE" = "dev" ]; then
    print_status "Starting development mode..."
    wails dev
else
    print_status "Building for production..."
    
    if [ -z "$PLATFORM" ]; then
        print_status "Building for current platform..."
        wails build
    else
        print_status "Building for platform: $PLATFORM"
        wails build -platform "$PLATFORM"
    fi
    
    print_success "Build completed successfully!"
    print_status "You can find the executable in the build/bin/ directory"
fi

print_success "Mobile Forensics Toolkit build script completed! 🚀"