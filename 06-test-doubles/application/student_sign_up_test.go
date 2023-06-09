package application

import (
	"dasalgadoc.com/go-testing/06-test-doubles/domain"
	"dasalgadoc.com/go-testing/06-test-doubles/infrastructure/database"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type studentSignUpTestScenario struct {
	student           domain.Student
	studentName       string
	studentAge        int
	studentVanillaSpy repository.StudentRepository
	studentMockSpy    *database.MockStudentRepository

	err  error
	test *testing.T
}

func TestStudentSignUp_SignUpStudent_Expected_Methods_Called(t *testing.T) {
	s := startStudentSignUpTestScenario(t)
	s.givenAStudentPrimitives("John Doe", 30)
	s.whenSignUpStudentIsCalledWithVanillaSpy()
	s.thenSearchMethodShouldBeCalled_VanillaCheck(1)
	s.thenSaveMethodShouldBeCalled_VanillaCheck(1)
}

func TestStudentSignUp_SignUpStudent_Expected_Methods_Called_Mock(t *testing.T) {
	s := startStudentSignUpTestScenario(t)
	s.givenAStudent("John Doe", 30)
	s.andStudentMockRepositoryIsNotFound()
	s.andStudentMockRepositorySaveStudent()
	s.whenSignUpStudentIsCalledWithMock()
	s.thenSearchMethodShouldBeCalled_MockCheck(1)
	s.thenSaveMethodShouldBeCalled_MockCheck(1)
}

func startStudentSignUpTestScenario(t *testing.T) *studentSignUpTestScenario {
	t.Parallel()

	return &studentSignUpTestScenario{
		studentVanillaSpy: database.NewVanillaSpyStudentRepository(),
		studentMockSpy:    new(database.MockStudentRepository),
		test:              t,
	}
}

func (s *studentSignUpTestScenario) givenAStudent(name string, age int) {
	s.studentName = name
	s.studentAge = age

	uid, err := domain.NewStudentId()
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

func (s *studentSignUpTestScenario) givenAStudentPrimitives(name string, age int) {
	s.studentName = name
	s.studentAge = age
}

func (s *studentSignUpTestScenario) andStudentMockRepositoryFoundStudent() {
	s.studentMockSpy.
		On("Search", mock.Anything).Return(s.student, nil).
		Once()
}

func (s *studentSignUpTestScenario) andStudentMockRepositoryIsNotFound() {
	s.studentMockSpy.
		On("Search", mock.Anything).Return(domain.Student{}, errors.New("something went wrong")).
		Once()
}

func (s *studentSignUpTestScenario) andStudentMockRepositorySaveStudent() {
	s.studentMockSpy.
		On("Save", mock.Anything).Return(nil).
		Once()
}

func (s *studentSignUpTestScenario) whenSignUpStudentIsCalledWithVanillaSpy() {
	target := NewStudentSignUp(s.studentVanillaSpy)
	s.err = target.SignUpStudent(s.studentName, s.studentAge)
}

func (s *studentSignUpTestScenario) whenSignUpStudentIsCalledWithMock() {
	target := NewStudentSignUp(s.studentMockSpy)
	s.err = target.SignUpStudent(s.studentName, s.studentAge)
}

func (s *studentSignUpTestScenario) thenSearchMethodShouldBeCalled_VanillaCheck(times int) {
	assert.Truef(s.test,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SearchCalled,
		"Search method was not called")

	assert.Equalf(s.test,
		times,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SearchCalledTimes,
		"Search method was not called %d times",
		times)
}

func (s *studentSignUpTestScenario) thenSaveMethodShouldBeCalled_VanillaCheck(times int) {
	assert.Truef(s.test,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SaveCalled,
		"Save method was not called")

	assert.Equal(s.test,
		times,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SaveCalledTimes,
		"Save method was not called %d times",
		times)
}

func (s *studentSignUpTestScenario) thenSearchMethodShouldBeCalled_MockCheck(times int) {
	s.studentMockSpy.AssertCalled(s.test, "Search", mock.Anything)
	s.studentMockSpy.AssertNumberOfCalls(s.test, "Search", times)
}

func (s *studentSignUpTestScenario) thenSaveMethodShouldBeCalled_MockCheck(times int) {
	s.studentMockSpy.AssertCalled(s.test, "Search", mock.Anything)
	s.studentMockSpy.AssertNumberOfCalls(s.test, "Save", times)
}
