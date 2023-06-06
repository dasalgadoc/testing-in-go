package application

import (
	"dasalgadoc.com/go-testing/07-deterministic/domain"
	"dasalgadoc.com/go-testing/07-deterministic/test"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

/* Deterministic test*/
func TestMorningStudentCanTakeClassInMorning(t *testing.T) {
	testData := map[string]any{
		"shift": "morning",
	}

	/* Mocking */
	domain.CurrentHour = func() int {
		return generateRandomNumber(6, 11)
	}

	student := test.StudentWithParameters(testData)
	classEnrollment := NewClassEnrollment()

	assert.NoError(t, classEnrollment.Enroll(*student))
}

func TestMorningStudentCantTakeClassAtNight(t *testing.T) {
	testData := map[string]any{
		"shift": "morning",
	}

	/* Mocking */
	domain.CurrentHour = func() int {
		return generateRandomNumber(18, 21)
	}

	student := test.StudentWithParameters(testData)
	classEnrollment := NewClassEnrollment()

	assert.Error(t, classEnrollment.Enroll(*student))
}

func generateRandomNumber(startHour, endHour int) int {
	return rand.Intn(endHour-startHour+1) + startHour
}
