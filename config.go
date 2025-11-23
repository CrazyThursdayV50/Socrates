package main

import (
	"github.com/CrazyThursdayV50/Socrates/server"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
	"github.com/CrazyThursdayV50/pkgo/trace/jaeger"
)

type Config struct {
	Log    sugar.Config
	Jaeger jaeger.Config
	Server server.Config
}
