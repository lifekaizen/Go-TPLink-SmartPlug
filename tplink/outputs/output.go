package tpoutput

import (
	"errors"

	tpdevices "github.com/mikemrm/Go-TPLink-SmartPlug/tplink/devices"
)

type Output interface {
	Write(tpdevices.TPDevices) error
}

type bad struct{}

func (b *bad) Write(devices tpdevices.TPDevices) error {
	return nil
}

type Outputer func() (error, Output)

var Outputs = map[string]Outputer{}

func AddOutput(name string, output Outputer) {
	Outputs[name] = output
}

func GetOutput(name string) (error, Output) {
	output, ok := Outputs[name]
	if !ok {
		return errors.New("Output not found."), &bad{}
	}
	return output()
}
