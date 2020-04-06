package metrics

import (
	"time"

	metricsServiceModelConfigAVA "github.com/ver13/ava/pkg/common/config/model/services/metrics"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
	urlUtilsAVA "github.com/ver13/ava/pkg/common/utils/url"
)

type Metrics struct {
	URL                  string        `mapstructure:"url"`
	Timeout              time.Duration `mapstructure:"timeout"`
	KeepAlive            time.Duration `mapstructure:"keepAlive"`
	TLSHandshakeTimeout  time.Duration `mapstructure:"tlsHandshakeTimeout"`
	CounterIsAvailable   bool          `mapstructure:"counterIsAvailable"`
	GaugeIsAvailable     bool          `mapstructure:"gaugeIsAvailable"`
	HistogramIsAvailable bool          `mapstructure:"histogramIsAvailable"`
	SummaryIsAvailable   bool          `mapstructure:"summaryIsAvailable"`
}

func (b *Metrics) Parser() (*metricsServiceModelConfigAVA.Metrics, *errorAVA.Error) {
	url, err := urlUtilsAVA.Parse(b.URL)
	if err != nil {
		return nil, err
	}
	return &metricsServiceModelConfigAVA.Metrics{
		URL:                  url,
		Name:                 "AVA Metrics Service",
		CounterIsAvailable:   b.CounterIsAvailable,
		GaugeIsAvailable:     b.GaugeIsAvailable,
		HistogramIsAvailable: b.HistogramIsAvailable,
		SummaryIsAvailable:   b.SummaryIsAvailable,
	}, nil
}

func (b *Metrics) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(b)
}
