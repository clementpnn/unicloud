package test

import (
	"backend/database"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDB) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestConnectDB(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		t.Fatalf("Erreur lors de la création du mock : %v", err)
	}
	defer db.Close()

	mock.ExpectPing().WillReturnError(nil)

	err = db.Ping()
	if err != nil {
		t.Fatalf("Ping échoué : %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Il y a des attentes non satisfaites: %v", err)
	}
}

func TestGetEnv(t *testing.T) {
	os.Unsetenv("TEST_ENV_VAR")
	result := database.GetEnv("TEST_ENV_VAR", "default")
	assert.Equal(t, "default", result)

	os.Setenv("TEST_ENV_VAR", "some_value")
	result = database.GetEnv("TEST_ENV_VAR", "default")
	assert.Equal(t, "some_value", result)
}
