package storage

import "github.com/Harsh971/GoLang_REST_API_CRUD/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int, error)
	GetStudentById(id int) (types.Student, error)
}
