package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
	"tsarka/models"
	"tsarka/repositories"
)

type FakeDB struct {
	mock.Mock
}

func (fdb *FakeDB) QueryRow(query string, args ...interface{}) *sql.Row {
	calledArgs := fdb.Called(append([]interface{}{query}, args...)...)
	return calledArgs.Get(0).(*sql.Row)
}

func (fdb *FakeDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	calledArgs := fdb.Called(append([]interface{}{query}, args...)...)
	return calledArgs.Get(0).(sql.Result), calledArgs.Error(1)
}

func TestCreateUser(t *testing.T) {
	fakeDB := new(FakeDB)
	repo := repositories.NewUserRepository(fakeDB)

	fakeDB.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(&sql.Row{}, nil)
	fakeDB.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	_, err := repo.CreateUser(&user)
	assert.NoError(t, err)

	fakeDB.AssertExpectations(t)
}

func TestMain(m *testing.M) {
	fakeDB := new(FakeDB)

	fakeDB.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(&sql.Row{}, nil)
	fakeDB.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	// Запуск тестов
	exitCode := m.Run()

	// Выход с указанием кода завершения
	os.Exit(exitCode)
}
