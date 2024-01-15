package util

import "fmt"

func TagFrom(tag string, value string) string {
	return fmt.Sprintf("<%s:%d>%s", tag, len(value), value)
}
