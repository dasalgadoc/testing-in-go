package database

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
)

type DummyStudentRepository struct{}

// Dummies has no responses
func NewDummyStudentRepository() repository.StudentRepository {
	return &DummyStudentRepository{}
}

func (r *DummyStudentRepository) Save(student domain.Student) error {
	return nil
}

func (r *DummyStudentRepository) Search(id domain.StudentId) (domain.Student, error) {
	return domain.Student{}, nil
}
