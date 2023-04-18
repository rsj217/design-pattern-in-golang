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

type Clientlication struct {
	factory BizFactory
	benefit Benefiter
}

func NewClientlication(factory BizFactory) *Clientlication {
	return &Clientlication{factory: factory}
}

type Config struct {
	discount   int64
	saleAmount int64
}

func Client() {
	conf := Config{int64(10), int64(20)}

	Client := NewClientlication(&HotelFactory{conf})
	discount := Client.factory.CreateBenefiter().GetDiscount()
	saleAmount := Client.factory.CreateCarder().GetSaleAmount()
	fmt.Printf("hotel voucher discount: %d, card: %d \n", discount, saleAmount)

	Client = NewClientlication(&ScenicFactory{conf})
	discount = Client.factory.CreateBenefiter().GetDiscount()
	saleAmount = Client.factory.CreateCarder().GetSaleAmount()
	fmt.Printf("scenic voucher discount: %d, card: %d \n", discount, saleAmount)
}
