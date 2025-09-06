package model

import "fmt"

type User struct {
	Id     string  `json:"id"`
	Tid    int64   `json:"tid"`
	Weight float32 `json:"weight"`
	Height float32 `json:"height"`
}

func (u *User) Recipient() string {
	return fmt.Sprint(u.Tid)
}
