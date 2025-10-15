# Changelog

All notable changes to the Mobile Forensics Toolkit will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release preparation
- Documentation improvements
- Build system enhancements

## [2.0.0] - 2025-01-16

### Added
- **Complete English Translation**: Rebuilt entire codebase with professional English interface
- **Modern UI Framework**: Upgraded to Vue 3 + TDesign components
- **Wails v2 Integration**: Desktop application framework for better performance
- **Comprehensive Documentation**: Professional README, setup guides, and API documentation
- **Enhanced Database Support**:
  - SQLCipher 3/4 encryption schemes
  - WeChat database decryption with multiple algorithms
  - XiaoHongShu static key decryption
  - MosGram/Bubble MD5-based decryption
  - WuKong IM series UserID-based decryption
- **Advanced Password Calculation**:
  - WeChat IMEI-based password generation
  - TikTok/Douyin dynamic password algorithms
  - Custom encryption scheme support
- **Data Extraction Tools**:
  - Message recovery with timeline reconstruction
  - Contact list and relationship graph analysis
  - Media extraction (images, videos, voice)
  - Metadata preservation and analysis
- **Registry Analysis**:
  - Windows registry forensic artifact extraction
  - System-level evidence collection
  - User activity timeline reconstruction
- **Security Tools**:
  - Smart dictionary-based brute force engine
  - Timestamp format conversion utilities
  - Cryptographic hash calculation and verification
  - File structure analysis tools

### Changed
- **Project Structure**: Reorganized codebase for better maintainability
- **Build System**: Improved build scripts with cross-platform support
- **Error Handling**: Enhanced error messages and user feedback
- **Performance**: Optimized database operations and UI responsiveness

### Security
- **Code Review**: Complete security audit of all modules
- **Input Validation**: Enhanced validation for all user inputs
- **Memory Safety**: Improved memory management in Go modules
- **Dependency Updates**: Updated all dependencies to latest secure versions

### Documentation
- **User Manual**: Comprehensive user guide with screenshots
- **API Reference**: Complete API documentation for developers
- **Legal Guidelines**: Clear usage guidelines and disclaimers
- **Contributing Guide**: Detailed contribution guidelines for community

### Technical Improvements
- **Go Modules**: Modern Go module system implementation
- **TypeScript Support**: Enhanced frontend type safety
- **Test Coverage**: Expanded unit test coverage
- **CI/CD Pipeline**: Automated testing and build processes

---

## Previous Versions

### [1.x.x] - Legacy Chinese Version
- Original Chinese implementation
- Basic forensic analysis capabilities
- Limited documentation and support

---

## Version Notes

### Supported Platforms
- **Windows**: Windows 10/11 (x64)
- **macOS**: macOS 10.15+ (Intel/Apple Silicon)
- **Linux**: Ubuntu 18.04+, CentOS 7+, Debian 10+

### Dependencies
- **Go**: 1.19 or later
- **Node.js**: 16 or later
- **Wails**: v2.9.0 or later

### Breaking Changes
- Complete rewrite from Chinese to English
- New module structure and import paths
- Updated configuration file formats
- Modified API endpoints and function signatures

### Migration Guide
Users migrating from the Chinese version should:
1. Backup existing configuration and data
2. Uninstall previous version completely
3. Install new English version following setup guide
4. Reconfigure settings and preferences
5. Import existing case data if compatible

---

**Note**: This changelog reflects the transformation from the original Chinese forensics tool to the professional English Mobile Forensics Toolkit.