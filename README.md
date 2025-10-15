# Mobile Forensics Toolkit

<div align="center">

![Version](https://img.shields.io/badge/version-2.0.0-blue.svg)
![Go](https://img.shields.io/badge/go-1.19+-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg)

A comprehensive digital forensics toolkit for mobile application database analysis, password recovery, and data extraction.

</div>

## 🔍 Overview

Mobile Forensics Toolkit is a professional-grade digital forensics application built with Go and Wails v2. It provides investigators with powerful tools for analyzing encrypted mobile application databases, recovering passwords, and extracting digital evidence from various messaging and social media platforms.

## ✨ Key Features

### 🔐 Database Decryption
- **SQLCipher Support**: SQLCipher 3/4 database decryption
- **WeChat Analysis**: Decrypt WeChat databases with various encryption schemes
- **Multi-Platform**: Support for XiaoHongShu, MosGram/Bubble, WuKong IM series
- **Automated Detection**: Intelligent database type detection and parameter selection

### 🔑 Password Calculation
- **WeChat Password Generation**: Calculate WeChat database passwords from device IMEI
- **TikTok/Douyin**: Generate passwords for TikTok and Douyin databases
- **Custom Algorithms**: Support for various proprietary password generation schemes
- **Bulk Processing**: Process multiple accounts simultaneously

### 📱 Data Extraction
- **Message Recovery**: Extract chat messages, media, and metadata
- **Contact Analysis**: Recover contact lists and relationship graphs
- **Timeline Analysis**: Reconstruct communication timelines
- **Media Extraction**: Recover images, videos, and voice messages

### 🖥️ Registry Analysis
- **Windows Registry**: Analyze Windows registry entries for digital evidence
- **System Artifacts**: Extract system-level forensic artifacts
- **User Activity**: Track user behavior and application usage

### 🔨 Additional Tools
- **Brute Force Engine**: Password cracking utilities with smart dictionary attacks
- **Timestamp Converter**: Convert between various timestamp formats
- **Hash Calculator**: Generate and verify cryptographic hashes
- **File Analysis**: Analyze file structures and metadata

## 🎯 Supported Applications

| Application | Database Type | Encryption | Status |
|-------------|---------------|------------|--------|
| WeChat (微信) | SQLCipher | AES-256 | ✅ Full Support |
| XiaoHongShu (小红书) | SQLCipher3 | Static Key | ✅ Full Support |
| TikTok/Douyin | Custom | Dynamic | ✅ Full Support |
| MosGram/Bubble | SQLCipher4 | MD5-based | ✅ Full Support |
| WuKong IM | SQLCipher4 | UserID-based | ✅ Full Support |
| Telegram | Custom | Various | 🚧 Partial Support |
| WhatsApp | SQLite/Encrypted | Crypt12/14 | 🚧 In Development |

## 🚀 Quick Start

### Prerequisites

- **Go 1.19+** - [Download](https://golang.org/dl/)
- **Node.js 16+** - [Download](https://nodejs.org/)
- **Wails CLI v2** - Install with: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/MobileForensicsToolkit.git
   cd MobileForensicsToolkit
   ```

2. **Install dependencies**
   ```bash
   # Install frontend dependencies
   cd frontend
   npm install
   cd ..
   
   # Install Go dependencies
   go mod tidy
   ```

3. **Development mode**
   ```bash
   wails dev
   ```

4. **Production build**
   ```bash
   wails build
   ```

### Using the Application

1. **Launch** the application
2. **Select a tool** from the sidebar navigation
3. **Upload target files** (databases, registry files, etc.)
4. **Configure parameters** as needed
5. **Run analysis** and review results
6. **Export findings** in various formats

## 🛠️ Technical Architecture

```
MobileForensicsToolkit/
├── analyzers/           # Core analysis modules
│   ├── database/       # Database decryption engines
│   ├── extractor/      # Data extraction tools
│   ├── passwdCalc/     # Password calculation algorithms
│   └── winreg/         # Windows registry analysis
├── tools/              # Utility tools
│   ├── cracker/        # Brute force engines
│   ├── timestamp/      # Time conversion utilities
│   └── hash/           # Cryptographic tools
├── frontend/           # Vue.js user interface
│   ├── src/
│   ├── components/
│   └── views/
├── utils/              # Shared utilities
└── docs/               # Documentation
```

### Technology Stack

- **Backend**: Go 1.19+ with Wails v2 framework
- **Frontend**: Vue 3 + TDesign UI components
- **Build Tool**: Vite for fast development
- **Encryption**: SQLCipher, AES, RC4 support
- **Database**: SQLite with various encryption schemes

## 📖 Documentation

- [Installation Guide](docs/INSTALLATION.md)
- [User Manual](docs/USER_GUIDE.md)
- [API Reference](docs/API.md)
- [Contributing Guidelines](CONTRIBUTING.md)
- [Troubleshooting](docs/TROUBLESHOOTING.md)

## 🔒 Legal & Ethical Use

**⚠️ IMPORTANT DISCLAIMER**

This toolkit is designed for legitimate digital forensics investigations and research purposes only. Users must:

- ✅ Have proper legal authorization before analyzing any device or data
- ✅ Comply with local laws and regulations regarding digital forensics
- ✅ Respect privacy rights and data protection laws
- ✅ Use only on devices/data you own or have explicit permission to analyze
- ❌ Not use for unauthorized access to systems or data
- ❌ Not use for illegal surveillance or privacy violations

**Users are solely responsible for ensuring their use complies with applicable laws.**

## 🤝 Contributing

We welcome contributions from the digital forensics community!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 🐛 Bug Reports & Feature Requests

Please use the [GitHub Issues](https://github.com/yourusername/MobileForensicsToolkit/issues) page to:
- Report bugs with detailed reproduction steps
- Request new features or improvements
- Ask questions about usage or implementation

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **SQLCipher Team** for the excellent database encryption library
- **Wails Framework** for enabling Go-based desktop applications
- **Digital Forensics Community** for continuous research and development
- **JetBrains** for supporting open source development

## 📊 Project Stats

- **Language**: Go (Backend), Vue.js (Frontend)
- **Lines of Code**: ~15,000+
- **Supported Platforms**: Windows, macOS, Linux
- **Active Development**: Yes
- **First Release**: 2025

---

<div align="center">

**Made with ❤️ for the Digital Forensics Community**

[🌟 Star this repo](https://github.com/yourusername/MobileForensicsToolkit) | [🔧 Report Issues](https://github.com/yourusername/MobileForensicsToolkit/issues) | [📚 Documentation](docs/)

</div>