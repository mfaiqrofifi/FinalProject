package userRegister_test

import (
	// "social_media/business/users"
	"context"
	_middlewares "social_media/app/routes/middleware"
	"social_media/business/users/userRegister"
	_mockUserRepository "social_media/business/users/userRegister/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.Repository
var userService userRegister.Usecase
var userJWT _middlewares.ConfigJWT
var user userRegister.DomainRegister

func setup() {
	userService = userRegister.NewUserUscase(&userRepository, time.Hour*1, userJWT)
	user = userRegister.DomainRegister{
		ID:       1,
		Name:     "Fifi",
		Email:    "mfaiqrofifi@gmail.com",
		Password: "abcdefg",
		Address:  "bangsri",
		Token:    "123sdf",
	}
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(user, nil).Once()
	t.Run("Test Case 1| Valid Register", func(t *testing.T) {

		users, err := userService.Register(context.Background(), "Faiq", "mfaiqrofifi@gmail.com", "Nganjuk", "12323da")
		assert.Nil(t, err)
		assert.Equal(t, "", users.Name)
	})
	t.Run("Test Case 2| Invalid Name", func(t *testing.T) {

		_, err := userService.Register(context.Background(), "", "mfaiqrofifi@gmail.com", "Nganjuk", "12323da")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3| Invalid Email", func(t *testing.T) {

		_, err := userService.Register(context.Background(), "Faiq", "", "Nganjuk", "12323da")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3| Invalid Password", func(t *testing.T) {

		_, err := userService.Register(context.Background(), "Faiq", "mfaiqrofifi@gmail.com", "Nganjuk", "")
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	setup()
	user.Password,_=encrypt.Hash(user.Password)
	userRepository.On("Login", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(user, nil).Once()
	t.Run("Test Case 1| Valid Login", func(t *testing.T) {
		users, err := userService.Login(context.Background(), "mfaiqrofifi@gmail.com", "123sdf")
		assert.Nil(t, err)
		assert.Equal(t, "Fifi", users.Name)
	})
	t.Run("Test Case 2| Invalid Email", func(t *testing.T) {
		_, err := userService.Login(context.Background(), "", "123sdf")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3| Invalid Password", func(t *testing.T) {
		_, err := userService.Login(context.Background(), "mfaiqrofifi@gmail.com", "")
		assert.NotNil(t, err)
	})
	t.Run("Test Case 4| Wrong Password",func(t *testing){
		userRepository.On("Login", 
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(user,business.ErrPassword).Once()
		_,err := userService.Login(context.Background(),"mfaiqrofifi@gmail.com","4532")
		assert.NotNil(t,err)
	})
}

func TestDeteilRegister(t *testing.T) {
	setup()
	userRepository.On("DeteilRegister", mock.Anything,
		mock.AnythingOfType("int")).Return([]userRegister.DomainRegister{}, nil).Once()
	t.Run("Test Case 1| Valid Deteil Register", func(t *testing.T) {
		_, err := userService.DeteilRegister(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func TestGetFriend(t *testing.T) {
	setup()
	userRepository.On("GetFriend", mock.Anything,
		mock.AnythingOfType("string")).Return([]userRegister.DomainRegister{}, nil).Once()
	t.Run("Test Case 1| Valid GetFriend", func(t *testing.T) {
		_, err := userService.GetFriend(context.Background(), "faiq")
		assert.Nil(t, err)
	})
	t.Run("Test Case 2| Invalid GetFriend", func(t *testing.T) {
		_, err := userService.GetFriend(context.Background(), "")
		assert.NotNil(t, err)
	})
}
