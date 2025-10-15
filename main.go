package main

import (
	"MobileForensicsToolkit/analyzers/database"
	"MobileForensicsToolkit/analyzers/extractor"
	"MobileForensicsToolkit/analyzers/passwdCalc"
	"MobileForensicsToolkit/analyzers/winreg"
	"MobileForensicsToolkit/tools/cracker"
	"MobileForensicsToolkit/tools/timestamp"
	"context"
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	decDb := database.NewDecryptDatabase()
	passwdCal := passwdCalc.NewPasswordCalculator()
	infoExtract := extractor.NewInfoExtractor()
	reg := winreg.NewRegistryAnalyzer()
	crack := cracker.NewForensicsCracker()
	tp := timestamp.NewTimeStampParser()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Mobile Forensics Toolkit",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		DisableResize:    true,
		DragAndDrop:      &options.DragAndDrop{DisableWebViewDrop: false, EnableFileDrop: true},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Windows:          &windows.Options{IsZoomControlEnabled: false},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			decDb.InitCtx(ctx)
			passwdCal.InitCtx(ctx)
			infoExtract.InitCtx(ctx)
			reg.InitCtx(ctx)
			crack.InitCtx(ctx)
			tp.InitCtx(ctx)
		},
		Bind: []interface{}{
			app,
			decDb,
			passwdCal,
			infoExtract,
			reg,
			crack,
			tp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
