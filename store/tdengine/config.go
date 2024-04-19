package tdengine

import "errors"

var (
	ErrEmptyHost       = errors.New("empty mysql host")
	ErrEmptyPort       = errors.New("empty mysql port")
	ErrEmptyUser       = errors.New("empty mysql user")
	ErrEmptyPass       = errors.New("empty mysql pass")
	ErrEmptyDB         = errors.New("empty mysql database")
	ErrEmptyDriverName = errors.New("empty mysql driver name")
)

type (
	// A OrmConf is a mysql config.
	OrmConf struct {
		Host       string `yaml:"Host"`
		Port       string `yaml:"Port"`
		User       string `yaml:"User"`
		Pass       string `yaml:"Pass"`
		Database   string `yaml:"Database"`
		DriverName string `yaml:"DriverName"`
	}
)

// Validate validates the MysqlConf.
func (rc OrmConf) Validate() error {
	if len(rc.Host) == 0 {
		return ErrEmptyHost
	}
	if len(rc.Port) == 0 {
		return ErrEmptyPort
	}
	if len(rc.User) == 0 {
		return ErrEmptyUser
	}
	if len(rc.Pass) == 0 {
		return ErrEmptyPass
	}
	if len(rc.Database) == 0 {
		return ErrEmptyDB
	}
	if len(rc.DriverName) == 0 {
		return ErrEmptyDriverName
	}

	return nil
}
