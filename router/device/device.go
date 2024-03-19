package device

import (
	"IOTProject/kernel"
	"IOTProject/router"
	"golang.org/x/net/context"
	"sync"
)

type (
	Project struct {
		Name string
		router.UnimplementedModule
	}
)

func (p *Project) Info() string {
	return p.Name
}

func (p *Project) PreInit(engine *kernel.Engine) error {
	return dao.InitPG(engine.MainPG.DB)
}

func (p *Project) Init(*kernel.Engine) error {
	return nil
}

func (p *Project) PostInit(*kernel.Engine) error {
	return nil
}

func (p *Project) Load(engine *kernel.Engine) error {
	// 加载flamego api
	AppProjectInit(engine.GIN)
	return nil
}

func (p *Project) Start(engine *kernel.Engine) error {
	return nil
}

func (p *Project) Stop(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (p *Project) OnConfigChange() func(*kernel.Engine) error {
	return func(engine *kernel.Engine) error {

		return nil
	}
}
