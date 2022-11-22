package main

import "fmt"

type shipLegend struct {
}

func (ship *shipLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用ship做OpenTheDoorOfRefrigerator")
	return nil
}
func (ship *shipLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用ship做PutElephantIntoRefrigerator")
	return nil
}
func (ship *shipLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用ship做CloseTheDoorOfRefrigerator")
	return nil
}
