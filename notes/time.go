package notes

import (
	"fmt"
	"time"
)

var week time.Duration

func TimeNotes() {
	t := time.Now()
	fmt.Println("time now is ", t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	week = 60 * 60 * 24 * 7 * 1e9
	weekForNow := t.Add(week)
	fmt.Println(weekForNow)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("21 Dec 2011 08:52"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
