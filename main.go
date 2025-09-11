package main

import (
	"tns-energo/config"

	"github.com/shopspring/decimal"
	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
	"github.com/sunshineOfficial/golib/goos"
)

func main() {
	configureDecimal()

	log := golog.NewLogger("tns-energo")
	log.Debug("service up")

	settings, err := config.Parse()
	if err != nil {
		log.Errorf("failed to parse config: %v", err)
		return
	}

	mainCtx, cancelMainCtx := goctx.Background().WithCancel()
	defer cancelMainCtx()

	app := NewApp(mainCtx, log, settings)

	if err = app.InitDatabases(); err != nil {
		log.Errorf("failed to init databases: %v", err)
		return
	}

	if err = app.InitServices(); err != nil {
		log.Errorf("failed to init services: %v", err)
		return
	}

	app.InitServer()

	if err = app.Start(); err != nil {
		log.Errorf("failed to start app: %v", err)
		return
	}

	goos.WaitTerminate(mainCtx, app.Stop)

	log.Debug("service down")
}

func configureDecimal() {
	decimal.DivisionPrecision = 2
	decimal.MarshalJSONWithoutQuotes = true
}
