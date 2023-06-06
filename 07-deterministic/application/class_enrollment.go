package application

import (
	"dasalgadoc.com/go-testing/07-deterministic/domain"
	"errors"
	"fmt"
)

type ClassEnrollment struct{}

func NewClassEnrollment() *ClassEnrollment {
	return &ClassEnrollment{}
}

func (c *ClassEnrollment) Enroll(student domain.Student) error {
	if !student.CanTakeClassNow() {
		return errors.New("student can't take class now")
	}
	fmt.Println("student enrolled...")
	return nil
}
