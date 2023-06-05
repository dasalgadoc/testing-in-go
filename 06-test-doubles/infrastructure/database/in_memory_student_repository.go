package database

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
	"errors"
)

// This is a fake
type InMemoryStudentRepository struct {
	Students map[domain.StudentId]domain.Student
}

func NewInMemoryStudentRepository() repository.StudentRepository {
	return &InMemoryStudentRepository{
		Students: make(map[domain.StudentId]domain.Student),
	}
}

func (r *InMemoryStudentRepository) Save(student domain.Student) error {
	r.Students[*student.ID] = student
	return nil
}

func (r *InMemoryStudentRepository) Search(id domain.StudentId) (domain.Student, error) {
	student, ok := r.Students[id]
	if !ok {
		return domain.Student{}, errors.New("student not found")
	}
	return student, nil
}
