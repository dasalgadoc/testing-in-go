package application

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
	"errors"
)

type StudentSignUp struct {
	studentRepo repository.StudentRepository
}

func NewStudentSignUp(studentRepo repository.StudentRepository) StudentSignUp {
	return StudentSignUp{studentRepo: studentRepo}
}

func (ss StudentSignUp) SignUpStudent(name string, age int) error {
	studentId, err := domain.NewStudentId()
	if err != nil {
		return err
	}

	_, err = ss.studentRepo.Search(*studentId)
	if err == nil {
		return errors.New("student already exists")
	}

	studentName, err := domain.NewStudentName(name)
	if err != nil {
		return err
	}

	studentAge, err := domain.NewStudentAge(age)
	if err != nil {
		return err
	}

	student, err := domain.NewStudentWithId(*studentId, *studentName, *studentAge)

	return ss.studentRepo.Save(*student)
}
