package storage

import "github.com/Harsh971/GoLang_REST_API_CRUD/internal/types"

type Storage interface {
	CreateStudent(name, email string, age int) (int, error)
	GetStudentByID(id int) (types.Student, error)
	GetStudentList() ([]types.Student, error)
}
