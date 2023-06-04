package entrypoints

import (
	"dasalgadoc.com/go-testing/03-acceptance/application"
	"dasalgadoc.com/go-testing/03-acceptance/infrastructure/dto"
	"github.com/gin-gonic/gin"
)

type StudentGetter struct {
	useCase application.StudentSearcher
}

func NewStudentGetter(useCase application.StudentSearcher) StudentGetter {
	return StudentGetter{useCase: useCase}
}

func (sg StudentGetter) Get(ginCtx *gin.Context) {
	id := ginCtx.Param("student_id")

	student, err := sg.useCase.SearchStudent(id)
	if err != nil {
		ginCtx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	studentDto := dto.StudentDTO{
		ID:   student.ID.Value(),
		Name: student.Name.Value(),
		Age:  student.Age.Value(),
	}

	ginCtx.JSON(200, studentDto)
}
