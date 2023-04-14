package models

import (
	"database/sql/driver"
	"encoding/json"
)

type UIntArr []uint

type User struct {
	Id         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionKey string  `json:"session_key"`
	Openid     string  `json:"openid"`
	IsVip      bool    `json:"is_vip"`
	IdolList   UIntArr `json:"idol_list" gorm:"type:longtext"`
	FanList    UIntArr `json:"fan_list" gorm:"type:longtext"`
}

// 入库。实现 driver.Valuer 接口，Value 返回 json value
func (j UIntArr) Value() (driver.Value, error) {
	if len(j) == 0 {
		return "", nil
	}

	bytes, err := json.Marshal(j)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}

// 出库。实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *UIntArr) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok || len(bytes) == 0 {
		return nil
	}

	result := UIntArr{}
	err := json.Unmarshal(bytes, &result)
	*j = UIntArr(result)
	return err
}
