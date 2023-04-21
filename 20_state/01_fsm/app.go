package state

import "fmt"

type Stater interface {
	transfer(string) Stater
}

var _ Stater = (*OneState)(nil)

type OneState struct{}

func (o *OneState) String() string {
	return "1"
}

func (o *OneState) transfer(s string) Stater {
	if s == "0" {
		return &ZeroState{}
	}
	return o
}

type ZeroState struct{}

func (z *ZeroState) String() string {
	return "0"
}

func (z *ZeroState) transfer(s string) Stater {
	if s == "1" {
		return &OneState{}
	}
	return z
}

type Machine struct {
	state Stater
}

func (m *Machine) Run(input string) {
	m.state = m.state.transfer(input)
}

func client() {
	s := "1001010001001010"

	machine := Machine{
		state: &ZeroState{},
	}

	for _, v := range s {
		machine.Run(string(v))
	}
	fmt.Println(machine.state)
}
