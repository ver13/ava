package timezones

type Timezones struct {
	supportedZones []string
}

type TimezonesI interface {
	GetSupported() []string
}

func New() TimezonesI {
	timezones := Timezones{}

	timezones.supportedZones = DefaultSupportedTimezones

	return &timezones
}

func (t *Timezones) GetSupported() []string {
	return t.supportedZones
}

func DefaultUserTimezone() map[string]string {
	defaultTimezone := make(map[string]string)
	defaultTimezone["useAutomaticTimezone"] = "true"
	defaultTimezone["automaticTimezone"] = ""
	defaultTimezone["manualTimezone"] = ""

	return defaultTimezone
}
