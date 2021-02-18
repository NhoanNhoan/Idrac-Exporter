package main

import (
	"fmt"
//	"github.com/intel-go/fastjson"
//	io_prometheus_client "github.com/prometheus/client_model/go"
//	"io/ioutil"
	"net/http"

	"github.com/alochym01/idrac-exporter/config"
	"github.com/alochym01/idrac-exporter/system"
	"github.com/alochym01/idrac-exporter/chassis"

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
	}
	fmt.Println(r.URL.Query().Get("idrac_host"))
	config.GOFISH, _ = gofish.Connect(conf)

	/*
	//

	fmt.Println("Running Test")
	processor := testUnmarshall(config.GOFISH)
	
	metric := prometheus.MustNewConstMetric(config.S_processor,
		prometheus.GaugeValue,
		0.0,
		processor.Actions,
			processor.Description,
			processor.Manufacturer,
			fmt.Sprintf("%v", processor.MaxSpeedMHz),
			fmt.Sprintf("%v", processor.MaxTDPWatts),
			processor.Model,
			fmt.Sprintf("%v", processor.ProcessorType),
			processor.Socket,
			processor.SubProcessors,
			fmt.Sprintf("%v", processor.TDPWatts),
			fmt.Sprintf("%v", processor.TotalCores),
			fmt.Sprintf("%v", processor.TotalEnabledCores),
			fmt.Sprintf("%v", processor.TotalThreads),
			processor.UUID,
	)
	//


	if err != nil {
		panic(err)
	}
	defer config.GOFISH.Logout()

	client := io_prometheus_client.Metric{}
	metric.Write(&client)
	fmt.Println(client.String())

	fmt.Println(metric.Desc())

		input := model.Metric(model.LabelSet{
			"a": "bb",
			"b": "c",
		})

		fmt.Println(input.String()
*/

	fmt.Println(" Connect successfull")

	mhandler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
	})
	mhandler.ServeHTTP(w, r)

}

func main() {
	fmt.Println("Running...")

	system := system.SystemCollector{}
	prometheus.Register(system)

	chassis := chassis.Chassis{}
	prometheus.Register(chassis)

	// Starting server
	http.HandleFunc("/metrics", metrichandler)
	http.ListenAndServe(":9000", nil)
}

/*
func testUnmarshall(client *gofish.APIClient) redfish.Processor {
	res, resErr := client.Get("/redfish/v1/Systems/System.Embedded.1/Processors/CPU.Socket.1")
	processor := redfish.Processor{}

	if nil != resErr {
		panic(resErr)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		contentBytes, readErr := ioutil.ReadAll(res.Body)

		if nil != readErr {
			panic(readErr)
		}

		contentString := string(contentBytes)
		fmt.Println(contentString)

		unmarshallErr := fastjson.Unmarshal(contentBytes, &processor)

		if nil != unmarshallErr {
			panic(unmarshallErr)
		}
	}

	return processor
}
 */
