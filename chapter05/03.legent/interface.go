package main

type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}
type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}

func (Refrigerator) Open() error {
	return nil
}
func (Refrigerator) Close() error {
	return nil
}
