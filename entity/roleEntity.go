package entity

import (
	"encoding/json"
	"fmt"
)

type Role struct {
	*AbstractEntity
	Name string `json:"name" db:"name"`
}
type RoleSlice []Role

func (rs *RoleSlice) Scan(src interface{}) error {
	if src == nil {
		*rs = make([]Role, 0)
		return nil
	}
	bytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("%T", src)
	}
	err := json.Unmarshal(bytes, rs)
	if err != nil {
		return err
	}
	return nil
}
