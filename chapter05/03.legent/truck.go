package main

import "fmt"

type truckLegend struct {
}

func (truck *truckLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 truckLegend 做OpenTheDoorOfRefrigerator")
	return nil
}
func (truck *truckLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 truckLegend 做PutElephantIntoRefrigerator")
	return nil
}
func (truck *truckLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 truckLegend 做CloseTheDoorOfRefrigerator")
	return nil
}
