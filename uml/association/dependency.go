package dependency

type CouponTpl struct {
	Id     int64
	amount int64
	name   string
}

type Act struct {
}

func (a Act) GetDisAmount(tpl CouponTpl) int64 {
	return tpl.amount
}
