package ghcp

import (
	"log"
	"sync"
	"strings"
)

var once sync.Once

type ghcpLogger struct {
	debug bool
}

var ghcpLog ghcpLogger

func debugLog() ghcpLogger {

	once.Do(func() {
		ghcpLog = ghcpLogger{}

		if strings.ToLower(getConfig().Debug) == "true" {
			ghcpLog.debug = true
		}
	})

	return ghcpLog
}

func (rcv ghcpLogger) printf(format string, v ...interface{}) {
	if rcv.debug {
		log.Printf("GHCP Debug: " + format, v...)
	}
}

func (rcv ghcpLogger) println(text string) {
	if rcv.debug {
		log.Println("GHCP Debug:", text)
	}
}