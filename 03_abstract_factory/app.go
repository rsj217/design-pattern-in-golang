package abstract_factory

import (
	"fmt"
)

type Benefiter interface {
	GetDiscount() int64
}

type HotelVoucher struct {
	discount int64
}

func (hv *HotelVoucher) GetDiscount() int64 {
	return hv.discount
}

type ScenicVoucher struct {
	discount int64
}

func (sv *ScenicVoucher) GetDiscount() int64 {
	return sv.discount
}

type Carder interface {
	GetSaleAmount() int64
}

type HotelCard struct {
	saleAmount int64
}

func (hc *HotelCard) GetSaleAmount() int64 {
	return hc.saleAmount
}

type ScenicCard struct {
	saleAmount int64
}

func (sc *ScenicCard) GetSaleAmount() int64 {
	return sc.saleAmount
}

type BizFactory interface {
	CreateBenefiter() Benefiter
	CreateCarder() Carder
}

type HotelFactory struct {
	config Config
}

func (hf *HotelFactory) CreateBenefiter() Benefiter {
	return &HotelVoucher{hf.config.discount}
}
func (hf *HotelFactory) CreateCarder() Carder {
	return &HotelCard{hf.config.saleAmount}
}

type ScenicFactory struct {
	config Config
}

func (sf *ScenicFactory) CreateBenefiter() Benefiter {
	return &ScenicVoucher{sf.config.discount}

}
func (sf *ScenicFactory) CreateCarder() Carder {
	return &ScenicCard{sf.config.saleAmount}
}

type Application struct {
	factory BizFactory
	benefit Benefiter
}

func NewApplication(factory BizFactory) *Application {
	return &Application{factory: factory}
}

type Config struct {
	discount   int64
	saleAmount int64
}

func App() {
	conf := Config{int64(10), int64(20)}

	app := NewApplication(&HotelFactory{conf})
	discount := app.factory.CreateBenefiter().GetDiscount()
	saleAmount := app.factory.CreateCarder().GetSaleAmount()
	fmt.Printf("hotel voucher discount: %d, card: %d \n", discount, saleAmount)

	app = NewApplication(&ScenicFactory{conf})
	discount = app.factory.CreateBenefiter().GetDiscount()
	saleAmount = app.factory.CreateCarder().GetSaleAmount()
	fmt.Printf("scenic voucher discount: %d, card: %d \n", discount, saleAmount)
}
