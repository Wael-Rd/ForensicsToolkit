# Installation Guide - Mobile Forensics Toolkit

This guide provides detailed instructions for installing and setting up the Mobile Forensics Toolkit on different operating systems.

## System Requirements

### Minimum Requirements
- **CPU**: 64-bit processor, 2 GHz or faster
- **Memory**: 4 GB RAM minimum, 8 GB recommended
- **Storage**: 2 GB free disk space
- **OS**: Windows 10/11, macOS 10.15+, or Linux (Ubuntu 18.04+)

### Development Requirements
- **Go**: 1.19 or later
- **Node.js**: 16 or later
- **npm**: 8 or later
- **Git**: Latest version

## Pre-built Binaries (Recommended)

### Windows

1. **Download the latest release**
   - Visit [Releases](https://github.com/yourusername/MobileForensicsToolkit/releases)
   - Download `MobileForensicsToolkit-windows-amd64.exe`

2. **Install**
   - Run the installer as Administrator
   - Follow the installation wizard
   - Accept the license agreement
   - Choose installation directory

3. **First Run**
   - Launch from Start Menu or Desktop shortcut
   - Windows Defender may show a warning (click "More info" → "Run anyway")
   - Grant necessary permissions when prompted

### macOS

1. **Download and Install**
   - Download `MobileForensicsToolkit-darwin-universal.dmg`
   - Open the DMG file
   - Drag the application to Applications folder

2. **Security Settings**
   - First launch: Right-click → "Open" to bypass Gatekeeper
   - Or go to System Preferences → Security & Privacy → Allow

3. **Verify Installation**
   - Launch from Applications folder
   - Check version: Menu → About

### Linux

#### Ubuntu/Debian
```bash
# Download DEB package
wget https://github.com/yourusername/MobileForensicsToolkit/releases/latest/download/MobileForensicsToolkit-linux-amd64.deb

# Install
sudo dpkg -i MobileForensicsToolkit-linux-amd64.deb
sudo apt-get install -f  # Fix dependencies if needed

# Run
MobileForensicsToolkit
```

#### CentOS/RHEL/Fedora
```bash
# Download RPM package
wget https://github.com/yourusername/MobileForensicsToolkit/releases/latest/download/MobileForensicsToolkit-linux-amd64.rpm

# Install
sudo rpm -i MobileForensicsToolkit-linux-amd64.rpm
# Or with yum/dnf
sudo yum install MobileForensicsToolkit-linux-amd64.rpm

# Run
MobileForensicsToolkit
```

#### Generic Linux (AppImage)
```bash
# Download AppImage
wget https://github.com/yourusername/MobileForensicsToolkit/releases/latest/download/MobileForensicsToolkit-linux-amd64.AppImage

# Make executable
chmod +x MobileForensicsToolkit-linux-amd64.AppImage

# Run
./MobileForensicsToolkit-linux-amd64.AppImage
```

## Building from Source

### Prerequisites Installation

#### Install Go

**Windows:**
1. Download from https://golang.org/dl/
2. Run the installer
3. Verify: `go version`

**macOS:**
```bash
# Using Homebrew
brew install go

# Or download from https://golang.org/dl/
```

**Linux:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# CentOS/RHEL/Fedora
sudo yum install golang
# or
sudo dnf install golang

# Verify
go version
```

#### Install Node.js

**Windows/macOS:**
- Download from https://nodejs.org/
- Run the installer

**Linux:**
```bash
# Ubuntu/Debian
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# CentOS/RHEL/Fedora
curl -fsSL https://rpm.nodesource.com/setup_18.x | sudo bash -
sudo yum install -y nodejs

# Verify
node --version
npm --version
```

#### Install Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Verify
wails version
```

### Building the Application

1. **Clone Repository**
   ```bash
   git clone https://github.com/yourusername/MobileForensicsToolkit.git
   cd MobileForensicsToolkit
   ```

2. **Install Dependencies**
   ```bash
   # Backend dependencies
   go mod tidy
   
   # Frontend dependencies
   cd frontend
   npm install
   cd ..
   ```

3. **Development Build**
   ```bash
   # Run in development mode
   wails dev
   ```

4. **Production Build**
   ```bash
   # Build for current platform
   wails build
   
   # Build for specific platform
   wails build -platform windows/amd64
   wails build -platform darwin/universal
   wails build -platform linux/amd64
   ```

### Build Outputs

Built applications will be in the `build/bin/` directory:
- **Windows**: `MobileForensicsToolkit.exe`
- **macOS**: `MobileForensicsToolkit.app`
- **Linux**: `MobileForensicsToolkit`

## Troubleshooting Installation

### Common Issues

#### Windows

**Issue**: "Windows protected your PC" warning
- **Solution**: Click "More info" → "Run anyway"
- **Alternative**: Right-click installer → Properties → Unblock

**Issue**: Missing Visual C++ Redistributable
- **Solution**: Install Microsoft Visual C++ Redistributable

#### macOS

**Issue**: "App can't be opened because it is from an unidentified developer"
- **Solution**: System Preferences → Security & Privacy → Allow
- **Alternative**: Right-click app → Open

**Issue**: Wails command not found
- **Solution**: Add Go bin to PATH: `export PATH=$PATH:$(go env GOPATH)/bin`

#### Linux

**Issue**: Missing dependencies
```bash
# Ubuntu/Debian
sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev

# CentOS/RHEL/Fedora
sudo yum install gtk3-devel webkit2gtk3-devel
```

**Issue**: Permission denied
```bash
# Make executable
chmod +x MobileForensicsToolkit

# Or run AppImage
chmod +x MobileForensicsToolkit-linux-amd64.AppImage
```

### Build Issues

**Issue**: Go modules error
```bash
# Clean module cache
go clean -modcache
go mod download
go mod tidy
```

**Issue**: Node.js dependencies error
```bash
# Clear npm cache
npm cache clean --force
rm -rf frontend/node_modules
cd frontend && npm install
```

**Issue**: Wails build fails
```bash
# Update Wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Check doctor
wails doctor
```

## Verification

After installation, verify the application works:

1. **Launch Application**
   - Check the application starts without errors
   - Verify the UI loads correctly

2. **Test Core Features**
   - Try the timestamp converter (no files needed)
   - Check if all menu items are accessible
   - Verify help and about dialogs work

3. **Check Logs**
   - Windows: `%APPDATA%\MobileForensicsToolkit\logs\`
   - macOS: `~/Library/Application Support/MobileForensicsToolkit/logs/`
   - Linux: `~/.config/MobileForensicsToolkit/logs/`

## Next Steps

After successful installation:
1. Read the [User Guide](USER_GUIDE.md)
2. Review [Legal Guidelines](../README.md#legal--ethical-use)
3. Join the community discussions
4. Report any issues on GitHub

## Support

If you encounter issues:
1. Check the [Troubleshooting Guide](TROUBLESHOOTING.md)
2. Search [existing issues](https://github.com/yourusername/MobileForensicsToolkit/issues)
3. Create a new issue with detailed information
4. Join community discussions for help

Remember: This tool is for legitimate forensics investigations only. Ensure you have proper authorization before analyzing any device or data.