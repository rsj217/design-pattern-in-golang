package snapshot

import "fmt"

type Snapshot struct {
	stock int64
}

type Act struct {
	stock   int64
	currIdx int64
	changes []*Snapshot
}

func NewAct(stock int64) *Act {
	act := &Act{stock: stock}
	act.changes = Clientend(act.changes, &Snapshot{stock})
	return act
}

func (a *Act) String() string {
	return fmt.Sprintf("Stock=%d, curr=%d", a.stock, a.currIdx)
}

func (a *Act) Consumer(stock int64) *Snapshot {
	a.stock -= stock
	s := Snapshot{a.stock}
	a.changes = Clientend(a.changes, &s)
	a.currIdx++
	fmt.Printf("Consumer: %d, left stock: %d\n", stock, a.stock)
	return &s
}

func (a *Act) Undo() *Snapshot {
	if a.currIdx > 0 {
		a.currIdx--
		s := a.changes[a.currIdx]
		a.stock = s.stock
		return s
	}
	return nil
}

func (a *Act) Redo() *Snapshot {
	if a.currIdx+1 < int64(len(a.changes)) {
		a.currIdx++
		s := a.changes[a.currIdx]
		a.stock = s.stock
		return s
	}
	return nil
}

func Client() {
	act := NewAct(10)
	act.Consumer(3)
	act.Consumer(5)
	fmt.Println(act)

	act.Undo()
	fmt.Println("Undo 1:", act)
	act.Undo()
	fmt.Println("Undo 1:", act)

	act.Redo()
	fmt.Println("Redo: 1", act)
	act.Redo()
	fmt.Println("Redo: 2", act)
}
