package prom

import (

	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"

)

func DelegateHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "'target' parameter must be specified", 400)
		return
	}
	moduleName := r.URL.Query().Get("module")
	if moduleName == "" {
		moduleName = "default"
	}


	registry := prometheus.NewRegistry()
	collector := DelegateCollector{Target: target, Module: moduleName}
	registry.MustRegister(collector)


	// Delegate http serving to Promethues client library, which will call collector.Collect.
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

type DelegateCollector struct {
	Target string
	Module string
}

func (c DelegateCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("dummy", "dummy", nil, nil)
}


var arr = []string{
	"a","b","c","d",
}

func (c DelegateCollector) Collect(ch chan<- prometheus.Metric) {
	r := rand.Intn(len(arr))


	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			"delegate",
			"Total SNMP time scrape took (walk and processing).",
			[]string{"target", "module", "var"},
			nil),
		prometheus.GaugeValue,
		float64(rand.Intn(10)),
		c.Target, c.Module, arr[r],
	)
}


