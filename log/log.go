// Package log wraps the standard log package in order to ensure that fatal
// errors are never discarded.
package log

import (
	l "log"
	"os"
)

// FatalWriter is the io.Writer that the logger falls back to for Fatal[f|ln]
// function calls.
var FatalWriter = os.Stderr

// Fatal is equivalent to Print() followed by a call to os.Exit(1). FatalWriter
// is used.
func Fatal(v ...interface{}) {
	SetOutput(FatalWriter)
	l.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
// FatalWriter is used.
func Fatalf(format string, v ...interface{}) {
	SetOutput(FatalWriter)
	l.Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
// FatalWriter is used.
func Fatalln(v ...interface{}) {
	SetOutput(FatalWriter)
	l.Fatalln(v...)
}

var (
	// Flags returns the output flags for the standard logger.
	Flags = l.Flags

	// Output writes the output for a logging event. The string s contains the
	// text to print after the prefix specified by the flags of the Logger. A
	// newline is appended if the last character of s is not already a newline.
	// Calldepth is the count of the number of frames to skip when computing
	// the file name and line number if Llongfile or Lshortfile is set; a value
	// of 1 will print the details for the caller of Output.
	Output = l.Output

	// Panic is equivalent to Print() followed by a call to panic().
	Panic = l.Panic

	// Panicf is equivalent to Printf() followed by a call to panic().
	Panicf = l.Panicf

	// Panicln is equivalent to Println() followed by a call to panic().
	Panicln = l.Panicln

	// Prefix returns the output prefix for the standard logger.
	Prefix = l.Prefix

	// Print calls Output to print to the standard logger. Arguments are
	// handled in the manner of fmt.Print.
	Print = l.Print

	// Printf calls Output to print to the standard logger. Arguments are
	// handled in the manner of fmt.Printf.
	Printf = l.Printf

	// Println calls l.Output to print to the logger. Arguments are handled in
	// the manner of fmt.Println.
	Println = l.Println

	// SetFlags sets the output flags for the logger.
	SetFlags = l.SetFlags

	// SetOutput sets the output destination for the logger.
	SetOutput = l.SetOutput

	// SetPrefix sets the output prefix for the logger.
	SetPrefix = l.SetPrefix
)
