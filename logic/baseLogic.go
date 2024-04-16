package logic

import (
	"authenticaiton-authorization/entity"
	"fmt"
)

// ILogic interface tanımı
type ILogic interface {
	Create(model entity.IModel) (*string, error)
	Get(e any)
	Update(e any)
	Delete(e any)
}

type AbstractLogic struct {
	ILogic
}

func (s *AbstractLogic) Create(e any) {
}

func (s *AbstractLogic) Update(e any) {
	fmt.Println("update")
}

func (s *AbstractLogic) Delete(e any) {
	fmt.Println("delete")
}
