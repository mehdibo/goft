# goft
![Tests](https://github.com/mehdibo/goft/workflows/Tests/badge.svg?branch=develop)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mehdibo/goft)
[![GitHub license](https://img.shields.io/github/license/mehdibo/goft)](https://github.com/mehdibo/goft/blob/develop/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/mehdibo/goft)](https://github.com/mehdibo/goft/issues)
![GitHub all releases](https://img.shields.io/github/downloads/mehdibo/goft/total)


A CLI tool to use 42's API

# Usage

Clone the repository with:
```bash
git clone https://github.com/mehdibo/goft.git
cd goft
```

### Locally
To use the binary locally, just compile the binary
```bash
make
```

It will create a binary with the name `goft` and voila!


### Globally
If you want to use goft globally just run
```bash
make install
```

goft is available wherever you go!


To uninstall it, simply run `make uninstall`

### Configuration

You need to configure your oAuth credentials in a *.yaml* file
If you ran `make install`, it created a `.goft.yaml` in your home folder.

Edit that with the appropriate values
