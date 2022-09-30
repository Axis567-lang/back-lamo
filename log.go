package main

import "log"

func logOutput(msg string) {
	_, _ = log.Writer().Write([]byte(msg + "\n"))
}
