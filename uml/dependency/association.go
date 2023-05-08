package association

type CouponTpl struct {
	Id     int64
	amount int64
	name   string
}

type Act struct {
	couponTplList []*CouponTpl
}
