# Ida

Ida is a voting assistant to help process voters for online music competitions. It is designed to work with contest spreadsheet files for [Melbourne](https://github.com/iune-melbourne/melbourne-archive).

## Usage

```
ida 0.2.0
Usage: ida COUNTRIES SPREADSHEET

Positional arguments:
  COUNTRIES              Path to JSON file with country information
  SPREADSHEET            Path to contest Excel spreadsheet

Options:
  --help, -h             display this help and exit
  --version              display version and exit
```

## Building from Source

To install you will first need Go installed on your computer. Instructions on how to download and install Go can be found on the [official website](https://golang.org/dl/).

Next, clone this repository and run the following commands. This will download any required dependencies, and build and install an executable for the program.

``` bash
# Download required dependencies
go get

# If cross-compiling, run the following script
./cross.sh

# Otherwise, simply run
go build
```
