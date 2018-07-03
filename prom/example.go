package prom

import (
	"github.com/prometheus/client_golang/prometheus"

	"log"
	"prom/dnspod/api"
	"prom/dnspod/g"
	"sync"
)

type exampleCollector struct {
	desc *prometheus.Desc
}

// NewTimeCollector returns a new Collector exposing the current system time in
// seconds since epoch.
func NewExampleCollector() (Collector, error) {
	return &exampleCollector{
		desc: prometheus.NewDesc(
			Namespace+"_record_list",
			"https://dnsapi.cn/Record.List",
			[]string{"name", "type", "line", "value", "ttl", "id"}, nil,
		//	[]string{"name", "type", "line", "value", "ttl"}, nil,
		),
	}, nil
}

func (c *exampleCollector) Update(ch chan<- prometheus.Metric) error {
	names := g.Config().RecordNames
	log.Println("Update----->", names)
	if len(names) < 1 {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(len(names))
	for _, name := range names {

		go func(real_name string) {
			records, err := api.GetExample(real_name)

			log.Println(real_name, "->", len(records), err)
			if err == nil && records != nil {
				for _, record := range records {
					ch <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue,
						float64(record.Weight),
						record.Name, record.Type, record.Line, record.Value, record.TTL, record.ID)
					//record.Name, record.Type, record.Line, record.Value, record.TTL)
				}
			}
			wg.Done()
		}(name)

	}
	wg.Wait()

	//ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, 123, "aa", "xx")
	return nil
}
