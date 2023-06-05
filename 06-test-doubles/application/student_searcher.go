package application

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
)

type StudentSearcher struct {
	studentRepo repository.StudentRepository
}

func NewStudentSearcher(studentRepo repository.StudentRepository) StudentSearcher {
	return StudentSearcher{studentRepo: studentRepo}
}

func (ss StudentSearcher) SearchStudent(id string) (domain.Student, error) {
	studentID, err := domain.NewStudentIdFromString(id)
	if err != nil {
		return domain.Student{}, err
	}

	return ss.studentRepo.Search(*studentID)
}
