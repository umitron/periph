// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package sn3218_test

import (
	"log"
	"time"

	"github.com/umitron/periph/conn/i2c/i2creg"
	"github.com/umitron/periph/experimental/devices/sn3218"
	"github.com/umitron/periph/host"
)

func Example() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	d, err := sn3218.New(b)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Halt()

	// By default, the device is disabled and brightness is 0 for all LEDs
	// So let's set the brightness to a low value and enable the device to
	// get started
	d.BrightnessAll(1)
	d.Sleep()

	// Switch LED 7 on
	if err := d.Switch(7, true); err != nil {
		log.Fatal("Error while switching LED", err)
	}

	time.Sleep(1000 * time.Millisecond)
}
