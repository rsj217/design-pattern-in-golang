package visitor

import "fmt"

type Visitor interface {
	VisitorForHotelWorker(*HotelWorker)
	VisitorForScenicWorker(*ScenicWorker)
}

type Worker interface {
	GetName()
	Accept(Visitor)
}

type HotelWorker struct {
	name string
}

func (hw *HotelWorker) GetName() string {
	return hw.name
}

func (hw *HotelWorker) Accept(visitor Visitor) {
	visitor.VisitorForHotelWorker(hw)
}

type ScenicWorker struct {
	name string
}

func (sw *ScenicWorker) GetName() string {
	return sw.name
}
func (sw *ScenicWorker) Accept(visitor Visitor) {
	visitor.VisitorForScenicWorker(sw)
}

type LogVisitor struct {
}

func (lv LogVisitor) VisitorForHotelWorker(hw *HotelWorker) {
	fmt.Printf("logVisitor record log for %v\n", hw.GetName())
}

func (lv LogVisitor) VisitorForScenicWorker(lw *ScenicWorker) {
	fmt.Printf("logVisitor record log for %v\n", lw.GetName())
}

type ExportVisitor struct{}

func (ev ExportVisitor) VisitorForHotelWorker(hw *HotelWorker) {
	fmt.Printf("ExportVisitor export data from %v\n", hw.GetName())
}

func (ev ExportVisitor) VisitorForScenicWorker(lw *ScenicWorker) {
	fmt.Printf("ExportVisitor export data from %v\n", lw.GetName())
}

func Client() {
	hw := &HotelWorker{"hotelWorker"}
	lw := &ScenicWorker{"ScenicWorker"}

	lv := LogVisitor{}
	hw.Accept(lv)
	lw.Accept(lv)

	ev := ExportVisitor{}
	hw.Accept(ev)
	lw.Accept(ev)

}
