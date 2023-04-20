package observer

import "fmt"

type Observer interface {
	Update(string)
	GetID() int64
}

type Subjecter interface {
	Register(Observer)
	UnRegister(Observer)
	NotifyAll()
}

type CouponManger struct {
	id int64
}

func (cm *CouponManger) Update(orderId string) {
	fmt.Printf("[CouponManger] handle event: update coupon status by orderId=%s\n", orderId)
}

func (cm *CouponManger) GetID() int64 {
	return cm.id
}

type LogManger struct {
	id int64
}

func (lm *LogManger) Update(orderId string) {
	fmt.Printf("[LogManger] handle event: write log for monitor via orderId=%s\n", orderId)
}

func (lm *LogManger) GetID() int64 {
	return lm.id
}

type Consumer struct {
	msg          string
	observerList []Observer
}

func (c *Consumer) Register(observer Observer) {
	c.observerList = append(c.observerList, observer)
}

func (c *Consumer) UnRegister(observer Observer) {
	size := len(c.observerList)
	for i, v := range c.observerList {
		if observer.GetID() == v.GetID() {
			c.observerList[size-1], c.observerList[i] = c.observerList[i], c.observerList[size-1]
			break
		}
	}
	c.observerList = c.observerList[0 : size-1]
}

func (c *Consumer) NotifyAll() {
	for _, v := range c.observerList {
		v.Update(c.msg)
	}
}
func (c *Consumer) ReceiveMsg(msg string) {
	fmt.Printf("[consumer]receive msg=%s\n", msg)
	c.msg = msg
	c.NotifyAll()
}

func Client() {
	cm := &CouponManger{
		1,
	}
	lm := &LogManger{
		2,
	}

	consumer := &Consumer{}

	consumer.Register(cm)
	consumer.Register(lm)
	msg := "202304050001"
	consumer.ReceiveMsg(msg)

}
