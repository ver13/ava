package http

import (
	"time"

	"github.com/ver13/ava/pkg/common/config/model/http"
	errorAVA "github.com/ver13/ava/pkg/common/error"
	serializerAVA "github.com/ver13/ava/pkg/common/serializer"
)

const (
	DialerTimeoutDefault       = 30 * time.Second
	DialerKeepAliveDefault     = 30 * time.Second
	DialerFallbackDelayDefault = 30 * time.Second
)

type DialerConfig struct {
	// DialerTimeout is the maximum amount of time a dial will wait for a connect to complete. If Deadline is also set, it may fail earlier.
	// The default is no timeout.
	//
	// When using TCP and dialing a host name with multiple IP addresses, the timeout may be divided between them.
	//
	// With or without a timeout, the operating system may impose its own earlier timeout. For instance, TCP timeouts are often around 3 minutes.
	DialerTimeout uint64 `mapstructure:"dialer_timeout,omitempty"`
	// FallbackDelay specifies the length of time to wait before spawning a fallback connection, when DualStack is enabled.
	// If zero, a default delay of 300ms is used.
	DialerFallbackDelay uint64 `mapstructure:"dialer_fallback_delay,omitempty"`
	// KeepAlive specifies the keep-alive period for an active network connection.
	// If zero, keep-alives are not enabled. Network protocols that do not support keep-alives ignore this field.
	DialerKeepAlive uint64 `mapstructure:"dialer_keep_alive,omitempty"`
}

func NewDialerConfig(dialerTimeout uint64, dialerFallbackDelay uint64, dialerKeepAlive uint64) (*http.Dialer, *errorAVA.Error) {
	panic("Not implemented.")
}

func NewDialerConfigDefault() (*http.Dialer, *errorAVA.Error) {
	panic("Not implemented.")
}

func (a *DialerConfig) ReadLocal(fileName string) (*http.Dialer, *errorAVA.Error) {
	panic("Not implemented.")
}

func (a *DialerConfig) Parser() (*http.Dialer, *errorAVA.Error) {
	var dialerTimeout time.Duration
	var dialerFallbackDelay time.Duration
	var dialerKeepAlive time.Duration

	if a.DialerTimeout == 0 {
		dialerTimeout = DialerTimeoutDefault
	} else {
		dialerTimeout = time.Duration(a.DialerTimeout * uint64(time.Second))
	}

	if a.DialerFallbackDelay == 0 {
		dialerFallbackDelay = DialerFallbackDelayDefault
	} else {
		dialerFallbackDelay = time.Duration(a.DialerFallbackDelay * uint64(time.Second))
	}

	if a.DialerKeepAlive == 0 {
		dialerKeepAlive = DialerKeepAliveDefault
	} else {
		dialerKeepAlive = time.Duration(a.DialerKeepAlive * uint64(time.Second))
	}

	return http.NewDialer(dialerTimeout, dialerFallbackDelay, dialerKeepAlive)
}

func (a *DialerConfig) Serializer(t serializerAVA.SerializerType) ([]byte, *errorAVA.Error) {
	serializer, errSerializer := serializerAVA.GetInstance().SerializerFactory(t)
	if errSerializer != nil {
		return nil, errSerializer
	}

	return serializer.Serializer(a)
}
