package utils

import "fmt"

type UniqueException struct {
	Message string
}

func (e *UniqueException) Error() string {
	return fmt.Sprintf("%s already exist.", e.Message)
}

type BusinessException struct {
	Message string
}

func (e *BusinessException) Error() string {
	return e.Message
}

type NotNullException struct {
	Message string
}

func (e *NotNullException) Error() string {
	return fmt.Sprintf("%s is not empty.", e.Message)
}
