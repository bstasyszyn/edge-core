/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package log

import (
	"sync"

	"github.com/trustbloc/edge-core/pkg/internal/logging/modlog"
)

// loggerProviderInstance is logger factory singleton - access only via loggerProvider()
//nolint:gochecknoglobals
var (
	loggerProviderInstance LoggerProvider
	mutex                  sync.RWMutex
)

// Initialize sets new custom logging provider which takes over logging operations.
// It is required to call this function before making any loggings for using custom loggers.
func Initialize(l LoggerProvider) {
	mutex.Lock()
	defer mutex.Unlock()

	loggerProviderInstance = l
	logger := loggerProviderInstance.GetLogger(loggerModule)
	logger.Debugf("Logger provider initialized")
}

func loggerProvider() LoggerProvider {
	mutex.RLock()
	provider := loggerProviderInstance
	mutex.RUnlock()

	if provider != nil {
		return provider
	}

	mutex.Lock()
	defer mutex.Unlock()

	// A custom logger must be initialized prior to the first log output
	// Otherwise the built-in logger is used
	loggerProviderInstance = &modlogProvider{}
	logger := loggerProviderInstance.GetLogger(loggerModule)
	logger.Debugf(loggerNotInitializedMsg)

	return loggerProviderInstance
}

// modlogProvider is a module based logger provider wrapped on given custom logging provider
// if custom logger provider is not provided, then default logger will be used
type modlogProvider struct {
	custom LoggerProvider
}

// GetLogger returns moduled logger implementation.
func (p *modlogProvider) GetLogger(module string) Logger {
	var logger Logger
	if p.custom != nil {
		logger = p.custom.GetLogger(module)
	} else {
		logger = modlog.NewDefLog(module)
	}

	return modlog.NewModLog(logger, module)
}
