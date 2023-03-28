package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Unixtime struct {
	timestamp int64
}

func NewUnixtime(_timestamp int64) *Unixtime {
	return &Unixtime{_timestamp}
}

func (u *Unixtime) AsString(loc *time.Location) string {
	tm := time.Unix(u.timestamp, 0)
	return tm.In(loc).Format("2006-01-02 15:04:05 MST")
}

func (u *Unixtime) Modify(modifier string) {
	var seconds = regexp.MustCompile(`^[\+\-]\d+ seconds?$`)
	var minutes = regexp.MustCompile(`^[\+\-]\d+ minutes?$`)
	var hours = regexp.MustCompile(`^[\+\-]\d+ hours?$`)
	var days = regexp.MustCompile(`^[\+\-]\d+ days?$`)
	var weeks = regexp.MustCompile(`^[\+\-]\d+ weeks?$`)
	var months = regexp.MustCompile(`^[\+\-]\d+ months?$`)
	var years = regexp.MustCompile(`^[\+\-]\d+ years?$`)

	var tm = time.Unix(u.timestamp, 0)

	modifier = strings.TrimSpace(modifier)

	switch {
	case seconds.MatchString(modifier):
		s := strings.Fields(modifier)
		i, _ := strconv.Atoi(s[0])
		tm = tm.Add(time.Duration(i) * time.Second)
	case minutes.MatchString(modifier):
		s := strings.Fields(modifier)
		i, _ := strconv.Atoi(s[0])
		tm = tm.Add(time.Duration(i) * time.Minute)
	case hours.MatchString(modifier):
		s := strings.Fields(modifier)
		i, _ := strconv.Atoi(s[0])
		tm = tm.Add(time.Duration(i) * time.Hour)
	case days.MatchString(modifier):
		s := strings.Fields(modifier)
		i, _ := strconv.Atoi(s[0])
		tm = tm.AddDate(0, 0, i)
	case weeks.MatchString(modifier):
		s := strings.Fields(modifier)
		i, _ := strconv.Atoi(s[0])
		tm = tm.AddDate(0, 0, 7*i)
	case months.MatchString(modifier):
		s := strings.Fields(modifier)
		i, _ := strconv.Atoi(s[0])
		tm = tm.AddDate(0, i, 0)
	case years.MatchString(modifier):
		s := strings.Fields(modifier)
		i, _ := strconv.Atoi(s[0])
		tm = tm.AddDate(i, 0, 0)
	default:
		fmt.Fprintln(os.Stderr, "Unable to parse modifier:", modifier)
		os.Exit(1)
	}

	u.timestamp = tm.Unix()
}
