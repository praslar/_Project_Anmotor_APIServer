package api

import "github.com/anmotor/internal/app/user"

func newUserService() (*user.Service, error) {
	s, err := dialDefaultMongoDB()
	if err != nil {
		return nil, err
	}

	repo := user.NewMongoDBRepo(s)
	return user.NewService(repo), nil
}
