package service

import (
	"context"
	"github.com/d-Una-Interviews/svc_driver/pkg/model"
	//"github.com/d-Una-Interviews/svc_driver/pkg/repository"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

var (
	driverServiceLogger, _ = zap.NewProduction(zap.Fields(zap.String("type", "service")))
)

type DriverService interface {
	FindAll(ctx context.Context) (error error, drivers []*model.Driver)
	FindAllPaged(ctx context.Context, limit int64, page int64) (error error, paged map[string]interface{})
	FindByRadius(ctx context.Context, radius int) (error error, drivers []*model.Driver)
	Create(ctx context.Context, driver *model.Driver) error
}

type driverService struct {
	Repo      model.DriverRepository
	Validator *validator.Validate
}

func (d driverService) FindAll(ctx context.Context) (error error, drivers []*model.Driver) {
	err, drivers := d.Repo.FindAll(ctx)
	if err != nil {
		driverServiceLogger.Error(err.Error(), zap.Error(err))
	}
	return err, drivers
}

func (d driverService) FindAllPaged(ctx context.Context, limit int64, page int64) (error error, paged map[string]interface{}) {
	err, driversPaged := d.Repo.FindAllPaged(ctx, limit, page)
	if err != nil {
		driverServiceLogger.Error(err.Error(), zap.Error(err))
	}
	return err, driversPaged
}

func (d driverService) FindByRadius(ctx context.Context, radius int) (error error, drivers []*model.Driver) {
	err, drivers := d.Repo.FindByRadius(ctx, radius)
	if err != nil {
		driverServiceLogger.Error(err.Error(), zap.Error(err))
	}
	return err, drivers
}

func (d driverService) Create(ctx context.Context, driver *model.Driver) error {
	err := d.Repo.Create(ctx, driver)
	if err != nil {
		driverServiceLogger.Error(err.Error(), zap.Error(err))
	}
	return nil
}

// NewdriverService Create a new User Service
func NewdriverService(repo model.DriverRepository) DriverService {
	return &driverService{
		Repo:      repo,
		Validator: validator.New(),
	}
}
