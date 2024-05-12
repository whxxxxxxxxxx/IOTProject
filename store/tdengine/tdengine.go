package tdengine

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

type (
	Orm struct {
		Host       string
		Port       string
		User       string
		Pass       string
		Database   string
		DriverName string

		*sql.DB
	}
	Option func(r *Orm)
)

var tdInit = sync.Once{}

func (r *Orm) GetOrm() *sql.DB {
	return r.DB
}

func (r *Orm) OrmConnectionUpdate(conf OrmConf) *Orm {
	orm, err := NewTDEngineOrm(conf)
	if err != nil {
		return r
	}
	return orm
}

func MustNewTDEngineOrm(conf OrmConf, opts ...Option) *Orm {
	orm, err := NewTDEngineOrm(conf, opts...)
	if err != nil {
		os.Exit(1)
	}
	return orm
}

func NewTDEngineOrm(conf OrmConf, opts ...Option) (*Orm, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	opts = append([]Option{WithAddr(conf.Host, conf.Port)}, opts...)
	opts = append([]Option{WithAuth(conf.User, conf.Pass)}, opts...)
	opts = append([]Option{WithDBName(conf.Database)}, opts...)
	opts = append([]Option{WithDriverName(conf.DriverName)}, opts...)
	return newOrm(opts...)
}

func WithAddr(host, port string) Option {
	return func(r *Orm) {
		r.Host = host
		r.Port = port
	}
}

func WithAuth(user, pass string) Option {
	return func(r *Orm) {
		r.Pass = pass
		r.User = user
	}
}

func WithDBName(db string) Option {
	return func(r *Orm) {
		r.Database = db
	}
}

func WithDriverName(driverName string) Option {
	return func(r *Orm) {
		r.DriverName = driverName
	}
}

func newOrm(opts ...Option) (*Orm, error) {
	m := &Orm{}
	for _, opt := range opts {
		opt(m)
	}
	var err error
	tdInit.Do(func() {
		dsn := fmt.Sprintf("%s:%s@http(%s:%s)/%s",
			m.User, m.Pass, m.Host, m.Port, m.Database)
		db, err := sql.Open(m.DriverName, dsn)
		_, err = db.Exec("create database if not exists " + m.Database + " duration 8h keep 1")
		if err != nil {
			return
		}
		m.DB = db
	})

	return m, err
}
