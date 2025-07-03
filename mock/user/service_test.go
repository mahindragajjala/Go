// user/service_test.go
package user_test

import (
	"testing"

	"user"
	"user/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIsEmailRegistered(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := user.NewUserService(mockRepo)

	t.Run("email exists", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail("test@example.com").
			Return(&user.User{ID: "1", Email: "test@example.com"}, nil)

		result := service.IsEmailRegistered("test@example.com")
		assert.True(t, result)
	})

	t.Run("email does not exist", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByEmail("missing@example.com").
			Return(nil, nil)

		result := service.IsEmailRegistered("missing@example.com")
		assert.False(t, result)
	})
}
