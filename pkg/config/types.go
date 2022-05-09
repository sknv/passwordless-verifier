package config

import (
	"time"
)

// Duration represents YAML and TOML decodeable time.Duration type.
type Duration time.Duration

// Duration returns standard library's time.Duration value.
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

// UnmarshalText implements encoding.TextUnmarshaler to decode time.Duration values from TOML.
func (d *Duration) UnmarshalText(text []byte) error {
	dur, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}

	*d = Duration(dur)
	return nil
}
