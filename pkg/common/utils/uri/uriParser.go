package uri

import (
	"net/url"
	"regexp"
	"strings"

	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
	loggerGmf "github.com/ValentinEncinasRojas/ava/pkg/common/logger"
	errorURIUtilsGmf "github.com/ValentinEncinasRojas/ava/pkg/common/utils/uri/error"
)

const (
	// BracketsRouterPatternBuilder uses brackets as route params delimiter
	BracketsRouterPatternBuilder = iota
	// ColonRouterPatternBuilder use a colon as route param delimiter
	ColonRouterPatternBuilder
	// DefaultMaxIdleConnsPerHost is the default value for the maxIdleConnsPerHost param
	DefaultMaxIdleConnsPerHost = 250

	// ConfigVersion is the current version of the error struct
	ConfigVersion = 2
)

var (
	// RoutingPattern to use during route conversion. By default, use the colon router pattern
	RoutingPattern         = ColonRouterPatternBuilder
	EndpointURLKeysPattern = regexp.MustCompile(`/{([a-zA-Z\-_0-9]+)\}`)
	hostPattern            = regexp.MustCompile(`(https?://)?([a-zA-Z0-9._\-]+)(:[0-9]{2,6})?/?`)
)

// URI implements the URIParser interface
type uri struct {
	int
}

func (u *uri) SetRoutingPattern(pattern int) {
	u.int = pattern
}
func SetRoutingPattern(pattern int) {
	GetInstance().SetRoutingPattern(pattern)
}

// CleanHosts applies the CleanHost method to every member of the received array of hosts
func (u *uri) CleanHosts(hosts []string) []string {
	cleaned := []string{}
	for i := range hosts {
		host, err := u.CleanHost(hosts[i])
		if err != nil {
			loggerGmf.Errorf("Host invalid %s. Ignored.", hosts[i])
		} else {
			cleaned = append(cleaned, host)
		}
	}
	return cleaned
}
func CleanHosts(hosts []string) []string {
	return GetInstance().CleanHosts(hosts)
}

// CleanHost sanitizes the received host
func (u *uri) CleanHost(host string) (string, errorGmf.ErrorGmfI) {
	matches := hostPattern.FindAllStringSubmatch(host, -1)
	if len(matches) != 1 {
		return "", errorURIUtilsGmf.InvalidHost(nil, "")
	}
	keys := matches[0][1:]
	if keys[0] == "" {
		keys[0] = "http://"
	}
	return strings.Join(keys, ""), nil
}
func CleanHost(host string) (string, errorGmf.ErrorGmfI) {
	return GetInstance().CleanHost(host)
}

// CleanPath trims all the extra slashes from the received URI path
func (u *uri) CleanPath(path string) string {
	return "/" + strings.TrimPrefix(path, "/")
}
func CleanPath(path string) string {
	return GetInstance().CleanPath(path)
}

// GetEndpointPath applies the proper replacement in the received path to generate valid route patterns
func (u *uri) GetEndpointPath(path string, params []string) string {
	result := path
	if u.int == ColonRouterPatternBuilder {
		for p := range params {
			parts := strings.Split(result, "?")
			parts[0] = strings.Replace(parts[0], "{"+params[p]+"}", ":"+params[p], -1)
			result = strings.Join(parts, "?")
		}
	}
	return result
}
func GetEndpointPath(path string, params []string) string {
	return GetInstance().GetEndpointPath(path, params)
}

// Parse parses rawurl into a URL structure.
func (u *uri) Parse(pattern string) (*url.URL, errorGmf.ErrorGmfI) {
	url, err := url.Parse(pattern)
	if err != nil {
		return nil, errorURIUtilsGmf.URLParseWrong(err, pattern)
	}
	return url, nil
}
func Parse(pattern string) (*url.URL, errorGmf.ErrorGmfI) {
	return GetInstance().Parse(pattern)
}
