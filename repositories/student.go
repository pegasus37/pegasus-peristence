package repositories

type StudentInterface interface {
	Insert(student *Student) (*Student, error)
	// BulkInsert(studentList []*Student) ([]*Student, error)
}

type Student struct {
	name string
}

func NewStudent() *Student {
	return &Student{}
}

func (studentModel Student) Insert(student *Student) (*Student, error) {

	return nil, nil
}

// func (student Student) BulkInsert(studentList []*Student) ([]*Student, error) {
// }
