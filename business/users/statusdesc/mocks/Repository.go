// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	statusdesc "social_media/business/users/statusdesc"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// MakingComment provides a mock function with given fields: ctx, comment, userId, statusId
func (_m *Repository) MakingComment(ctx context.Context, comment string, userId int, statusId int) (statusdesc.DetailStatus, error) {
	ret := _m.Called(ctx, comment, userId, statusId)

	var r0 statusdesc.DetailStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) statusdesc.DetailStatus); ok {
		r0 = rf(ctx, comment, userId, statusId)
	} else {
		r0 = ret.Get(0).(statusdesc.DetailStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) error); ok {
		r1 = rf(ctx, comment, userId, statusId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakingLike provides a mock function with given fields: ctx, userId, statusId
func (_m *Repository) MakingLike(ctx context.Context, userId int, statusId int) (statusdesc.DetailStatus, error) {
	ret := _m.Called(ctx, userId, statusId)

	var r0 statusdesc.DetailStatus
	if rf, ok := ret.Get(0).(func(context.Context, int, int) statusdesc.DetailStatus); ok {
		r0 = rf(ctx, userId, statusId)
	} else {
		r0 = ret.Get(0).(statusdesc.DetailStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, userId, statusId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
