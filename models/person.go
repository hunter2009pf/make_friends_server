package models

import (
	"database/sql/driver"
	"encoding/json"
)

type StrArr []string

type Person struct {
	Id                uint   `json:"id"`
	Name              string `json:"name"`
	AvatarUrl         string `json:"avatar_url" gorm:"default:''"`
	IsBoy             bool   `json:"is_boy"`
	Age               uint   `json:"age"`
	Height            uint   `json:"height"`
	Weight            uint   `json:"weight"`
	Star              string `json:"star"`
	BirthPlace        string `json:"birth_place"`
	CurrentPlace      string `json:"current_place"`
	MarryStatus       string `json:"marry_status"`
	Degree            string `json:"degree"`
	Job               string `json:"job"`
	MonthSalary       string `json:"month_salary"`
	HouseStatus       string `json:"house_status"`
	CarStatus         string `json:"car_status"`
	UniqueChildStatus string `json:"unique_child_status"`
	Wechat            string `json:"wechat"`
	Photos            StrArr `json:"photos"`
}

func (t *StrArr) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t StrArr) Value() (driver.Value, error) {
	return json.Marshal(t)
}
