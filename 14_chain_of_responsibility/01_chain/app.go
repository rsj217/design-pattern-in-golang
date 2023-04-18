package chainofreponsibility

import (
	"errors"
	"fmt"
	"time"
)

type Filter interface {
	Check(*Request) error
}

type Request struct {
	adiu        string
	poiid       string
	benefitType int64
	actTime     time.Time
}

type ADiuFilter struct {
	value string
}

func (a ADiuFilter) Check(request *Request) error {
	fmt.Printf("check adiu=%s\n", request.adiu)
	if a.value != request.adiu {
		return errors.New("adiu err")
	}
	return nil
}

type PoiidFilter struct {
	value string
}

func (p PoiidFilter) Check(request *Request) error {
	fmt.Printf("check poiid=%s\n", request.poiid)
	if p.value != request.poiid {
		return errors.New("poiid err")
	}
	return nil
}

type BenefitTypeFilter struct {
	value int64
}

func (b BenefitTypeFilter) Check(request *Request) error {
	fmt.Printf("check benefitType=%d\n", request.benefitType)
	if b.value != request.benefitType {
		return errors.New("benefitType err")
	}
	return nil
}

type ActTimeFilter struct {
	value time.Time
}

func (a ActTimeFilter) Check(request *Request) error {
	fmt.Printf("check ActTime=%s\n", request.actTime)
	if a.value != request.actTime {
		return errors.New("actTime err")
	}
	return nil
}

func Chain(request *Request, filters []Filter) error {
	for _, v := range filters {
		if err := v.Check(request); err != nil {
			return err
		}
	}
	return nil
}

func Client() {
	request := &Request{
		"adiu_1234",
		"BX10010",
		1,
		time.Now(),
	}
	filters := []Filter{
		ADiuFilter{"adiu_1234"},
		PoiidFilter{"BX10010"},
		ActTimeFilter{time.Now()},
		BenefitTypeFilter{1},
	}

	err := Chain(request, filters)
	fmt.Println("result err=", err)

	filters = []Filter{
		ADiuFilter{"adiu_1234"},
		PoiidFilter{"BX10010"},
		BenefitTypeFilter{1},
	}

	err = Chain(request, filters)
	fmt.Println("result err=", err)
}
