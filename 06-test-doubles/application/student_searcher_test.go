package application

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/infrastructure/database"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type studentSearcherTestScenario struct {
	student         domain.Student
	retrieveStudent domain.Student
	studentFake     repository.StudentRepository
	studentStub     repository.StudentRepository
	studentMock     *database.MockStudentRepository

	err  error
	test *testing.T
}

func TestStudentSearcher_SearchStudentOnFakeRepo(t *testing.T) {
	s := startStudentSearcherTestScenario(t)
	s.givenStudent("45215570-0296-11ee-8566-acde48001122", "John Doe", 30)
	s.andStudentIsOnFakeRepository()
	s.whenSearchStudentWithFakeRepo("45215570-0296-11ee-8566-acde48001122")

	assert.NoError(s.test, s.err)
	assert.Equal(s.test, s.student, s.retrieveStudent)
}

func TestStudentSearcher_SearchStudentOnStubRepo(t *testing.T) {
	s := startStudentSearcherTestScenario(t)
	s.givenStudent("45215570-0296-11ee-8566-acde48001122", "John Doe", 30)
	s.whenSearchStudentWithStubRepo("45215570-0296-11ee-8566-acde48001122")

	assert.NoError(s.test, s.err)
	assert.Equal(s.test, s.student, s.retrieveStudent)
}

func TestStudentSearcher_SearchStudentMockRepo(t *testing.T) {
	s := startStudentSearcherTestScenario(t)
	s.givenStudent("45215570-0296-11ee-8566-acde48001122", "John Doe", 30)
	s.andStudentMockRepositoryIsOk()
	s.whenSearchStudentWithMockRepo("45215570-0296-11ee-8566-acde48001122")

	assert.NoError(s.test, s.err)
	assert.Equal(s.test, s.student, s.retrieveStudent)
}

/*-- Steps --*/
func startStudentSearcherTestScenario(t *testing.T) *studentSearcherTestScenario {
	t.Parallel()

	return &studentSearcherTestScenario{
		studentFake: database.NewInMemoryStudentRepository(),
		studentStub: database.NewStubStudentRepository(),
		studentMock: new(database.MockStudentRepository),
		test:        t,
	}
}

func (s *studentSearcherTestScenario) givenStudent(id string, name string, age int) {
	uid, err := domain.NewStudentIdFromString(id)
	if err != nil {
		s.test.Fatal(err)
	}
	studentName, err := domain.NewStudentName(name)
	if err != nil {
		s.test.Fatal(err)
	}
	studentAge, err := domain.NewStudentAge(age)
	if err != nil {
		s.test.Fatal(err)
	}

	student, err := domain.NewStudentWithId(*uid, *studentName, *studentAge)
	if err != nil {
		s.test.Fatal(err)
	}

	s.student = *student
}

func (s *studentSearcherTestScenario) andStudentIsOnFakeRepository() {
	err := s.studentFake.Save(s.student)
	if err != nil {
		s.test.Fatal(err)
	}
}

func (s *studentSearcherTestScenario) andStudentMockRepositoryIsOk() {
	s.studentMock.On("Search", mock.Anything).Return(s.student, nil).Once()
}

func (s *studentSearcherTestScenario) andStudentMockRepositoryIsNotFound() {
	s.studentMock.On("Search", mock.Anything).Return(domain.Student{}, errors.New("something went wrong")).Once()
}

func (s *studentSearcherTestScenario) whenSearchStudentWithFakeRepo(id string) {
	target := NewStudentSearcher(s.studentFake)
	s.retrieveStudent, s.err = target.SearchStudent(id)
}

func (s *studentSearcherTestScenario) whenSearchStudentWithStubRepo(id string) {
	target := NewStudentSearcher(s.studentStub)
	s.retrieveStudent, s.err = target.SearchStudent(id)
}

func (s *studentSearcherTestScenario) whenSearchStudentWithMockRepo(id string) {
	target := NewStudentSearcher(s.studentMock)
	s.retrieveStudent, s.err = target.SearchStudent(id)
}
