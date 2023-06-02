package infrastructure

import (
	"dasalgadoc.com/go-testing/02-integration/domain"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlStudentRepository struct {
	db *sql.DB
}

func NewMysqlStudentRepository(user, password, host, port, database string) *MysqlStudentRepository {
	connectionString := gerConnectionString(user, password, host, port, database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return &MysqlStudentRepository{
		db: db,
	}
}

func gerConnectionString(user, password, host, port, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
}

/*-- Repository --*/
func (r *MysqlStudentRepository) Save(student domain.Student) error {
	stmt, err := r.db.Prepare("INSERT INTO student(id, name, age) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.ID.Value(), student.Name.Value(), student.Age.Value())
	if err != nil {
		return err
	}

	return nil
}

func (r *MysqlStudentRepository) Search(id domain.StudentId) (domain.Student, error) {
	row := r.db.QueryRow("SELECT id, name, age FROM student WHERE id = ?", id.Value())
	var i, s string
	var a int
	err := row.Scan(i, s, a)
	if err != nil {
		return domain.Student{}, err
	}
	uid, err := domain.NewStudentIdFromString(i)
	if err != nil {
		return domain.Student{}, err
	}

	name, err := domain.NewStudentName(s)
	if err != nil {
		return domain.Student{}, err
	}

	age, err := domain.NewStudentAge(a)
	if err != nil {
		return domain.Student{}, err
	}

	var student *domain.Student
	student, err = domain.NewStudentWithId(*uid, *name, *age)
	if err != nil {
		return domain.Student{}, err
	}

	return *student, nil
}
