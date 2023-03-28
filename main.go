package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/docopt/docopt-go"
)

const version = "v0.1"

const usage = `A command-line tool for reporting and manipulating Unix timestamps.

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
`

type Config struct {
	Timestamp   int64
	AsLocaltime bool
	AsUTC       bool
	Modifier    string
}

func exitOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// Parse and validate command-line arguments
func getConfig() Config {

	var (
		retval  Config
		opts    docopt.Opts
		t       time.Time
		s       string
		dateStr string = time.Now().UTC().Format("2006-01-02")
		timeStr string = time.Now().UTC().Format("15:04:05")
		err     error
	)

	opts, err = docopt.ParseArgs(usage+" ", nil, version)
	exitOnError(err)

	retval.AsLocaltime, _ = opts.Bool("--localtime")
	retval.AsUTC, _ = opts.Bool("--utc")

	// Date
	if opts["--date"] != nil {
		dateStr, err = opts.String("--date")
		if _, err = time.Parse("2006-01-02", dateStr); err != nil {
			fmt.Fprintln(os.Stderr, "Unable to parse date:", dateStr)
			os.Exit(1)
		}

		if opts["--time"] != nil {
			timeStr, err = opts.String("--time")
			if _, err = time.Parse("15:04:05", timeStr); err != nil {
				fmt.Fprintln(os.Stderr, "Unable to parse time:", timeStr)
				os.Exit(1)
			}
		}
	}

	// Timestamp
	t, err = time.Parse("2006-01-0215:04:05", dateStr+timeStr) // default to "now"
	exitOnError(err)
	retval.Timestamp = t.Unix()
	if opts["--seconds"] != nil {
		s, err = opts.String("--seconds")
		if s != "" {
			retval.Timestamp, err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Unable to parse timestamp:", s)
				os.Exit(1)
			}
		}
	}

	// Modifier
	if opts["--modifier"] != nil {
		retval.Modifier, err = opts.String("--modifier")
	}

	return retval
}

func main() {
	var config = getConfig()
	var ut = NewUnixtime(config.Timestamp)

	if config.Modifier != "" {
		ut.Modify(config.Modifier)
	}

	if config.AsUTC {
		fmt.Println(ut.AsString(time.UTC))
	} else if config.AsLocaltime {
		fmt.Println(ut.AsString(time.Local))
	} else {
		fmt.Println(ut.timestamp)
	}
}
