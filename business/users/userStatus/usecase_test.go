package userStatus_test

import (
	// "social_media/business/users"
	"context"
	_middlewares "social_media/app/routes/middleware"
	"social_media/business/users/userStatus"
	_mockUserRepository "social_media/business/users/userStatus/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.Repository
var userService userStatus.Usecase
var userJWT _middlewares.ConfigJWT
var user userStatus.UserStatus

func setup() {
	userService = userStatus.NewUserUscase(&userRepository, time.Hour*1, userJWT)
	user = userStatus.UserStatus{
		ID:     1,
		Status: "hello Faiq",
		IdUser: 2,
	}
}

func TestCreateStatus(t *testing.T) {
	setup()
	userRepository.On("CreateStatus", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int")).Return(user, nil).Once()
	t.Run("Test Case 1| Valid Create Status", func(t *testing.T) {

		users, err := userService.CreateStatus(context.Background(), "Nganjuk", 1)
		assert.Nil(t, err)
		assert.Equal(t, "hello Faiq", users.Status)
	})
	t.Run("Test Case 2| Invalid Create Status", func(t *testing.T) {

		_, err := userService.CreateStatus(context.Background(), "", 1)
		assert.NotNil(t, err)
	})
}
func TestSeeComment(t *testing.T) {
	setup()
	userRepository.On("SeeComment", mock.Anything,
		mock.AnythingOfType("int")).Return([]userStatus.UserStatus{}, nil).Once()
	t.Run("Test Case 1|Valid UserStatus", func(t *testing.T) {
		_, err := userService.SeeComment(context.Background(), 1)
		assert.Nil(t, err)
	})
}
