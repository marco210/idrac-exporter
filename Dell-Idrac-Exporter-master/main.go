package main

import (
	"fmt"
	"net/http"

	"github.com/alochym01/idrac-exporter/chassis"
	"github.com/alochym01/idrac-exporter/config"
	"github.com/alochym01/idrac-exporter/system"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stmcginnis/gofish"
)

func metrichandler(w http.ResponseWriter, r *http.Request) {
	// var err error
	conf := gofish.ClientConfig{
		Endpoint: r.URL.Query().Get("idrac_host"),
		Username: config.Idracuser,
		Password: config.Idracpassword,
		Insecure: true,
	} // struct connection đến Redfish Service: url + username + password + authen

	fmt.Println(r.URL.Query().Get("idrac_host"))

	var err error
	config.GOFISH, err = gofish.Connect(conf)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer config.GOFISH.Logout()

	fmt.Println(" Connect successful")

	mhandler := promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		})
	mhandler.ServeHTTP(w, r)
}

func main() {
	const PORT = "9000"
	fmt.Println("Server listening at ", PORT)

	// Listen all interfaces at port 9000
	const IP_ADDRESS = ":" + PORT

	system := system.SystemCollector{}
	prometheus.Register(system)

	chassis := chassis.Chassis{}
	prometheus.Register(chassis)

	// Starting server
	http.HandleFunc("/metrics", metrichandler)
	http.ListenAndServe(IP_ADDRESS, nil)
}
