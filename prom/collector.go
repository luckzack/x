package prom

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

const Namespace = "namespace"
const (
	defaultCollectors = "example"
)

var Factories = make(map[string]func() (Collector, error))

func init() {
	Factories["example"] = NewExampleCollector
}

func warnDeprecated(collector string) {
	log.Println("The %s collector is deprecated and will be removed in the future!", collector)
}

type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Update(ch chan<- prometheus.Metric) error
}

type typedDesc struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

func (d *typedDesc) mustNewConstMetric(value float64, labels ...string) prometheus.Metric {
	return prometheus.MustNewConstMetric(d.desc, d.valueType, value, labels...)
}
