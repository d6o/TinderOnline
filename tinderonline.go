package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/disiqueira/tindergo"
)

func main() {
	token := flag.String("token", "", "Your Facebook Token.")

	flag.Parse()

	if *token == "" {
		fmt.Println("You must provide a valid Facebook Token.")
		os.Exit(2)
	}

	t := tindergo.New()

	err := t.Authenticate(*token)
	checkError(err)

	profile, err := t.Profile()
	checkError(err)

	fmt.Println("Your Profile:")
	fmt.Println("Name: " + profile.Name)
	fmt.Println("")

	friends, err := t.Friends()
	now := time.Now()

	fmt.Printf("|%40s|%70s|\n", "Your Friend", "Last Time Online")
	fmt.Printf("|%40s|%70s|\n", "", "")
	for _, e := range friends {
		user, err := t.User(e.UserID)
		checkError(err)
		fmt.Printf("|%40s|%70s|\n", e.Name, diffForHumans(user.PingTime, now))
	}
}

func diffForHumans(a, b time.Time) string {
	year, month, day, hour, min, sec := diff(a, b)
	s := ""
	p := false

	if year > 0 {
		s += fmt.Sprintf("%02d", year) + " year(s) "
		p = true
	}

	if month > 0 || p {
		s += fmt.Sprintf("%02d", month) + " month(s) "
		p = true
	}

	if day > 0 || p {
		s += fmt.Sprintf("%02d", day) + " day(s) "
		p = true
	}

	if hour > 0 || p {
		s += fmt.Sprintf("%02d", hour) + " hour(s) "
		p = true
	}

	if min > 0 || p {
		s += fmt.Sprintf("%02d", min) + " min(s) "
		p = true
	}

	if !p {
		return "ONLINE NOW!!!"
	}

	if sec > 0 || p {
		s += fmt.Sprintf("%02d", sec) + " sec(s) "
		p = true
	}

	return s + "ago"
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
