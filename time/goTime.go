package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2019, 11, 20, 13, 44, 56, 5564321, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p("+1:", now.Add(1)) //1 naos
	p(now.Add(-32))

	secs := now.Unix()
	nanoSecs := now.UnixNano()
	p(secs)
	p(nanoSecs)
	millSecs := nanoSecs / 1000000
	p(millSecs)

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanoSecs))

	p(now.Format(time.RFC3339))
	t, e := time.Parse(time.RFC3339, "2019-04-01T16:33:42+08:00")
	p(t, e)

	p(t.Format("04:05PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "04:05PM")
	p(e)
	_, e = time.Parse(ansic, "Mon Jan 2 03:04:05 2006")
	p(e)
}
