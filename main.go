package main

import (
	"embed"
	"encoding/json"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"log"
	"net/http"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

type Feather struct {
	Content  string `json:"content"`
	Origin   string `json:"origin"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (a *App) GetFeather() Feather {
	url := "https://v1.jinrishici.com/all.json"
	client, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer client.Body.Close()
	}
	if client.StatusCode == 200 {
		feather := Feather{}
		err = json.NewDecoder(client.Body).Decode(&feather)
		if err != nil {
			log.Fatal(err)
		} else {
			return feather
		}
	}
	return Feather{}
}

func main() {

	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:             "古诗词，好心情",
		Width:             480,
		Height:            200,
		MinWidth:          480,
		MinHeight:         200,
		MaxWidth:          480,
		MaxHeight:         200,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
		},
		// Mac平台特定选项
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 true,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameAccessibilityHighContrastVibrantLight,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "古诗词，好心情",
				Message: "一款随机的古诗文应用",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
