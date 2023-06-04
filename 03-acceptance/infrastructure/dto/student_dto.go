package dto

type StudentDTO struct {
	ID   string `json:"student_id"`
	Name string `json:"student_name"`
	Age  int    `json:"student_age"`
}
