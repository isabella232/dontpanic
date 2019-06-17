package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"code.cloudfoundry.org/dontpanic/osreporter"
	"code.cloudfoundry.org/dontpanic/plugins/command"
)

const extractDir = "/var/vcap/data/tmp"

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("could not determine hostname")
		hostname = "UNKNOWN-HOSTNAME"
	}

	osReporter := osreporter.New(extractDir, hostname, time.Now(), os.Stdout)

	osReporter.RegisterEchoStream("Date", "date.log", command.New("date"))
	osReporter.RegisterEchoStream("Uptime", "uptime.log", command.New("uptime"))

	if err := osReporter.Run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
