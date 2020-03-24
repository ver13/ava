package url

import (
	"net/url"
	"regexp"
	"strings"

	errorAVA "github.com/ver13/ava/pkg/common/error"
	loggerAVA "github.com/ver13/ava/pkg/common/logger"
	errorURIUtilsAVA "github.com/ver13/ava/pkg/common/utils/url/error"
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

// URL implements the URL interface
type URL int

func (u *URL) SetRoutingPattern(pattern int) {
	*u = URL(pattern)
}
func SetRoutingPattern(pattern int) {
	GetInstance().SetRoutingPattern(pattern)
}

// CleanHosts applies the CleanHost method to every member of the received array of hosts
func (u *URL) CleanHosts(hosts []string) []string {
	cleaned := []string{}
	for i := range hosts {
		host, err := u.CleanHost(hosts[i])
		if err != nil {
			loggerAVA.Errorf("Host invalid %s. Ignored.", hosts[i])
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
func (u *URL) CleanHost(host string) (string, *errorAVA.Error) {
	matches := hostPattern.FindAllStringSubmatch(host, -1)
	if len(matches) != 1 {
		return "", errorURIUtilsAVA.InvalidHost(nil, "")
	}
	keys := matches[0][1:]
	if keys[0] == "" {
		keys[0] = "http://"
	}
	return strings.Join(keys, ""), nil
}
func CleanHost(host string) (string, *errorAVA.Error) {
	return GetInstance().CleanHost(host)
}

// CleanPath trims all the extra slashes from the received URL path
func (u *URL) CleanPath(path string) string {
	return "/" + strings.TrimPrefix(path, "/")
}
func CleanPath(path string) string {
	return GetInstance().CleanPath(path)
}

// GetEndpointPath applies the proper replacement in the received path to generate valid route patterns
func (u *URL) GetEndpointPath(path string, params []string) string {
	result := path
	if *u == ColonRouterPatternBuilder {
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
func (u *URL) Parse(pattern string) (*url.URL, *errorAVA.Error) {
	url, err := url.Parse(pattern)
	if err != nil {
		return nil, errorURIUtilsAVA.URLParseWrong(err, pattern)
	}
	return url, nil
}
func Parse(pattern string) (*url.URL, *errorAVA.Error) {
	return GetInstance().Parse(pattern)
}
