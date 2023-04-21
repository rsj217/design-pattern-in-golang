package state

import "fmt"

type Event int64

const (
	LockEvent Event = iota + 1
	PaidEvent
	RefundEvent
)

type Stater interface {
	Transfer(Event) Stater
}

type NormalState struct {
}

func (ns *NormalState) String() string {
	return "NormalState"
}

func (ns *NormalState) Transfer(event Event) Stater {
	if event == LockEvent {
		return &LockState{}
	}
	return ns
}

type LockState struct {
}

func (ls *LockState) String() string {
	return "LockState"
}

func (ls *LockState) Transfer(event Event) Stater {
	switch event {
	case PaidEvent:
		return &UsedState{}
	case RefundEvent:
		return &NormalState{}
	}
	return ls
}

type UsedState struct {
}

func (us *UsedState) String() string {
	return "UsedState"
}
func (us *UsedState) Transfer(event Event) Stater {
	if event == RefundEvent {
		return &NormalState{}
	}
	return us
}

type ExpiredState struct {
}

func (es *ExpiredState) String() string {
	return "ExpiredState"
}
func (es *ExpiredState) Transfer(event Event) Stater {
	return es
}

type StateFSM struct {
	state Stater
}

func (sf *StateFSM) Run(event Event) {
	sf.state = sf.state.Transfer(event)
}

func client() {

	fsm := StateFSM{&NormalState{}}

	fsm.Run(LockEvent)
	fmt.Println(fsm.state)

	fsm.Run(LockEvent)
	fmt.Println(fsm.state)

	fsm.Run(PaidEvent)
	fmt.Println(fsm.state)

	fsm.Run(RefundEvent)
	fmt.Println(fsm.state)

}
