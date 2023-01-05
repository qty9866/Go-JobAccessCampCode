package main

import (
	"Learning-JobAccess-Camp/chapter12/02.practice/frinterface"
	"Learning-JobAccess-Camp/pkg/apis"
)

var _ frinterface.ServeInterface = &dbRank{}

type dbRank struct {
}

func (d dbRank) RegisterPersonalInformation(pi *apis.PersonalInformation) error {
	//TODO implement me
	panic("implement me")
}

func (d dbRank) UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbRank) GetFatRate(name string) (*apis.PersonalRank, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbRank) GetTop() ([]*apis.PersonalRank, error) {
	//TODO implement me
	panic("implement me")
}
