package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// time layout constants named by the most precise granularity afforded
const (
	Nanoseconds  = "2006-01-02T15:04:05.999999999Z07:00"
	Microseconds = "2006-01-02T15:04:05.999999Z07:00"
	Milliseconds = "2006-01-02T15:04:05.999Z07:00"
	Seconds      = "2006-01-02T15:04:05Z07:00"
	Minutes      = "2006-01-02T15:04Z07:00"
	Date         = "2006-01-02" // "2006-01-02Z07:00"
)

func timeLayout(nanoseconds, microseconds, milliseconds, seconds, minutes, date bool) string {
	if nanoseconds {
		return Nanoseconds
	} else if microseconds {
		return Microseconds
	} else if milliseconds {
		return Milliseconds
	} else if seconds {
		return Seconds
	} else if minutes {
		return Minutes
	} else if date {
		return Date
	}
	// default to seconds
	return Seconds
}

func epochMultiplier(nanoseconds, microseconds, milliseconds, seconds, minutes, date bool) int64 {
	if nanoseconds {
		return 1
	} else if microseconds {
		return 1000
	} else if milliseconds {
		return 1000000
	}
	// minutes and date are ignored; default to seconds
	return 1000000000
}

func formatTime(t time.Time, epoch, nanoseconds, microseconds, milliseconds, seconds, minutes, date bool) string {
	if epoch {
		multiplier := epochMultiplier(nanoseconds, microseconds, milliseconds, seconds, minutes, date)
		ticks := float64(t.UnixNano()) / float64(multiplier)
		return fmt.Sprintf("%.0f", ticks)
	} else {
		layout := timeLayout(nanoseconds, microseconds, milliseconds, seconds, minutes, date)
		return t.Format(layout)
	}
}

func main() {
	now := time.Now()
	env_now := os.Getenv("NOW")
	if env_now != "" {
		env_now_time, err := time.Parse(Nanoseconds, env_now)
		if err != nil {
			fmt.Println("Error in NOW value;", err)
			os.Exit(65) // EX_DATAERR
		}
		now = env_now_time.Local()
	}

	var epoch, d, m, s, ms, ns, local bool
	flag.BoolVar(&epoch, "epoch", false, "epoch ticks")
	// if more than one of -d|-m|-s|-ms|-ns are given, use the most precise
	flag.BoolVar(&d, "d", false, "date only")
	flag.BoolVar(&m, "m", false, "up to minutes")
	flag.BoolVar(&s, "s", false, "up to seconds")
	flag.BoolVar(&ms, "ms", false, "up to milliseconds")
	flag.BoolVar(&ns, "ns", false, "up to nanoseconds")
	// -local has no effect when -epoch is also used
	flag.BoolVar(&local, "local", false, "local time")
	flag.Parse()

	if !local {
		now = now.UTC()
	}

	now_string := formatTime(now, epoch, ns, false, ms, s, m, d)
	fmt.Println(now_string)
}
