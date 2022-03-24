package main

import (
	"os"

	conductor "github.com/netflix/conductor/client/go"
	"github.com/netflix/conductor/client/go/metrics"
	"github.com/netflix/conductor/client/go/settings"
	"github.com/netflix/conductor/client/go/task/sample"
	log "github.com/sirupsen/logrus"
)

//Example init function that shows how to configure logging
//Using json formatter and changing level to Debug
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	//Stdout, change to a file for production use case
	log.SetOutput(os.Stdout)

	// Set to debug for demonstration.  Change to Info for production use cases.
	log.SetLevel(log.DebugLevel)
}

func main() {
	go metrics.ProvideMetrics(
		"/metrics",
		2112,
	)

	var metricsCollector = metrics.NewMetricsCollector()

	var authenticationSettings = settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)

	var httpSettings = settings.NewHttpSettingsWithBaseUrlAndDebug(
		"https://play.orkes.io/api",
		true,
	)

	c := conductor.NewConductorWorker(
		authenticationSettings,
		httpSettings,
		metricsCollector,
		1,
		5000,
	)

	c.Start("go_task_example", "", sample.Task_1_Execution_Function, true)
}
