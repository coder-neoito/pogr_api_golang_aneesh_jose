// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/pogr_api_golang_aneesh_jose/profile_achievements/src/models"
	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// AchievementsRepository is an autogenerated mock type for the AchievementsRepository type
type AchievementsRepository struct {
	mock.Mock
}

// GetUserAchievements provides a mock function with given fields: ctx, userID
func (_m *AchievementsRepository) GetUserAchievements(ctx context.Context, userID primitive.ObjectID) ([]models.Achievement, error) {
	ret := _m.Called(ctx, userID)

	var r0 []models.Achievement
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) ([]models.Achievement, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) []models.Achievement); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Achievement)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAchievementsRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAchievementsRepository creates a new instance of AchievementsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAchievementsRepository(t mockConstructorTestingTNewAchievementsRepository) *AchievementsRepository {
	mock := &AchievementsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}