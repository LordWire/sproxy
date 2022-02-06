package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pseidemann/finish"
	"github.com/spf13/viper"
)

var backendList ServiceBackends
var configuration Configuration

func init() {
	// viper configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic(err)
	}

	// there's no other service, so it's just initialized with a hardcoded Services[0]
	// The model is correctly setup and extensible for future extensions though.
	backendList.loadBackends(configuration.Proxy.Services[0])
	registerMetrics()
}

func main() {

	// setup the http server and attach the route handlers
	listenAddress := configuration.Proxy.Listen.Address +
		":" + strconv.Itoa(configuration.Proxy.Listen.Port)
	http.HandleFunc("/", instrumentedReverseProxy("randomproxy", serveReverseProxy))
	http.HandleFunc("/admin/health", health) // a dummy health endpoint
	http.Handle("/admin/metrics", promhttp.Handler())
	srv := &http.Server{Addr: listenAddress}
	fin := finish.New()
	fin.Add(srv)

	// start server
	log.Println("Proxy listening on: " + listenAddress)
	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}

	}()
	fin.Wait()

}
