package main

import "fmt"

type logLevel int

const (
	levelTrace logLevel = iota
	levelDebug
	levelInfo
	levelWarning
	levelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (level logLevel) name() string {
	if level < levelTrace || level > levelError {
		return "Unknown"
	}

	return levelNames[level]
}

func printLogLevel(level logLevel) {
	fmt.Printf("Log level: %d - %s\n", level, level.name())
}

func main() {
	printLogLevel(levelTrace)
	printLogLevel(levelDebug)
	printLogLevel(levelInfo)
	printLogLevel(levelWarning)
	printLogLevel(levelError)
	printLogLevel(8)
}
