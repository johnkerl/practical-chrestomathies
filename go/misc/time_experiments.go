package main
import (
	"fmt"
	"time"
)

func main() {
	t    := time.Now()
	YYYY := t.Year()
	MM   := int(t.Month())
	DD   := t.Day()
	hh   := t.Hour()
	mm   := t.Minute()
	ss   := t.Second()
	nnn  := t.Nanosecond()

	_, offSecs := t.Zone()

	fmt.Printf("Default format: %s\n", t)
	fmt.Printf("\n")

	fmt.Printf("t.Unix():     %d\n", t.Unix())
	fmt.Printf("t.UnixNano(): %d\n", t.UnixNano())
	fmt.Printf("\n")

	fmt.Printf("YYYY: %04d\n", YYYY)
	fmt.Printf("MM  : %02d\n", MM)
	fmt.Printf("DD  : %02d\n", DD)
	fmt.Printf("hh  : %02d\n", hh)
	fmt.Printf("mm  : %02d\n", mm)
	fmt.Printf("ss  : %02d\n", ss)
	fmt.Printf("nnn : %09d\n", nnn)
	fmt.Printf("o  : %d\n", offSecs)
	fmt.Printf("\n")

	myFormats := []string {
		"1987-01-23",
		"12345678 12:34:56.000",
		"2010-06-14 hh:mm:45.000",
		"YYYYMMDD hh:mm:ss",
		"2013-05-17 23:21:49.876504 -0400",
		"2013-05-17 23:21:49.876504 -0400 EDT",
	}
	for _, myFormat := range(myFormats) {
		fmt.Printf("%-25s -> %s\n", myFormat, t.Format(myFormat))
	}
	fmt.Printf("%-25s -> %s\n", "[time.RFC3339Nano]",
		t.Format(time.RFC3339Nano))
	fmt.Printf("%-25s -> %s\n", "[t.String()]",
		t.Format(t.String()))

	offSign := "+"
	absOffHrMin := offSecs
	if offSecs < 0 {
		offSign = "-"
		absOffHrMin = -absOffHrMin
	}
	absOffHrMin /= 60
	whatIReallyWanted := fmt.Sprintf(
		"%04d%02d%02d %02d:%02d:%02d.%09d %s%02d%02d",
		YYYY, MM, DD, hh, mm, ss, nnn, offSign,
		absOffHrMin / 60, absOffHrMin % 60)
	fmt.Printf("\n")

	fmt.Printf("What I really wanted: %s\n", whatIReallyWanted)

//
//
//What I really wanted: 20130517 23:07:28.534816000
}
