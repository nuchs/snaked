package main

import (
	"errors"
	"flag"
	"fmt"
)

var (
	ErrTooSmall = errors.New("Width & height must  be >= 10")
	ErrBadSpeed = errors.New("Speed must be between 1 & 60")
)

type Config struct {
	Width  int
	Height int
	Speed  int
}

func (c Config) String() string {
	return fmt.Sprintf("width=%d, height=%d, speed=%d", c.Width, c.Height, c.Speed)
}

func (c Config) Validate() error {
	if c.Width < 10 || c.Height < 10 {
		return ErrTooSmall
	}
	if c.Speed < 1 || c.Speed > 60 {
		return ErrBadSpeed
	}
	return nil
}

func FromFlags() (Config, error) {
	var cfg Config

	flag.IntVar(&cfg.Width, "width", 40, "Board width (>= 10)")
	flag.IntVar(&cfg.Height, "height", 40, "Board height (>= 10)")
	flag.IntVar(&cfg.Speed, "speed", 8, "Ticks per second (1..60)")
	flag.Parse()

	return cfg, cfg.Validate()
}
