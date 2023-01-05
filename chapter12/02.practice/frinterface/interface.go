package frinterface

import "Learning-JobAccess-Camp/pkg/apis"

type ServeInterface interface {
	RegisterPersonalInformation(pi *apis.PersonalInformation) error
	UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error)
	GetFatRate(name string) (*apis.PersonalRank, error)
	GetTop() ([]*apis.PersonalRank, error)
}
