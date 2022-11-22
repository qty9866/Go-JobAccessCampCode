package main

import "fmt"

type manLegend struct {
}

func (man *manLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manLegend 做OpenTheDoorOfRefrigerator")
	return nil
}
func (man *manLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 manLegend 做PutElephantIntoRefrigerator")
	return nil
}
func (man *manLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manLegend 做CloseTheDoorOfRefrigerator")
	return nil
}
