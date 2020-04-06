package metricsService

import (
	"net/url"
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

type Metrics struct {
	URL                  *url.URL
	Name                 string
	Timeout              time.Duration
	KeepAlive            time.Duration
	TLSHandshakeTimeout  time.Duration
	CounterIsAvailable   bool
	GaugeIsAvailable     bool
	HistogramIsAvailable bool
	SummaryIsAvailable   bool
}

func NewMetrics(URL *url.URL, name string, timeout time.Duration, keepAlive time.Duration, TLSHandshakeTimeout time.Duration, counterIsAvailable bool, gaugeIsAvailable bool, histogramIsAvailable bool, summaryIsAvailable bool) (*Metrics, *errorAVA.Error) {
	return &Metrics{URL: URL, Name: name, Timeout: timeout, KeepAlive: keepAlive, TLSHandshakeTimeout: TLSHandshakeTimeout, CounterIsAvailable: counterIsAvailable, GaugeIsAvailable: gaugeIsAvailable, HistogramIsAvailable: histogramIsAvailable, SummaryIsAvailable: summaryIsAvailable}, nil
}

func (b *Metrics) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(b)
}
