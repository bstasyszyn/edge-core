/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import (
	"github.com/trustbloc/edge-core/pkg/internal/logging/metadata"
)

//nolint:lll
const (
	// loggerNotInitializedMsg is used when a logger is not initialized before logging
	loggerNotInitializedMsg = "Default logger initialized (please call log.Initialize() if you wish to use a custom logger)"
	loggerModule            = "edge_core_pkg_log"
)

// New creates and returns a Logger implementation based on given module name.
// note: the underlying logger instance is lazy initialized on first use.
// To use your own logger implementation provide logger provider in 'Initialize()' before logging any line.
// If 'Initialize()' is not called before logging any line then default logging implementation will be used.
func New(module string) Logger {
	return loggerProvider().GetLogger(module)
}

// SetLevel - setting log level for given module
//  Parameters:
//  module is module name
//  level is logging level
//
// If not set default logging level is info
func SetLevel(module string, level Level) {
	metadata.SetLevel(module, metadata.Level(level))
}

// GetLevel - getting log level for given module
//  Parameters:
//  module is module name
//
//  Returns:
//  logging level
//
// If not set default logging level is info
func GetLevel(module string) Level {
	return Level(metadata.GetLevel(module))
}

// GetAllLevels - getting all set log levels
//  Returns:
//  module names and their associated logging levels
//
// If not set default logging level is info
func GetAllLevels() map[string]Level {
	metadataLevels := metadata.GetAllLevels()

	// Convert to the Level type in this package
	levels := make(map[string]Level)
	for module, logLevel := range metadataLevels {
		levels[module] = Level(logLevel)
	}

	return levels
}

// IsEnabledFor - Check if given log level is enabled for given module
//  Parameters:
//  module is module name
//  level is logging level
//
//  Returns:
//  is logging enabled for this module and level
//
// If not set default logging level is info
func IsEnabledFor(module string, level Level) bool {
	return metadata.IsEnabledFor(module, metadata.Level(level))
}

// ParseLevel returns the log level from a string representation.
//  Parameters:
//  level is logging level in string representation
//
//  Returns:
//  logging level
func ParseLevel(level string) (Level, error) {
	l, err := metadata.ParseLevel(level)
	return Level(l), err
}

// ParseString returns string representation of given log level
//  Parameters:
//  level is logging level represented as an int
//
//  Returns:
//  logging level in string representation
func ParseString(level Level) string {
	return metadata.ParseString(metadata.Level(level))
}

// ShowCallerInfo - Show caller info in log lines for given log level and module
//  Parameters:
//  module is module name
//  level is logging level
//
// note: based on implementation of custom logger, callerinfo info may not be available for custom logging provider
func ShowCallerInfo(module string, level Level) {
	metadata.ShowCallerInfo(module, metadata.Level(level))
}

// HideCallerInfo - Do not show caller info in log lines for given log level and module
//  Parameters:
//  module is module name
//  level is logging level
//
// note: based on implementation of custom logger, callerinfo info may not be available for custom logging provider
func HideCallerInfo(module string, level Level) {
	metadata.HideCallerInfo(module, metadata.Level(level))
}

// IsCallerInfoEnabled - returns if caller info enabled for given log level and module
//  Parameters:
//  module is module name
//  level is logging level
//
//  Returns:
//  is caller info enabled for this module and level
//
// note: based on implementation of custom logger, callerinfo info may not be available for custom logging provider
func IsCallerInfoEnabled(module string, level Level) bool {
	return metadata.IsCallerInfoEnabled(module, metadata.Level(level))
}
