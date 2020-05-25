package bike

import (
	"context"
	"fmt"

	"github.com/anmotor/internal/app/status"
	"github.com/anmotor/internal/app/types"
	db "github.com/anmotor/internal/pkg/database"
	"github.com/anmotor/internal/pkg/uuid"
	"github.com/anmotor/internal/pkg/validator"

	"github.com/sirupsen/logrus"
)

type (
	mongoRepository interface {
		Create(ctx context.Context, user *types.Bike) error

		FindByNumber(ctx context.Context, number string) (*types.Bike, error)
	}

	Service struct {
		mongo mongoRepository
	}
)

func New(mongo mongoRepository) *Service {
	return &Service{
		mongo: mongo,
	}
}

// Manage Bikes

func (s *Service) Create(ctx context.Context, req *types.CreateBike) (*types.Bike, error) {

	if err := validator.Validate(req); err != nil {
		logrus.Errorf("Fail to create bike due to invalid req, %w", err)
		return nil, fmt.Errorf(err.Error()+" err: %w", status.Gen().BadRequest)
	}

	//Check duplicate project's name of this PM
	existingProject, err := s.mongo.FindByNumber(ctx, req.Number)

	if err != nil && !db.IsErrNotFound(err) {
		logrus.Errorf("failed to check existing bike by name: %v", err)
		return nil, fmt.Errorf("failed to check existing bike by name: %w", err)
	}

	if existingProject != nil {
		logrus.Errorf("Bike already exsit")
		return nil, status.Bike().BikeDuplicate
	}

	//Create new project
	bike := &types.Bike{
		BikeID: uuid.New(),
		Color:  req.Color,
		Cost:   req.Cost,
		Name:   req.Name,
		Number: req.Number,
		Rate:   0,
		Status: types.Avaiable,
	}

	if err = s.mongo.Create(ctx, bike); err != nil {
		logrus.Errorf("failed to create bike: %v", err)
		return nil, fmt.Errorf("failed to create bike, %w", err)
	}

	return bike, nil
}
