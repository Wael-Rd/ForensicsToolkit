# Contributing to Mobile Forensics Toolkit

We welcome contributions from the digital forensics community! This document outlines the process for contributing to the Mobile Forensics Toolkit.

## Code of Conduct

By participating in this project, you are expected to uphold our Code of Conduct:
- Be respectful and inclusive
- Focus on what is best for the community
- Show empathy towards other community members
- Use welcoming and inclusive language

## How to Contribute

### Reporting Bugs

1. **Check existing issues** to avoid duplicates
2. **Use the bug report template** when creating new issues
3. **Provide detailed information**:
   - Operating system and version
   - Go version
   - Steps to reproduce the issue
   - Expected vs actual behavior
   - Screenshots or logs if applicable

### Requesting Features

1. **Check existing feature requests** to avoid duplicates
2. **Use the feature request template**
3. **Clearly describe the use case** and why it would be valuable
4. **Consider implementation complexity** and potential security implications

### Contributing Code

#### Prerequisites

- Go 1.19 or later
- Node.js 16 or later
- Wails CLI v2
- Git

#### Development Setup

1. **Fork the repository**
   ```bash
   git clone https://github.com/yourusername/MobileForensicsToolkit.git
   cd MobileForensicsToolkit
   ```

2. **Install dependencies**
   ```bash
   cd frontend && npm install && cd ..
   go mod tidy
   ```

3. **Run in development mode**
   ```bash
   wails dev
   ```

#### Making Changes

1. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Follow coding standards**:
   - Use `gofmt` for Go code formatting
   - Follow Vue.js style guide for frontend code
   - Add comments for complex logic
   - Write unit tests for new functionality

3. **Test your changes**
   ```bash
   go test ./...
   cd frontend && npm test
   ```

4. **Commit with clear messages**
   ```bash
   git commit -m "feat: add support for new database format"
   ```

#### Pull Request Process

1. **Update documentation** if needed
2. **Ensure tests pass** and add new tests for your changes
3. **Update CHANGELOG.md** with your changes
4. **Create a pull request** with:
   - Clear title and description
   - Reference to related issues
   - Screenshots or demos if applicable

### Security Considerations

- **Never commit sensitive data** (passwords, keys, database files)
- **Review code for security vulnerabilities**
- **Follow secure coding practices**
- **Report security issues privately** to the maintainers

## Development Guidelines

### Backend (Go)

- Use meaningful variable and function names
- Add error handling for all operations
- Follow Go conventions and idioms
- Add unit tests for new functionality
- Document exported functions and types

### Frontend (Vue.js)

- Use composition API for new components
- Follow TDesign component guidelines
- Implement proper error handling
- Add TypeScript types where beneficial
- Ensure responsive design

### Documentation

- Update README.md for major changes
- Add inline code comments for complex logic
- Create or update user guides as needed
- Include examples in documentation

## Legal and Ethical Guidelines

- **Ensure contributions are ethical** and support legitimate forensics use
- **Respect intellectual property** rights
- **Do not contribute code** that facilitates illegal activities
- **Follow responsible disclosure** for security vulnerabilities

## Recognition

Contributors will be recognized in:
- CONTRIBUTORS.md file
- Release notes for significant contributions
- GitHub contributors section

## Getting Help

- **Discord/Slack**: Join our community chat
- **GitHub Discussions**: For general questions and ideas
- **GitHub Issues**: For bug reports and feature requests
- **Email**: Contact maintainers directly for sensitive issues

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

Thank you for contributing to the Mobile Forensics Toolkit! 🚀