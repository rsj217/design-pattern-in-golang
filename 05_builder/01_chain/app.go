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

type Builder struct {
	userManger UserManger
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build() UserManger {
	return b.userManger
}

func (b *Builder) UserId(userId int) *Builder {
	b.userManger.userId = userId
	return b
}

func (b *Builder) Adiu(adiu string) *Builder {
	b.userManger.adiu = adiu
	return b
}

func (b *Builder) bizType(bizType string) *Builder {
	b.userManger.bizType = bizType
	return b
}

func (b *Builder) cpcode(cpcode string) *Builder {
	b.userManger.cpcode = cpcode
	return b
}

func (b *Builder) orderId(orderId string) *Builder {
	b.userManger.orderId = orderId
	return b
}

func client() {
	// 售卖
	builder := NewBuilder()
	userManger := builder.
		UserId(622217).
		bizType("hotel").
		Adiu("jdser80llmc").
		cpcode("10006").
		Build()
	fmt.Println(userManger.userId)

	// 开卡
	userManger = builder.UserId(622217).bizType("hotel").orderId("0100234578987").Build()
	fmt.Println(userManger.orderId)

}
