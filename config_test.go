package main

import "testing"

func TestBadConfig(t *testing.T) {
	testCases := []struct {
		desc string
		cfg  Config
		want error
	}{
		{"Small width", Config{Width: 9, Height: 10, Speed: 1}, ErrTooSmall},
		{"Small height", Config{Width: 10, Height: 9, Speed: 1}, ErrTooSmall},
		{"Small speed", Config{Width: 10, Height: 10, Speed: 0}, ErrBadSpeed},
		{"Big speed", Config{Width: 10, Height: 10, Speed: 61}, ErrBadSpeed},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.cfg.Validate()
			if err == nil {
				t.Fatal("Expected error got nil")
			}
			if err != tC.want {
				t.Fatalf("Want(%s), got(%s)", tC.want, err)
			}
		})
	}
}
