package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/yiwenlong/ServiceUIDemo-desktop/controller"
	"os"
)

type LaunchdUIApp struct {
	app        *widgets.QApplication
	mainWindow *MainWindow
	systemTray *SystemTray
	serverCtl  *controller.ServerController
}

func NewApp() *LaunchdUIApp {
	app := LaunchdUIApp{
		app: widgets.NewQApplication(len(os.Args), os.Args),
	}
	app.mainWindow = NewMainWindow(&app)
	app.systemTray = NewSystemTray(&app)
	app.serverCtl = controller.NewServerController(app.HomePath())
	return &app
}

func (lapp *LaunchdUIApp) HomePath() string {
	dir := core.NewQDir2(lapp.app.ApplicationDirPath())
	dir.Cd("../../")
	return dir.AbsolutePath()
}

func (lapp *LaunchdUIApp) Launch() {
	lapp.mainWindow.Launch()
	lapp.app.Exec()
}

func (lapp *LaunchdUIApp) Exit() {
	lapp.app.Exit(0)
}

func (lapp *LaunchdUIApp) ShowMainWindow() {
	if lapp.mainWindow == nil {
		return
	}
	lapp.mainWindow.Show()
}

func (lapp *LaunchdUIApp) LaunchSystemTray() {
	lapp.systemTray.Launch()
}

func (lapp *LaunchdUIApp) CloseSystemTray() {
	lapp.systemTray.Close()
}
