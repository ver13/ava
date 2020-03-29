package metricsService

import (
	"time"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	uriUtilsAVA "github.com/ver13/ava/pkg/common/utils/url"
)

type MetricsConfigAVA struct {
	URL                  string        `mapstructure:"url"`
	Timeout              time.Duration `mapstructure:"timeout"`
	KeepAlive            time.Duration `mapstructure:"keepAlive"`
	TLSHandshakeTimeout  time.Duration `mapstructure:"tlsHandshakeTimeout"`
	CounterIsAvailable   bool          `mapstructure:"counterIsAvailable"`
	GaugeIsAvailable     bool          `mapstructure:"gaugeIsAvailable"`
	HistogramIsAvailable bool          `mapstructure:"histogramIsAvailable"`
	SummaryIsAvailable   bool          `mapstructure:"summaryIsAvailable"`
}

func (b *MetricsConfigAVA) Parser() (*MetricsService, *errorAVA.Error) {
	url, err := uriUtilsAVA.Parse(b.URL)
	if err != nil {
		return nil, err
	}
	return &MetricsService{
		UrlPattern:           url,
		ServiceName:          "AVA Metrics Service",
		CounterIsAvailable:   b.CounterIsAvailable,
		GaugeIsAvailable:     b.GaugeIsAvailable,
		HistogramIsAvailable: b.HistogramIsAvailable,
		SummaryIsAvailable:   b.SummaryIsAvailable,
	}, nil
}

func (b *MetricsConfigAVA) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(b)
}
