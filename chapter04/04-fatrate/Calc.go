package main

import (
	"log"

	gobmi "github.com/armstrongli/go-bmi"
)

type Calc struct{}

func (Calc) BMI(p *Person) error {
	bmi, err := gobmi.BMI(p.weight, p.weight)
	if err != nil {
		log.Println("error occurred when calculating bmi")
		return err
	}
	p.bmi = bmi
	return nil
}

func (Calc) Fatrate(p *Person) error {
	p.fat = gobmi.CalcFatRate(p.bmi, p.age, p.sex)
	return nil
}
