package app

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	// nolint
	ctx   context.Context
	Bind  []any
	Menu  *menu.Menu
	isMax bool
}

func NewApp(
	serv *Service,
) *App {
	app := &App{Bind: []any{serv}}
	app.Menu = app.menuMain()

	return app
}

func (p *App) menuMain() *menu.Menu {
	var (
		main     = menu.NewMenu()
		helpMenu = main.AddSubmenu("Help")
	)

	p.menuWindow(main)

	helpMenu.AddText("About", nil, p.about)
	helpMenu.AddSeparator()
	helpMenu.AddText("Quit", keys.CmdOrCtrl("q"), p.quit)

	return main
}

func (p *App) menuWindow(main *menu.Menu) {
	win := main.AddSubmenu("Window")
	win.AddText("Maximise", nil, func(_ *menu.CallbackData) { runtime.WindowToggleMaximise(p.ctx) })
	win.AddText("Minimise", nil, func(_ *menu.CallbackData) { runtime.WindowMinimise(p.ctx) })
	win.AddSeparator()
	win.AddText("Fullscreen", keys.Key("f11"), p.fullscreen)
}

func (p *App) fullscreen(_ *menu.CallbackData) {
	if runtime.WindowIsFullscreen(p.ctx) {
		runtime.WindowUnfullscreen(p.ctx)

		if p.isMax {
			runtime.WindowMaximise(p.ctx)
		} else {
			runtime.WindowUnmaximise(p.ctx)
		}

		return
	}

	p.isMax = runtime.WindowIsMaximised(p.ctx)

	if !p.isMax {
		const _sleep = time.Millisecond * 100

		runtime.WindowMaximise(p.ctx)
		time.Sleep(_sleep)
	}

	runtime.WindowFullscreen(p.ctx)
}

func HideMenu(item *menu.Menu) {
	if item == nil {
		return
	}

	for _, men := range item.Items {
		men.Hide()
		HideMenu(men.SubMenu)
	}
}

func (p *App) about(_ *menu.CallbackData) {
	runtime.EventsEmit(p.ctx, "about", true)
}

func (p *App) quit(_ *menu.CallbackData) {
	runtime.Quit(p.ctx)
}

func (p *App) Startup(ctx context.Context) {
	p.ctx = ctx
}
