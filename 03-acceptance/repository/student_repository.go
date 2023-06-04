package repository

import "dasalgadoc.com/go-testing/03-acceptance/domain"

type StudentRepository interface {
	Save(student domain.Student) error
	Search(id domain.StudentId) (domain.Student, error)
}
