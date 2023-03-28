# unixtime

A command-line tool for reporting and manipulating Unix timestamps.

## Installation

`unixtime` will run on most Linux and Mac OS X systems.

To install it, just find the appropriate one-liner below - based on the destination O/S and architecture - and copy-paste it into your terminal.

Feel free to change the install dir - `$HOME/bin` in the examples below - to be something more appropriate for your needs.

### Linux (32-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/unixtime/releases/latest/download/unixtime-linux-386.tar.gz | tar -zxf - -C $HOME/bin
```

### Linux (64-bit)

```
curl -s -L -o - https://github.com/alasdairmorris/unixtime/releases/latest/download/unixtime-linux-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Intel)

```
curl -s -L -o - https://github.com/alasdairmorris/unixtime/releases/latest/download/unixtime-darwin-amd64.tar.gz | tar -zxf - -C $HOME/bin
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o - https://github.com/alasdairmorris/unixtime/releases/latest/download/unixtime-darwin-arm64.tar.gz | tar -zxf - -C $HOME/bin
```

### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/unixtime@latest
```


## Usage

```
A command-line tool for reporting and manipulating Unix timestamps.

Usage:
  unixtime [-L | -U] [-s <timestamp>] [-m <modifier>]
  unixtime [-L | -U] [-d <date> [-t <time>]] [-m <modifier>]
  unixtime -h | --help
  unixtime --version

Options:
  -s, --seconds <ts>     The Unix timestamp to display/adjust [default is "now"]
  -d, --date <date>      The date to convert to a timestamp [default is "today"]
                         (format YYYY-MM-DD)
  -t, --time <time>      The time to convert to a timestamp [default is "now"]
                         (format HH:mm:ss)
  -m, --modifier <m>     String to describe how the time should be modified e.g.
                             "+1 hour" -> Add 1 hour
                             "-16 minutes" -> Subtract 16 minutes
                         Valid units are: second(s), minute(s), hour(s), day(s),
                         week(s), month(s), year(s)
  -L, --localtime        Format output as localtime.
  -U, --utc              Format output as UTC.

Homepage: https://github.com/alasdairmorris/unixtime

```

## Examples

```
$ unixtime
1680026055
```

```
$ unixtime -m "-30 days"
1677437662
```

```
$ unixtime -m "-30 days" --localtime
2023-02-26 18:55:19 GMT
```
