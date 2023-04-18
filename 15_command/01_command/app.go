package command

import "fmt"

type Act struct {
	stock int64
}

func (a *Act) Produce(stock int64) {
	a.stock += stock
	fmt.Printf("Produce %d, stock is now %d\n", stock, a.stock)
}

func (a *Act) Consume(stock int64) bool {
	if a.stock-stock >= 0 {
		a.stock -= stock
		fmt.Printf("Consume %d, stock is now %d\n", stock, a.stock)
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
}

type Action int64

const (
	Produce Action = iota
	Consume
)

var _ Command = (*ActCommand)(nil)

type ActCommand struct {
	succeeded bool
	stock     int64
	act       *Act
	action    Action
}

func NewActCommand(act *Act, action Action, stock int64) *ActCommand {
	return &ActCommand{stock: stock, act: act, action: action}
}

func (a *ActCommand) Call() {
	switch a.action {
	case Produce:
		a.act.Produce(a.stock)
		a.succeeded = true
	case Consume:
		a.succeeded = a.act.Consume(a.stock)
	}
}

func (a *ActCommand) Undo() {
	if !a.succeeded {
		return
	}
	switch a.action {
	case Produce:
		a.act.Consume(a.stock)
	case Consume:
		a.act.Produce(a.stock)
	}
}

func Client() {
	act := &Act{}
	cmd1 := NewActCommand(act, Produce, 100)
	cmd1.Call()

	cmd2 := NewActCommand(act, Consume, 30)
	cmd2.Call()

	fmt.Println(act)

	cmd2.Undo()

	fmt.Println(act)

}
