package web

import (
	"fmt"
	"github.com/loganetherton/pm-go/config"
	"github.com/loganetherton/pm-go/utils"
	"github.com/loganetherton/pm-go/web/routes"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Init() {
	tracer.Start(
		tracer.WithService("pm"),
		tracer.WithEnv(config.AppEnv),
	)
	defer tracer.Stop()
	httpTrace := httptrace.NewServeMux()
	fmt.Println(httpTrace)
	defer utils.Recover("Unhandled exception", "And another")

	routes.SetupRoutes()
}
