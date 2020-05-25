package status

import (
	"os"
	"sync"

	"github.com/anmotor/internal/pkg/status"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	//Status format from status pkg
	Status = status.Status

	GenStatus struct {
		Success    Status
		NotFound   Status `yaml:"not_found"`
		BadRequest Status `yaml:"bad_request"`
		Internal   Status
	}

	BikeStatus struct {
		BikeNotFound  Status `yaml:"not_found_bike"`
		BikeDuplicate Status `yaml:"duplicated_bike"`
	}

	UserStatus struct {
		UserNotFound Status `yaml:"not_found_user"`
	}

	AuthStatus struct {
		InvalidUserPassword Status `yaml:"invalid_user_password"`
	}

	statuses struct {
		Gen  GenStatus
		Bike BikeStatus
		User UserStatus
		Auth AuthStatus
	}
)

var (
	all  *statuses
	once sync.Once
)

// Init load statuses from the given config file.
// Init panics if cannot access or error while parsing the config file.
func Init(conf string) {
	once.Do(func() {
		f, err := os.Open(conf)
		if err != nil {
			logrus.Errorf("Fail to open status file, %v", err)
			panic(err)
		}
		all = &statuses{}
		if err := yaml.NewDecoder(f).Decode(all); err != nil {
			logrus.Errorf("Fail to parse status file data to statuses struct, %v", err)
			panic(err)
		}
	})
}

// all return all registered statuses.
// all will load statuses from configs/Status.yml if the statuses has not initalized yet.
func load() *statuses {
	conf := os.Getenv("STATUS_PATH")
	if conf == "" {
		conf = "configs/status.yml"
	}
	Init(conf)
	return all
}

func Gen() GenStatus {
	return load().Gen
}

func Bike() BikeStatus {
	return load().Bike
}

func User() UserStatus {
	return load().User
}

func Success() Status {
	return Gen().Success
}

func Auth() AuthStatus {
	return load().Auth
}
