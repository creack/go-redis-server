// +build !debug

package redis

// Debugf does nothing in non-debug mode.
func Debugf(format string, a ...interface{}) {
}
