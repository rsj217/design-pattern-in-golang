package command

import "fmt"

type Transfer interface {
	Produce(int64)
	Consume(int64) bool
}

type Act struct {
	stock int64
}

func NewAct() *Act {
	return &Act{}
}

func (a *Act) Produce(stock int64) {
	a.stock += stock
	fmt.Printf("[Act] Produce %d, stock is now %d\n", stock, a.stock)
}

func (a *Act) Consume(stock int64) bool {
	if a.stock-stock >= 0 {
		a.stock -= stock
		fmt.Printf("[Act] Consume %d, stock is now %d\n", stock, a.stock)
		return true
	}
	return false
}

type Score struct {
	token int64
}

func NewScore() *Score {
	return &Score{}
}

func (s *Score) String() string {
	return fmt.Sprintf("&{%d %d}", s.token, s.getScore())
}

func (s *Score) getScore() int64 {
	return s.token * 5
}

func (s *Score) Produce(token int64) {
	s.token += token
	fmt.Printf("[Score] Produce %d, token=%d score=%d\n", token, s.token, s.getScore())
}

func (s *Score) Consume(token int64) bool {
	if s.token-token >= 0 {
		s.token -= token
		fmt.Printf("[Score] Consume %d, token=%d score=%d\n", token, s.token, s.getScore())
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(value bool)
}

type Action int64

const (
	Produce Action = iota
	Consume
)

var _ Command = (*TransactCommand)(nil)

type TransactCommand struct {
	succeeded  bool
	token      int64
	transacter Transfer
	action     Action
}

func NewTransactCommand(act Transfer, action Action, token int64) *TransactCommand {
	return &TransactCommand{token: token, transacter: act, action: action}
}

func (a *TransactCommand) Succeeded() bool {
	return a.succeeded
}

func (a *TransactCommand) SetSucceeded(value bool) {
	a.succeeded = value
}

func (a *TransactCommand) Call() {
	switch a.action {
	case Produce:
		a.transacter.Produce(a.token)
		a.succeeded = true
	case Consume:
		a.succeeded = a.transacter.Consume(a.token)
	}
}

func (a *TransactCommand) Undo() {
	if !a.succeeded {
		return
	}
	switch a.action {
	case Produce:
		a.transacter.Consume(a.token)
	case Consume:
		a.transacter.Produce(a.token)
	}
}

var _ Command = (*CompositeActCommand)(nil)

type CompositeActCommand struct {
	commands []Command
}

func (cac *CompositeActCommand) Call() {
	for _, v := range cac.commands {
		v.Call()
	}
}

func (cac *CompositeActCommand) Undo() {
	for idx := range cac.commands {
		cac.commands[len(cac.commands)-1-idx].Undo()
	}
}

func (cac *CompositeActCommand) Succeeded() bool {
	//TODO implement me
	panic("implement me")
}

func (cac *CompositeActCommand) SetSucceeded(value bool) {
	//TODO implement me
	panic("implement me")
}

var _ Command = (*TokenTransferCommand)(nil)

type TokenTransferCommand struct {
	CompositeActCommand
	from, to Transfer
	token    int64
}

func NewTokenTransferCommand(from, to Transfer, token int64) *TokenTransferCommand {
	t := &TokenTransferCommand{from: from, to: to, token: token}
	t.commands = append(t.commands, NewTransactCommand(from, Consume, token))
	t.commands = append(t.commands, NewTransactCommand(to, Produce, token))
	return t
}

func (tc TokenTransferCommand) Call() {
	ok := true
	for _, v := range tc.commands {
		if ok {
			v.Call()
			ok = v.Succeeded()
		} else {
			v.SetSucceeded(false)
		}
	}
}

func Client() {
	act := NewAct()
	cmd1 := NewTransactCommand(act, Produce, 100)
	cmd2 := NewTransactCommand(act, Consume, 30)

	cmd1.Call()
	cmd2.Call()
	fmt.Println(act)

	cmd2.Undo()
	cmd1.Undo()
	fmt.Println(act)

	fmt.Println("Transfer ....")

	from := NewAct()
	from.Produce(100)
	to := NewScore()
	to.Produce(20)

	fmt.Printf("from=%v to=%v \n", from, to)

	tc := NewTokenTransferCommand(from, to, 30)
	tc.Call()

	fmt.Printf("from=%v to=%v \n", from, to)

	tc.Undo()

	fmt.Printf("from=%v to=%v \n", from, to)

}
