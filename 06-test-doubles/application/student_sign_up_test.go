package application

import (
	"dasalgadoc.com/go-testing/06-test-doubles/infrastructure/database"
	"dasalgadoc.com/go-testing/06-test-doubles/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

type studentSignUpTestScenario struct {
	studentName       string
	studentAge        int
	studentVanillaSpy repository.StudentRepository

	err  error
	test *testing.T
}

func TestStudentSignUp_SignUpStudent_Expected_Methods_Called(t *testing.T) {
	s := startStudentSignUpTestScenario(t)
	s.givenAStudentPrimitives("John Doe", 30)
	s.whenSignUpStudentIsCalled()
	s.thenSearchMethodShouldBeCalled(1)
	s.thenSaveMethodShouldBeCalled(1)
}

func startStudentSignUpTestScenario(t *testing.T) *studentSignUpTestScenario {
	t.Parallel()

	return &studentSignUpTestScenario{
		studentVanillaSpy: database.NewVanillaSpyStudentRepository(),
		test:              t,
	}
}

func (s *studentSignUpTestScenario) givenAStudentPrimitives(name string, age int) {
	s.studentName = name
	s.studentAge = age
}

func (s *studentSignUpTestScenario) whenSignUpStudentIsCalled() {
	target := NewStudentSignUp(s.studentVanillaSpy)
	s.err = target.SignUpStudent(s.studentName, s.studentAge)
}

func (s *studentSignUpTestScenario) thenSearchMethodShouldBeCalled(times int) {
	assert.Truef(s.test,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SearchCalled,
		"Search method was not called")

	assert.Equalf(s.test,
		times,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SearchCalledTimes,
		"Search method was not called %d times",
		times)
}

func (s *studentSignUpTestScenario) thenSaveMethodShouldBeCalled(times int) {
	assert.Truef(s.test,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SaveCalled,
		"Save method was not called")

	assert.Equal(s.test,
		times,
		s.studentVanillaSpy.(*database.VanillaSpyStudentRepository).SaveCalledTimes,
		"Save method was not called %d times",
		times)
}
