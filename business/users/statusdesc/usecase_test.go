package statusdesc_test

import (
	// "social_media/business/users"
	"context"
	// _middlewares "social_media/app/routes/middleware"
	"social_media/business/users/statusdesc"
	_mockUserRepository "social_media/business/users/statusdesc/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.Repository
var userService statusdesc.Usecase

// var userJWT _middlewares.ConfigJWT
var user statusdesc.DetailStatus

func setup() {
	userService = statusdesc.NewUserUscase(&userRepository, time.Hour*1)
	user = statusdesc.DetailStatus{
		ID:       1,
		Coment:   "hallo",
		UserID:   12,
		Like:     4,
		Dislike:  4,
		StatusId: 17,
	}
}

func TestMakingComment(t *testing.T) {
	setup()
	userRepository.On("MakingComment", mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
		mock.AnythingOfType("int")).Return(user, nil).Once()
	t.Run("Test Case 1| Valid Making Comment", func(t *testing.T) {

		users, err := userService.MakingComment(context.Background(), "Altera", 1, 2)
		assert.Nil(t, err)
		assert.Equal(t, "hallo", users.Coment)
	})
	t.Run("Test Case 2| Invalid Making Comment", func(t *testing.T) {

		_, err := userService.MakingComment(context.Background(), "", 1, 2)
		assert.NotNil(t, err)
	})

}

func TestMakingLike(t *testing.T) {
	setup()
	userRepository.On("MakingLike", mock.Anything,
		mock.AnythingOfType("int"),
		mock.AnythingOfType("int")).Return(user, nil).Once()
	t.Run("Test Case 1| Valid Making Like", func(t *testing.T) {

		users, err := userService.MakingLike(context.Background(), 1, 2)
		assert.Nil(t, err)
		assert.Equal(t, "hallo", users.Coment)
	})
}
