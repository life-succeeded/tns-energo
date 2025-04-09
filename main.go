package main

import (
	"os"
	"tns-energo/config"
	"tns-energo/lib/authorize"
	libconfig "tns-energo/lib/config"
	"tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	libos "tns-energo/lib/os"

	"github.com/shopspring/decimal"
)

func main() {
	configureDecimal()

	log := liblog.NewLogger("tns-energo")
	log.Debug("service up")

	var settings config.Settings
	if err := libconfig.Parse(&settings); err != nil {
		log.Errorf("failed to parse config: %v", err)
		return
	}

	// нет времени объяснять, суй секрет себе в кэш
	authorize.CachedSecret = []byte(settings.Auth.Secret)

	mainCtx, cancelMainCtx := ctx.Background().WithCancel()
	defer cancelMainCtx()

	app := NewApp(mainCtx, log, settings)

	if err := app.InitDatabases(os.DirFS("./"), "database/migrations/postgres"); err != nil {
		log.Errorf("failed to init databases: %v", err)
		return
	}

	if err := app.InitServices(); err != nil {
		log.Errorf("failed to init services: %v", err)
		return
	}

	app.InitServer()

	if err := app.Start(); err != nil {
		log.Errorf("failed to start app: %v", err)
		return
	}

	libos.WaitTerminate(mainCtx, app.Stop)

	log.Debug("service down")
}

func configureDecimal() {
	decimal.DivisionPrecision = 2
	decimal.MarshalJSONWithoutQuotes = true
}
