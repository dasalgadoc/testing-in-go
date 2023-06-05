package database

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
)

type StubStudentRepository struct{}

// Stubs has fixed responses
func NewStubStudentRepository() repository.StudentRepository {
	return &StubStudentRepository{}
}

func (r *StubStudentRepository) Save(student domain.Student) error {
	return nil
}

func (r *StubStudentRepository) Search(id domain.StudentId) (domain.Student, error) {
	studentName, _ := domain.NewStudentName("John Doe")
	studentAge, _ := domain.NewStudentAge(30)

	return domain.Student{
		ID:   &id,
		Name: studentName,
		Age:  studentAge,
	}, nil
}
