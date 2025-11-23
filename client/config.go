package client

import "time"

type Config struct {
	URL          string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
