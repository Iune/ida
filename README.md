# Ida

Ida is a voting assistant to help process voters for online music competitions. It is designed to work with contest
files for [Melbourne](https://github.com/Iune/melbourne).

## Installation

```
dotnet publish -c Release -r osx-x64 --self-contained false 
```

The resulting build is located at `Ida/bin/Release/net5.0/osx-x64/publish/`.

## Usage

```
USAGE: ida [--help] --contest <path>

OPTIONS:

    --contest <path>      JSON file containing contest details
    --help                display this list of options.
```
