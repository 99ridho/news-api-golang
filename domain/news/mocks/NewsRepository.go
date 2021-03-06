// Code generated by mockery v1.0.0. DO NOT EDIT.

package newsmocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "gitlab.com/99ridho/news-api/models"

// NewsRepository is an autogenerated mock type for the NewsRepository type
type NewsRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *NewsRepository) Delete(ctx context.Context, id int64) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int64) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchById provides a mock function with given fields: ctx, id
func (_m *NewsRepository) FetchById(ctx context.Context, id int64) (*models.News, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.News
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.News); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByParams provides a mock function with given fields: ctx, params
func (_m *NewsRepository) FetchByParams(ctx context.Context, params *models.FetchNewsParam) ([]*models.News, error) {
	ret := _m.Called(ctx, params)

	var r0 []*models.News
	if rf, ok := ret.Get(0).(func(context.Context, *models.FetchNewsParam) []*models.News); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.FetchNewsParam) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchBySlug provides a mock function with given fields: ctx, slug
func (_m *NewsRepository) FetchBySlug(ctx context.Context, slug string) (*models.News, error) {
	ret := _m.Called(ctx, slug)

	var r0 *models.News
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.News); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByStatus provides a mock function with given fields: ctx, status
func (_m *NewsRepository) FetchByStatus(ctx context.Context, status models.NewsStatus) ([]*models.News, error) {
	ret := _m.Called(ctx, status)

	var r0 []*models.News
	if rf, ok := ret.Get(0).(func(context.Context, models.NewsStatus) []*models.News); ok {
		r0 = rf(ctx, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.NewsStatus) error); ok {
		r1 = rf(ctx, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, _a1
func (_m *NewsRepository) Store(ctx context.Context, _a1 *models.News) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *models.News) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.News) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *NewsRepository) Update(ctx context.Context, _a1 *models.News) (*models.News, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *models.News
	if rf, ok := ret.Get(0).(func(context.Context, *models.News) *models.News); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.News)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.News) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
