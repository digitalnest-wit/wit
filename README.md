# wit 💻

## Project Description

`wit` is a powerful command-line interface (CLI) tool designed specifically for Digital NEST to streamline and automate common computer operations. This utility aims to simplify IT management and enhance productivity for the organization's technical staff.

## Features

- **Software Deployment**: Simplify software installation process
- **System Management**: Quickly perform system-level operations

## Installation

### Prerequisites
- Administrative/sudo access
- Go (version 1.20 or higher) - Required for building from source
- Download and install from the [official Go website](https://golang.org/dl/)
- Verify installation with `go version`
- Git
- Basic command-line knowledge

### Install from Source
```bash
# Clone the repository
git clone https://github.com/digitalnest/wit.git
cd wit

# Verify Go installation
go version

# Build the project
go build -o wit

# Optional: Install the binary system-wide
sudo mv wit /usr/local/bin/
```

## Usage

```bash
# Basic command structure
wit [command] [options]

# Example commands
wit help   # Display usage and help for specific commands
wit config # Install Homebrew software and VS Code extensions via a configuration file
```

## Contributing

Contributions are welcome! Please see `CONTRIBUTING.md` for guidelines on submitting pull requests.

## License

This project is licensed under the MIT License. See `LICENSE` file for details.

## Support

For issues and questions, please open a GitHub issue or contact the Digital NEST IT team.

## Changelog

See `CHANGELOG.md` for version history and updates.
