package main

type PersonalInformation struct {
	Id     int64   `json:"id,omitempty" gorm:"primaryKey;column:ID"`
	Name   string  `json:"name,omitempty" gorm:"column:NAME"`
	Sex    string  `json:"sex,omitempty" gorm:"column:SEX"`
	Tall   float32 `json:"tall,omitempty" gorm:"column:TALL"`
	Weight float32 `json:"weight,omitempty" gorm:"column:WEIGHT"`
	Age    int64   `json:"age,omitempty" gorm:"column:AGE"`
}

func (*PersonalInformation) TableName() string {
	return "personal_information"
}
