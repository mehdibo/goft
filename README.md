# goft
![Tests](https://github.com/mehdibo/goft/workflows/Tests/badge.svg?branch=develop)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mehdibo/goft)
[![GitHub license](https://img.shields.io/github/license/mehdibo/goft)](https://github.com/mehdibo/goft/blob/develop/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/mehdibo/goft)](https://github.com/mehdibo/goft/issues)
![GitHub all releases](https://img.shields.io/github/downloads/mehdibo/goft/total)


Goft makes it easier to interact with 42's API without having to deal with the authentication
nor all the repetitive parsing and validation.

The tool comes with built in commands to do common tasks but also makes it easier to send raw requests.

 * [Installation](#installation)
   * [Install script](#install-script)
   * [Compiling sources](#compiling-sources)
 * [Configuration](#configuration)

## Installation

### Install script
You can use the provided script to easily install the latest version, just run:
```shell
curl https://raw.githubusercontent.com/mehdibo/goft/main/install.sh | bash
```
This script will download the binary and make it available globally, it will also download a sample config file
if it doesn't already exist.
### Compiling sources
If you want to compile from the source code, you need to have Golang installed in your machine.

Clone the repository:
```shell
git clone https://github.com/mehdibo/goft.git
```

Run:
```shell
# Or make all_platforms
# To compile for all platforms
make

# You can also use this to install the binary globally
make install
```


## Configuration
After installing Goft you will have to modify `~/.goft.yml` with your credentials.

This file makes it easier to run commands without having to pass the `--config` flag every time.
