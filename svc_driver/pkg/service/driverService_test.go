package service

import (
	"context"
	"github.com/d-Una-Interviews/svc_driver/pkg/model"
	p "github.com/gobeam/mongo-go-pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

type driverRepositoryMock struct {
	mock.Mock
}

func (d *driverRepositoryMock) FindByRadius(ctx context.Context, radius int) (error, []*model.Driver) {
	args := d.Called(ctx, radius)
	return args.Error(0), args.Get(1).([]*model.Driver)
}

func (d *driverRepositoryMock) FindAll(ctx context.Context) (error, []*model.Driver) {
	args := d.Called(ctx)
	return args.Error(0), args.Get(1).([]*model.Driver)
}

func (d *driverRepositoryMock) FindAllPaged(ctx context.Context, limit int64, page int64) (error error, paged map[string]interface{}) {
	args := d.Called(ctx, limit, page)
	return args.Error(0), args.Get(1).(map[string]interface{})
}

func (d *driverRepositoryMock) Create(ctx context.Context, driver *model.Driver) error {
	args := d.Called()
	return args.Error(0)
}

func TestFindByRadiusWithoutResults(t *testing.T) {
	var drivers []*model.Driver
	modelDriverRepositoryMock := driverRepositoryMock{}
	todo := context.TODO()
	modelDriverRepositoryMock.On("FindByRadius", todo, 2).Return(nil, drivers)

	service := driverService{
		Repo:      &modelDriverRepositoryMock,
		Validator: validator.New(),
	}
	_, driversResponse := service.FindByRadius(todo, 2)
	assert.Equal(t, drivers, driversResponse)
	assert.Equal(t, len(drivers), 0)
}

func TestFindByRadiusWithResults(t *testing.T) {
	var drivers []*model.Driver
	var driver = model.Driver{}
	drivers = append(drivers, &driver)

	modelDriverRepositoryMock := driverRepositoryMock{}
	todo := context.TODO()
	modelDriverRepositoryMock.On("FindByRadius", todo, 2).Return(nil, drivers)

	service := driverService{
		Repo:      &modelDriverRepositoryMock,
		Validator: validator.New(),
	}
	_, driversResponse := service.FindByRadius(todo, 2)
	assert.Equal(t, drivers, driversResponse)
	assert.Equal(t, len(drivers), 1)
}

func TestFindAllWithoutResults(t *testing.T) {
	var drivers []*model.Driver
	modelDriverRepositoryMock := driverRepositoryMock{}
	todo := context.TODO()
	modelDriverRepositoryMock.On("FindAll", todo).Return(nil, drivers)

	service := driverService{
		Repo:      &modelDriverRepositoryMock,
		Validator: validator.New(),
	}
	_, driversResponse := service.FindAll(todo)
	assert.Equal(t, drivers, driversResponse)
	assert.Equal(t, len(drivers), 0)
}

func TestFindAllResults(t *testing.T) {
	var drivers []*model.Driver
	var driver = model.Driver{}
	drivers = append(drivers, &driver)

	modelDriverRepositoryMock := driverRepositoryMock{}
	todo := context.TODO()
	modelDriverRepositoryMock.On("FindAll", todo).Return(nil, drivers)

	service := driverService{
		Repo:      &modelDriverRepositoryMock,
		Validator: validator.New(),
	}
	_, driversResponse := service.FindAll(todo)
	assert.Equal(t, drivers, driversResponse)
	assert.Equal(t, len(drivers), 1)
}

func TestFindAllPaged(t *testing.T) {
	var drivers []model.Driver
	var driver = model.Driver{}
	var paginatedData p.PaginatedData

	drivers = append(drivers, driver)

	var resp = map[string]interface{}{}
	resp["drivers"] = drivers //Store the token in the response
	resp["paging"] = paginatedData.Pagination
	limit := int64(2)
	page := int64(1)

	modelDriverRepositoryMock := driverRepositoryMock{}
	todo := context.TODO()
	modelDriverRepositoryMock.On("FindAllPaged", todo, limit, page).Return(nil, resp)

	service := driverService{
		Repo:      &modelDriverRepositoryMock,
		Validator: validator.New(),
	}
	_, serviceResponse := service.FindAllPaged(todo, limit, page)
	assert.Equal(t, serviceResponse, resp)
	assert.Equal(t, serviceResponse["drivers"], drivers)
}
