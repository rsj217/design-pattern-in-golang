package builder

import (
	"fmt"
)

type UserManger struct {
	userId  int
	adiu    string
	bizType string
	cpcode  string
	orderId string
}

type Option func(u *UserManger) error

func WithUserId(userId int) Option {
	return func(u *UserManger) error {
		u.userId = userId
		return nil
	}
}

func WithADiu(adiu string) Option {
	return func(u *UserManger) error {
		u.adiu = adiu
		return nil
	}
}

func WithBizType(bizType string) Option {
	return func(u *UserManger) error {
		u.bizType = bizType
		return nil
	}
}
func WithCpcode(cpcode string) Option {
	return func(u *UserManger) error {
		u.cpcode = cpcode
		return nil
	}
}
func WithOrderId(orderId string) Option {
	return func(u *UserManger) error {
		u.orderId = orderId
		return nil
	}
}

func NewUserManger(opts ...Option) *UserManger {
	u := &UserManger{}
	for _, opt := range opts {
		if err := opt(u); err != nil {
			panic("err")
		}
	}
	return u
}

func client() {

	// 售卖
	userManger := NewUserManger(
		WithUserId(622217),
		WithBizType("hotel"),
		WithADiu("jdser80llmc"),
		WithCpcode("10006"),
	)
	fmt.Println(userManger.userId)

	// 开卡
	userManger = NewUserManger(WithUserId(622217), WithBizType("hotel"), WithOrderId("0100234578987"))
	fmt.Println(userManger.orderId)

}
