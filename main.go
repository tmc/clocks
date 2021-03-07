package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

var defaultTimeFormatString = "Mon Jan 2 15:04 MST"

var (
	flagFormatString = flag.String("fmt", defaultTimeFormatString, "time format string (golang time pkg)")
	flagTimezones    = flag.String("tzs", "UTC,America/Los_Angeles", "time zones")
)

func main() {
	flag.Parse()
	runtime.LockOSThread()

	app := cocoa.NSApp_WithDidLaunch(nsApp)
	app.Run()
}

func nsApp(n objc.Object) {
	obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
	obj.Retain()
	obj.Button().SetTitle("")

	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
			}
			t := time.Now()
			parts := []string{}

			for _, tz := range strings.Split(*flagTimezones, ",") {
				loc, err := time.LoadLocation(tz)
				if err != nil {
					parts = append(parts, fmt.Sprintf("issue with zone: %v", err))
					continue
				}
				parts = append(parts, t.In(loc).Format(*flagFormatString))
			}
			obj.Button().SetTitle(strings.Join(parts, ", "))
		}
	}()
}
