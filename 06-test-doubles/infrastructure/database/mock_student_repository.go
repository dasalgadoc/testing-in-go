package database

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"github.com/stretchr/testify/mock"
)

type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) Save(student domain.Student) error {
	results := m.Called(student)
	return results.Error(0)
}

func (m *MockStudentRepository) Search(id domain.StudentId) (domain.Student, error) {
	results := m.Called(id)
	return results.Get(0).(domain.Student), results.Error(1)
}
