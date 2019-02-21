# Ida

Ida is a voting assistant to help process voters for online music competitions. It is designed to work with contest spreadsheet files for [Melbourne](https://github.com/iune-melbourne/melbourne-archive).

## Usage

```
ida 0.1.0
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

Next, clone this repository and run the following two commands. This will download any required dependencies, build and install an executable for the program.

```
go get
go install
```