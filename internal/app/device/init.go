package device

import (
	"IOTProject/internal/app"
	"IOTProject/internal/app/device/dao"
	"IOTProject/internal/app/device/router"
	"IOTProject/kernel"
	"context"

	"sync"
)

type (
	Device struct {
		Name string
		app.UnimplementedModule
	}
)

func (p *Device) Info() string {
	return p.Name
}

func (p *Device) PreInit(engine *kernel.Engine) error {
	err := dao.InitMS(engine.SKLMySQL.DB)
	return err
}

func (p *Device) Init(*kernel.Engine) error {
	return nil
}

func (p *Device) PostInit(*kernel.Engine) error {
	return nil
}

func (p *Device) Load(engine *kernel.Engine) error {
	// 加载flamego api
	router.AppDeviceInit(engine.GIN)
	return nil
}

func (p *Device) Start(engine *kernel.Engine) error {
	return nil
}

func (p *Device) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Device) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {

		return nil
	}
}
