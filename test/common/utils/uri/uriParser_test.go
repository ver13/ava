package uri_test

import (
	"testing"

	. "github.com/ValentinEncinasRojas/ava/pkg/common/utils/uri"
)

func TestURIParser_cleanHosts(t *testing.T) {
	samples := []string{
		"supu",
		"127.0.0.1",
		"https://supu.local/",
		"http://127.0.0.1",
		"supu_42.local:8080/",
		"http://127.0.0.1:8080",
	}

	expected := []string{
		"http://supu",
		"http://127.0.0.1",
		"https://supu.local",
		"http://127.0.0.1",
		"http://supu_42.local:8080",
		"http://127.0.0.1:8080",
	}

	result := CleanHosts(samples)
	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("want: %s, have: %s\n", expected[i], result[i])
		}
	}
}

func TestURIParser_cleanPath(t *testing.T) {
	samples := []string{
		"supu/{tupu}",
		"supu/{tupu}{supu}",
		"/supu/{tupu}",
		"/supu.local/",
		"supu_supu.txt",
		"supu_42.local?a=8080",
		"supu/supu/supu?a=1&b=2",
		"debug/supu/supu?a=1&b=2",
	}

	expected := []string{
		"/supu/{tupu}",
		"/supu/{tupu}{supu}",
		"/supu/{tupu}",
		"/supu.local/",
		"/supu_supu.txt",
		"/supu_42.local?a=8080",
		"/supu/supu/supu?a=1&b=2",
		"/debug/supu/supu?a=1&b=2",
	}

	SetRoutingPattern(BracketsRouterPatternBuilder)

	for i := range samples {
		if have := CleanPath(samples[i]); expected[i] != have {
			t.Errorf("want: %s, have: %s\n", expected[i], have)
		}
	}
}
