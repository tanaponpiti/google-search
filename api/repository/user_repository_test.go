package repository_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"server-side/model"
	"server-side/repository"
	"server-side/utility"
	"testing"
)

func setupDatabase() *gorm.DB {
	uniqueID := uuid.New().String()
	connectionString := "file:" + uniqueID + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.SearchResult{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func TestUserRepositoryInsertAndGetUser(t *testing.T) {
	db := setupDatabase()
	repo := repository.NewUserRepository(db)
	require.NoError(t, repo.Migrate())
	userCreateInput := model.UserCreate{
		Username: "testUser",
		Password: "password123",
		Name:     "Test User",
	}
	user, err := repo.InsertUser(userCreateInput)
	require.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.Equal(t, userCreateInput.Username, user.Username)
	assert.NotEqual(t, userCreateInput.Password, user.Password, "Password should be hashed")
	assert.True(t, utility.CheckPasswordHash(userCreateInput.Password, user.Password))
	assert.Equal(t, userCreateInput.Name, user.Name)
	fetchedUser, err := repo.GetUser(user.ID)
	require.NoError(t, err)
	assert.Equal(t, user.Username, fetchedUser.Username)
	assert.Equal(t, user.Name, fetchedUser.Name)
}

func TestUserRepositoryGetUserByUsername(t *testing.T) {
	db := setupDatabase()
	repo := repository.NewUserRepository(db)
	require.NoError(t, repo.Migrate())
	userCreateInput := model.UserCreate{
		Username: "uniqueUser",
		Password: "password123",
		Name:     "Unique User",
	}
	repo.InsertUser(userCreateInput)
	fetchedUser, err := repo.GetUserByUsername(userCreateInput.Username)
	require.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, userCreateInput.Username, fetchedUser.Username)
}

func TestUserRepositoryUpdateUser(t *testing.T) {
	db := setupDatabase()
	repo := repository.NewUserRepository(db)
	require.NoError(t, repo.Migrate())
	userCreateInput := model.UserCreate{
		Username: "updateUser",
		Password: "oldPassword",
		Name:     "Old Name",
	}
	user, _ := repo.InsertUser(userCreateInput)

	newName := "New Name"
	newPassword := "newPassword123"
	updateData := model.UserUpdate{
		Name:     &newName,
		Password: &newPassword,
	}
	updatedUser, err := repo.UpdateUser(user.ID, updateData)
	require.NoError(t, err)
	assert.Equal(t, newName, updatedUser.Name)
	assert.True(t, utility.CheckPasswordHash(newPassword, updatedUser.Password))
}

func TestUserRepositoryDeleteUser(t *testing.T) {
	db := setupDatabase()
	repo := repository.NewUserRepository(db)
	require.NoError(t, repo.Migrate())
	userCreateInput := model.UserCreate{
		Username: "deleteUser",
		Password: "password123",
		Name:     "Delete User",
	}
	user, _ := repo.InsertUser(userCreateInput)
	err := repo.DeleteUser(user.ID)
	require.NoError(t, err)
	foundUser, err := repo.GetUser(user.ID)
	assert.Nil(t, foundUser)
}
