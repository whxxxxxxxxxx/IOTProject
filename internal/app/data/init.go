package data

import (
	"IOTProject/internal/app"
	"IOTProject/internal/app/data/dao"
	"IOTProject/internal/app/data/router"
	"IOTProject/internal/app/data/service"
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
	err := dao.InitTD(engine.TDEngine.DB)
	return err
}

func (p *Data) Init(engine *kernel.Engine) error {

	service.CreateDataFromJs()

	err := service.SaveDataToDB()
	return err
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
