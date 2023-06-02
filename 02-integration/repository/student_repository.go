package repository

import "dasalgadoc.com/go-testing/02-integration/domain"

type StudentRepository interface {
	Save(student domain.Student) error
	Search(id domain.StudentId) (domain.Student, error)
}
