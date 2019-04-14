package repositories_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pegasus37/pegasus-peristence/repositories"
	"github.com/stretchr/testify/assert"
)

func TestNewStudent(t *testing.T) {
	t.Run("type of new empty student object", func(t *testing.T) {
		expected := "*repositories.Student"
		student := repositories.NewStudent()

		if reflect.TypeOf(student).String() != expected {
			t.Errorf("expected %s, got %s", expected, reflect.TypeOf(student).String())
		}
	})

	t.Run("created new empty student object", func(t *testing.T) {
		expected := &repositories.Student{}
		student := repositories.NewStudent()

		fmt.Println(reflect.TypeOf(student))
		fmt.Println(reflect.TypeOf(expected))

		assert.Equal(t, *expected, *student)
	})
}

func TestInsert(t *testing.T) {
	t.Run("insert success", func(t *testing.T) {
		student := repositories.NewStudent()

		result, err := student.Insert(student)

		assert.Nil(t, result)
		assert.Nil(t, err)
	})
}
