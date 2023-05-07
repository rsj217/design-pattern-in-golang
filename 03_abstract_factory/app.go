package abstract_factory

import (
	"fmt"
)

// 权益接口
type Benefiter interface {
	GetDisAmount() int64
}

type HotelVoucher struct {
	disAmount int64
}

func (hv *HotelVoucher) GetDisAmount() int64 {
	return hv.disAmount
}

type ScenicVoucher struct {
	disAmount int64
}

func (sv *ScenicVoucher) GetDisAmount() int64 {
	return sv.disAmount
}

// 会员卡接口
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

// 抽象工厂接口
type BizFactory interface {
	CreateBenefiter() Benefiter
	CreateCarder() Carder
}

// 酒店工厂
type HotelFactory struct {
	config Config
}

func (hf *HotelFactory) CreateBenefiter() Benefiter {
	return &HotelVoucher{hf.config.disAmount}
}
func (hf *HotelFactory) CreateCarder() Carder {
	return &HotelCard{hf.config.saleAmount}
}

// 景区工厂
type ScenicFactory struct {
	config Config
}

func (sf *ScenicFactory) CreateBenefiter() Benefiter {
	return &ScenicVoucher{sf.config.disAmount}

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
	disAmount  int64
	saleAmount int64
}

func client() {
	conf := Config{int64(10), int64(20)}

	client := NewApplication(&HotelFactory{conf})
	disAmount := client.factory.CreateBenefiter().GetDisAmount()
	saleAmount := client.factory.CreateCarder().GetSaleAmount()
	fmt.Printf("hotel voucher disAmount: %d, card: %d \n", disAmount, saleAmount)

	client = NewApplication(&ScenicFactory{conf})
	disAmount = client.factory.CreateBenefiter().GetDisAmount()
	saleAmount = client.factory.CreateCarder().GetSaleAmount()
	fmt.Printf("scenic voucher disAmount: %d, card: %d \n", disAmount, saleAmount)
}
