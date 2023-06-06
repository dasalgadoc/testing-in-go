package test

import (
	"dasalgadoc.com/go-testing/07-deterministic/domain"
	"math/rand"
)

func StudentWithParameters(data map[string]any) *domain.Student {
	dataId, ok := data["id"]
	if !ok {
		id, err := domain.NewStudentId()
		if err != nil {
			panic(err)
		}
		dataId = id.Value()
	}
	studentId, err := domain.NewStudentIdFromString(dataId.(string))
	if err != nil {
		panic(err)
	}
	var studentShift *domain.StudentShift
	dataShift, ok := data["shift"]
	if !ok {
		studentShift, err = domain.NewStudentShiftFromHour(rand.Intn(16) + 6)
		if err != nil {
			panic(err)
		}
	} else {
		studentShift, err = domain.NewStudentShift(dataShift.(string))
		if err != nil {
			panic(err)
		}
	}

	return &domain.Student{
		ID:    studentId,
		Shift: studentShift,
	}
}
