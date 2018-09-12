package tracking

import (
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	TimeFmt = "2006-01-02"
)

func mustParseFloat32(s string) float32 {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Fatalf("error parsing float %q: %v", s, err)
	}
	return float32(v)
}

func mustParseTime(s string) time.Time {
	t, err := time.Parse(TimeFmt, s)
	if err != nil {
		log.Fatalf("error parsing time %q: %v", s, err)
	}
	return t
}

func removeTrailingSpaces(s string) string {
	parts := strings.Split(s, "\n")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return strings.Join(parts, "\n")
}
