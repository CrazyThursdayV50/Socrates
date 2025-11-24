package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/CrazyThursdayV50/Socrates/internal/repository/chatter/impl/gemini"
	"github.com/CrazyThursdayV50/Socrates/server"
	"github.com/CrazyThursdayV50/pkgo/config"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
	"github.com/CrazyThursdayV50/pkgo/trace/jaeger"
	"github.com/spf13/pflag"
)

var dir string
var name string
var ext string

func init() {
	pflag.StringVar(&dir, "d", ".", "config dir")
	pflag.StringVar(&name, "n", "config", "config name")
	pflag.StringVar(&ext, "ext", "yml", "config extension")
}

func main() {
	pflag.Parse()

	// load config
	cfg, err := config.GetConfig[Config](dir, name, ext)
	if err != nil {
		panic(err)
	}

	logger := sugar.New(&cfg.Log)

	var wg sync.WaitGroup
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	goo.Go(func() {
		defer cancel()
		var signalChan = make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		<-signalChan
	})

	tracer, err := jaeger.New(ctx, &cfg.Jaeger, logger)
	if err != nil {
		panic(err)
	}

	gemini := gemini.New(&cfg.Chatter, logger)
	err = gemini.LoadSystem()
	if err != nil {
		panic(err)
	}

	server := server.New(
		&cfg.Server,
		logger,
		tracer,
		gemini,
	)
	server.Init()
	server.Run(ctx, &wg)
}
