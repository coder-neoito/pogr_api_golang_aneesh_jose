package profile_achievements

import (
	"context"
	"testing"

	"github.com/pogr_api_golang_aneesh_jose/profile_achievements/src/mocks"
	"github.com/pogr_api_golang_aneesh_jose/profile_achievements/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_achievementsService_GetUserAchievements(t *testing.T) {
	sampleUserID := "645a242587e0f691ee173e52"

	// assert repository called where objectID is valid and the object type is primitive.ObjectID
	t.Run("with valid object userID string", func(t *testing.T) {
		mockRepository := &mocks.AchievementsRepository{}
		service := achievementsService{
			repository: mockRepository,
		}

		mockRepository.On("GetUserAchievements", mock.Anything, mock.AnythingOfType("primitive.ObjectID")).Once().Return([]models.Achievement{}, nil)
		service.GetUserAchievements(context.TODO(), sampleUserID)
		mockRepository.AssertCalled(t, "GetUserAchievements", mock.Anything, mock.Anything)
	})

	// test for valid object type userID
	t.Run("without valid object userID string", func(t *testing.T) {
		mockRepository := &mocks.AchievementsRepository{}
		service := achievementsService{
			repository: mockRepository,
		}

		_, err := service.GetUserAchievements(context.TODO(), "sampleInvalid")
		assert.Error(t, err)
		mockRepository.AssertNotCalled(t, "GetUserAchievements", mock.Anything, mock.Anything)
	})
}
