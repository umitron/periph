// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package tm1637_test

import (
	"log"

	"github.com/umitron/periph/conn/gpio/gpioreg"
	"github.com/umitron/periph/devices/tm1637"
	"github.com/umitron/periph/host"
)

func Example() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	clk := gpioreg.ByName("GPIO6")
	data := gpioreg.ByName("GPIO12")
	if clk == nil || data == nil {
		log.Fatal("Failed to find pins")
	}
	dev, err := tm1637.New(clk, data)
	if err != nil {
		log.Fatalf("failed to initialize tm1637: %v", err)
	}
	if err := dev.SetBrightness(tm1637.Brightness10); err != nil {
		log.Fatalf("failed to set brightness on tm1637: %v", err)
	}
	if _, err := dev.Write(tm1637.Clock(12, 00, true)); err != nil {
		log.Fatalf("failed to write to tm1637: %v", err)
	}
}
