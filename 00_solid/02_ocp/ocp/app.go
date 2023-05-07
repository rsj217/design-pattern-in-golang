package ocp

import "fmt"

type Specification interface {
	IsSatisfied(c Coupon) bool
}

type Cp int

const (
	Fliggy Cp = iota
	CTrip
)

type CpSpecification struct {
	cp Cp
}

func (c CpSpecification) IsSatisfied(coupon Coupon) bool {
	return c.cp == coupon.cp
}

type Scene int

const (
	SearchScene Scene = iota
	ActScene
)

type SceneSpecification struct {
	scene Scene
}

func (s SceneSpecification) IsSatisfied(coupon Coupon) bool {
	return s.scene == coupon.scene
}

type Coupon struct {
	name  string
	cp    Cp
	scene Scene
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(coupon Coupon) bool {
	return a.first.IsSatisfied(coupon) && a.second.IsSatisfied(coupon)
}

type BetterFilter struct {
}

func (b BetterFilter) Filter(coupons []Coupon, spec Specification) []*Coupon {
	ret := make([]*Coupon, 0)
	for i, v := range coupons {
		if spec.IsSatisfied(v) {
			ret = append(ret, &coupons[i])
		}
	}
	return ret
}

func client() {
	f1 := Coupon{"满减券", Fliggy, SearchScene}
	f2 := Coupon{"折扣券", Fliggy, ActScene}
	c := Coupon{"满减券", CTrip, SearchScene}

	coupons := []Coupon{f1, f2, c}
	cpSpecification := CpSpecification{Fliggy}

	f := BetterFilter{}

	for _, v := range f.Filter(coupons, cpSpecification) {
		fmt.Printf("%s is Fliggy\n", v.name)
	}

	andSpecification := AndSpecification{cpSpecification, SceneSpecification{SearchScene}}
	for _, v := range f.Filter(coupons, andSpecification) {
		fmt.Printf("%s is Fliggy and Scene\n", v.name)
	}
}
