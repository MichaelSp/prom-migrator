// This file and its contents are licensed under the Apache License 2.0.
// Please see the included NOTICE for copyright information and
// LICENSE for a copy of the license.

package utils

import (
	"fmt"
	"strings"
)

type HeadersFlag struct {
	Headers map[string]string
}

func (f HeadersFlag) String() string {
	var s string
	for k, vs := range f.Headers {
		s += k + "=" + vs
	}
	return s
}

func (f HeadersFlag) Set(value string) error {
	k, v, found := strings.Cut(value, ":")
	if !found {
		return fmt.Errorf("HTTP header values should be specified with the format `key:value`")
	}
	f.Headers[k] = v
	return nil
}
