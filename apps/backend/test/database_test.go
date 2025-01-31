package test

import (
	"backend/database"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock pour simuler un environnement de base de données
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
	// Création d'un mock de la DB avec l'option pour surveiller les pings
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // Active la surveillance des pings
	if err != nil {
		t.Fatalf("Erreur lors de la création du mock : %v", err)
	}
	defer db.Close()

	// On configure le mock pour simuler un appel à Ping()
	mock.ExpectPing().WillReturnError(nil) // Aucune erreur lors du ping

	// Vérification du comportement de la méthode Ping sur le mock
	err = db.Ping()
	if err != nil {
		t.Fatalf("Ping échoué : %v", err)
	}

	// Vérifie que toutes les attentes du mock ont été satisfaites
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Il y a des attentes non satisfaites: %v", err)
	}
}

func TestGetEnv(t *testing.T) {
	// Test sans définir de variable d'environnement
	os.Unsetenv("TEST_ENV_VAR")
	result := database.GetEnv("TEST_ENV_VAR", "default")
	assert.Equal(t, "default", result)

	// Test avec une variable d'environnement définie
	os.Setenv("TEST_ENV_VAR", "some_value")
	result = database.GetEnv("TEST_ENV_VAR", "default")
	assert.Equal(t, "some_value", result)
}
