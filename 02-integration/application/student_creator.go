package application

import (
	"dasalgadoc.com/go-testing/02-integration/domain"
	"dasalgadoc.com/go-testing/02-integration/repository"
)

type StudentCreator struct {
	studentRepo repository.StudentRepository
}

func NewStudentCreator(studentRepo repository.StudentRepository) StudentCreator {
	return StudentCreator{studentRepo: studentRepo}
}

func (sc StudentCreator) CreateStudent(name string, age int) error {
	studentName, err := domain.NewStudentName(name)
	if err != nil {
		return err
	}

	studentAge, err := domain.NewStudentAge(age)
	if err != nil {
		return err
	}

	student, err := domain.NewStudent(*studentName, *studentAge)
	if err != nil {
		return err
	}

	return sc.studentRepo.Save(*student)
}
