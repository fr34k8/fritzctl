package cliapp

import (
	"strconv"

	"github.com/bpicode/fritzctl/assert"
	"github.com/bpicode/fritzctl/fritz"
	"github.com/mitchellh/cli"
)

type temperatureCommand struct {
}

func (cmd *temperatureCommand) Help() string {
	return "Set the temperature of a HKR device. Example usage: fritzctl temperature 21.0 MY_HKR_DEVICE"
}

func (cmd *temperatureCommand) Synopsis() string {
	return "set the temperature of a HKR device"
}

func (cmd *temperatureCommand) Run(args []string) int {
	assert.StringSliceHasAtLeast(args, 2, "insufficient input: two parameters expected.")
	temp, errorParse := strconv.ParseFloat(args[0], 64)
	assert.NoError(errorParse, "cannot parse temperature value:", errorParse)
	f := fritz.New(clientLogin())
	err := f.Temperature(temp, args[1:]...)
	assert.NoError(err, "error setting temperature:", err)
	return 0
}

func temperature() (cli.Command, error) {
	p := temperatureCommand{}
	return &p, nil
}
