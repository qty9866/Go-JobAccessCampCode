package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	r := TestBox{}
	var c Close = &r

	switch cDetail := c.(type) {
	case Refrigerator:
		fmt.Println("是预期的冰箱")
		fmt.Println(cDetail.Size)
	case TestBox:
		fmt.Println("这是个box，不能当做冰箱使用")
	}

	refrigerator, ok := checkIsRefrigerator(c)
	if ok {
		fmt.Println("是个冰箱，开门装大象")
	} else {
		fmt.Println("不是个冰箱")
	}

	/*	r2 := c.(Refrigerator)
		fmt.Println(r2.Size)*/
}

func checkIsRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator)
	return r, ok
}

type TestBox struct {
}

func (tb TestBox) Close() error {
	return nil
}
