package data

import (
	"IOTProject/internal/app"
	"IOTProject/internal/app/data/router"
	"IOTProject/kernel"
	"context"
	"sync"
)

type (
	Data struct {
		Name string
		app.UnimplementedModule
	}
)

func (p *Data) Info() string {
	return p.Name
}

func (p *Data) PreInit(engine *kernel.Engine) error {
	//dao.InitMS(engine.SKLMySQL.DB)
	return nil
}

func (p *Data) Init(*kernel.Engine) error {
	return nil
}

func (p *Data) PostInit(*kernel.Engine) error {
	return nil
}

func (p *Data) Load(engine *kernel.Engine) error {
	// 加载flamego api
	router.AppDataInit(engine.GIN)
	return nil
}

func (p *Data) Start(engine *kernel.Engine) error {
	return nil
}

func (p *Data) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Data) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {
		return nil
	}
}
