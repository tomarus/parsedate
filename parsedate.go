package parsedate
//
// Parsedate function, original idea by:
// https://groups.google.com/d/msg/golang-nuts/mc0VmhNajn0/rmVjvfNNV1sJ
//
// These dates are sampled from about 1 billion usenet articles.
// It should match most of user supplied data formats.
//
// tommy@chiparus.org 2014
// This is free public domain software, please copy.
//

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type formatInfo struct {
	format string
	needed string // all the missing information from the format
}

var formats = []formatInfo{
	{time.RFC1123, ""},
	{time.RFC1123Z, ""},
	{time.RFC3339, ""},
	{time.RFC3339Nano, ""},
	{"02 Jan 06 15:04:05", ""},
	{"02 Jan 06 15:04:05 +-0700", ""},
	{"02 Jan 06 15:4:5 MST", ""},
	{"02 Jan 2006 15:04:05", ""},
	{"2 Jan 2006 15:04:05", ""},
	{"2 Jan 2006 15:04:05 MST", ""},
	{"2 Jan 2006 15:04:05 -0700", ""},
	{"2 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"02 January 2006 15:04", ""},
	{"02 Jan 2006 15:04 MST", ""},
	{"02 Jan 2006 15:04:05 MST", ""},
	{"02 Jan 2006 15:04:05 -0700", ""},
	{"02 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"Mon, 2 Jan  15:04:05 MST 2006", ""},
	{"Mon, 2 Jan 15:04:05 MST 2006", ""},
	{"Mon, 02 Jan 2006 15:04:05", ""},
	{"Mon, 02 Jan 2006 15:04:05 (MST)", ""},
	{"Mon, 2 Jan 2006 15:04:05", ""},
	{"Mon, 2 Jan 2006 15:04:05 MST", ""},
	{"Mon, 2 Jan 2006 15:04:05 -0700", ""},
	{"Mon, 2 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"Mon, 02 Jan 06 15:04:05 MST", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 MST", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 (MST)", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 (MST-07:00)", ""},
	{"Mon, 02 Jan 2006 15:04:05 -0700 (MST MST)", ""},
	{"Mon, 02 Jan 2006 15:04 -0700", ""},
	{"Mon, 02 Jan 2006 15:04 -0700 (MST)", ""},
	{"Mon Jan 02 15:05:05 2006 MST", ""},
	{"Monday, 02 Jan 2006 15:04 -0700", ""},
	{"Monday, 02 Jan 2006 15:04:05 -0700", ""},
}

func Parse(s string) (time.Time, error) {
	for _, f := range formats {
		format := f.format
		t, err := time.Parse(format, s)
		if err != nil {
			continue
		}
		t, err = time.Parse(f.format+f.needed, s+time.Now().Format(f.needed))
		if err != nil {
			panic("unexpectedly failed parse")
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("Failed to parse time: %s", s)
}

func maintest() {
	reader := bufio.NewReader(os.Stdin)

	for {
		str, _, err := reader.ReadLine()
		if err == io.EOF {
			return
		}

		//t, err := parse(string(str), time.Now())
		_, err = Parse(string(str))
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			//fmt.Printf("%v\n", t)
		}
	}
}
