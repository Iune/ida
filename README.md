# Ida

Ida is a voting assistant to help process voters for online music competitions. It is designed to work with contest spreadsheet files for [Melbourne](https://github.com/iune-melbourne/melbourne-archive).

## Usage

```
ida 0.1.1
Usage: ida COUNTRIES SPREADSHEET

Positional arguments:
  COUNTRIES              Path to tab-separated file with country information
  SPREADSHEET            Path to tab-separated contest file

Options:
  --help, -h             display this help and exit
  --version              display version and exit
```

## Building from Source

To install you will first need Go installed on your computer. Instructions on how to download and install Go can be found on the [official website](https://golang.org/dl/).

Next, clone this repository and run the following commands. This will download any required dependencies, and build and install an executable for the program.

```
# Download required dependencies
go get

# If cross-compiling on a platform other than Windows, run the following two commands
github.com/konsorten/go-windows-terminal-sequences
go get github.com/mattn/go-colorable

# If cross-compiling, run the following script
./cross.sh

# Otherwise, simply run
go build
```

**Note:** By default, `go get` does not download the two packages listed above on platforms other than Windows. In order to cross compile on platforms other than Windows, these two packages must be explicitly downloaded as shown above.