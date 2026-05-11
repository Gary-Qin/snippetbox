package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// custom flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// initialize app
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, app.routes())

	// exit
	logger.Error(err.Error())
	os.Exit(1)
}
