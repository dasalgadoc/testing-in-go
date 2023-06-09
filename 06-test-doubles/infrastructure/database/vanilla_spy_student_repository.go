package database

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
	"github.com/pkg/errors"
)

type VanillaSpyStudentRepository struct {
	SearchCalled      bool
	SearchCalledTimes int
	SaveCalled        bool
	SaveCalledTimes   int
}

// VanillaSpy checks if the method was called
func NewVanillaSpyStudentRepository() repository.StudentRepository {
	return &VanillaSpyStudentRepository{
		SearchCalled:      false,
		SearchCalledTimes: 0,
		SaveCalled:        false,
		SaveCalledTimes:   0,
	}
}

func (r *VanillaSpyStudentRepository) Save(student domain.Student) error {
	r.SaveCalled = true
	r.SaveCalledTimes += 1
	return nil
}

func (r *VanillaSpyStudentRepository) Search(id domain.StudentId) (domain.Student, error) {
	r.SearchCalled = true
	r.SearchCalledTimes += 1
	return domain.Student{}, errors.New("student not found")
}
