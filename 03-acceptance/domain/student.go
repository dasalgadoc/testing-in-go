package domain

type Student struct {
	ID   *StudentId
	Name *StudentName
	Age  *StudentAge
}

func NewStudent(name StudentName, age StudentAge) (*Student, error) {
	id, err := NewStudentId()
	if err != nil {
		return nil, err
	}
	return &Student{ID: id, Name: &name, Age: &age}, nil
}

func NewStudentWithId(id StudentId, name StudentName, age StudentAge) (*Student, error) {
	return &Student{ID: &id, Name: &name, Age: &age}, nil
}
