package flags

import (
	"flag"
	"fmt"
	"os"
)

func Flags() (int, string) {
	help := flag.Bool("help", false, "help panel")
	port := flag.Int("port", 8080, "Port number")
	dir := flag.String("dir", "data", "Path to the directory")
	flag.Parse()

	if *help {
		showHelp()
		os.Exit(0)
	}
	return *port, *dir
}

func showHelp() {
	helpText := `
Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>] 
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.
`
	fmt.Println(helpText)
}
