package main

import (
	"log"

	"Learning-JobAccess-Camp/pkg/apis"
	goBMI "github.com/armstrongli/go-bmi"
)

type Calc struct {
	continental string
}

func (c *Calc) BMI(person *apis.PersonalInformation) (float64, error) {
	bmi, err := goBMI.BMI(float64(person.Weight), float64(person.Tall))
	if err != nil {
		log.Println("error when calculating bmi:", err)
		return -1, err
	}
	return bmi, nil
}

func (c *Calc) FatRate(person *apis.PersonalInformation) (float64, error) {
	bmi, err := c.BMI(person)
	if err != nil {
		return -1, err
	}
	return goBMI.CalcFatRate(bmi, int(person.Age), person.Sex), nil
}
