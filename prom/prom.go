package prom

import (
	"fmt"
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"time"
	"log"
)

var (
	scrapeDurationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "scrape", "collector_duration_seconds"),
		"node_exporter: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	scrapeSuccessDesc = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "scrape", "collector_success"),
		"node_exporter: Whether a collector succeeded.",
		[]string{"collector"},
		nil,
	)
)

type NodeCollector struct {
	collectors map[string]Collector
}

// Describe implements the prometheus.Collector interface.
func (n NodeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- scrapeDurationDesc
	ch <- scrapeSuccessDesc
}

// Collect implements the prometheus.Collector interface.
func (n NodeCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(n.collectors))
	for name, c := range n.collectors {
		go func(name string, c Collector) {
			execute(name, c, ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
}

func execute(name string, c Collector, ch chan<- prometheus.Metric) {

	begin := time.Now()
	err := c.Update(ch)
	duration := time.Since(begin)

	//var success float64

	if err != nil {
		log.Printf("ERROR: %s collector failed after %fs: %s", name, duration.Seconds(), err)
		//	success = 0
	} else {
		//log.Printf("OK: %s collector succeeded after %fs.", name, duration.Seconds())
		//	success = 1
	}
	//	ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, duration.Seconds(), name)
	//	ch <- prometheus.MustNewConstMetric(scrapeSuccessDesc, prometheus.GaugeValue, success, name)
}

func loadCollectors(list string) (map[string]Collector, error) {
	collectors := map[string]Collector{}
	for _, name := range strings.Split(list, ",") {
		fn, ok := Factories[name]
		if !ok {
			return nil, fmt.Errorf("collector '%s' not available", name)
		}
		c, err := fn()
		if err != nil {
			return nil, err
		}
		collectors[name] = c
	}
	return collectors, nil
}


func Init(){

	collectors, err := loadCollectors(defaultCollectors)
	if err != nil {
		log.Fatalf("Couldn't load collectors: %s", err)
	}

	log.Println("Enabled collectors:")
	for n := range collectors {
		log.Printf(" - %s", n)
	}

	if err := prometheus.Register(NodeCollector{collectors: collectors}); err != nil {
		log.Fatalf("Couldn't register collector: %s", err)
	}
}
