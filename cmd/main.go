package main

import (
	"hot/internal/pkg/flags"
	"hot/internal/pkg/server"
)

func main() {
	port, dir := flags.Flags()
	server.Start(port, dir)
}
