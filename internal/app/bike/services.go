package bike

import (
	"context"
	"fmt"

	"github.com/anmotor/internal/app/status"
	"github.com/anmotor/internal/app/types"
	db "github.com/anmotor/internal/pkg/database"
	"github.com/anmotor/internal/pkg/uuid"
	"github.com/anmotor/internal/pkg/validator"
	"github.com/imdario/mergo"

	"github.com/sirupsen/logrus"
)

type (
	mongoRepository interface {
		Create(ctx context.Context, user *types.Bike) error
		Update(ctx context.Context, id string, req *types.UpdateBike) error
		Delete(ctx context.Context, id string) error

		FindByNumber(ctx context.Context, number string) (*types.Bike, error)
		FindByBikeID(ctx context.Context, bikeID string) (*types.Bike, error)

		FindAll(ctx context.Context) ([]*types.Bike, error)
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

	bikeExisted, err := s.mongo.FindByNumber(ctx, req.Number)

	if err != nil && !db.IsErrNotFound(err) {
		logrus.Errorf("failed to check existing bike by name: %v", err)
		return nil, fmt.Errorf("failed to check existing bike by name: %w", err)
	}

	if bikeExisted != nil {
		logrus.Errorf("Bike already exsit")
		return nil, status.Bike().BikeDuplicate
	}

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

func (s *Service) Update(ctx context.Context, id string, req types.UpdateBike) (*types.UpdateBike, error) {

	bikeFromDB, err := s.mongo.FindByBikeID(ctx, id)

	if err != nil && !db.IsErrNotFound(err) {
		logrus.Errorf("failed to check existing Bike by id: %v", err)
		return nil, fmt.Errorf("failed to check existing bike by id: %w", err)
	}

	if db.IsErrNotFound(err) {
		logrus.Errorf("Bike doesn't exist")
		return nil, status.Bike().BikeNotFound
	}

	bike := types.UpdateBike{
		Color:  bikeFromDB.Color,
		Cost:   bikeFromDB.Cost,
		Name:   bikeFromDB.Name,
		Number: bikeFromDB.Number,
		Status: bikeFromDB.Status,
	}
	fmt.Println("before", req)
	fmt.Println("fromdb", bike)
	//update only change field
	if err := mergo.Merge(&req, bike); err != nil {
		return nil, err
	}
	fmt.Println("after", req)
	fmt.Println("after fromdb", bike)

	if err = s.mongo.Update(ctx, id, &req); err != nil {
		logrus.Errorf("failed to update bike: %v", err)
		return nil, fmt.Errorf("failed to update bike, %w", err)
	}

	return &req, nil
}

func (s *Service) FindAll(ctx context.Context) ([]*types.Bike, error) {

	bikes, err := s.mongo.FindAll(ctx)

	info := make([]*types.Bike, 0)

	for _, bike := range bikes {
		info = append(info, &types.Bike{
			BikeID:    bike.BikeID,
			Name:      bike.Name,
			Color:     bike.Color,
			Cost:      bike.Cost,
			CreatedAt: bike.CreatedAt,
			Number:    bike.Number,
			Rate:      bike.Rate,
			Status:    bike.Status,
			UpdateAt:  bike.UpdateAt,
		})
	}

	return info, err
}

func (s *Service) FindByID(ctx context.Context, id string) (*types.Bike, error) {

	bike, err := s.mongo.FindByBikeID(ctx, id)

	return bike, err
}

func (s *Service) Delete(ctx context.Context, id string) error {

	if err := s.mongo.Delete(ctx, id); err != nil {
		logrus.Errorf("Fail to delete bike due to %v", err)
		return status.Bike().BikeNotFound
	}
	return nil
}
