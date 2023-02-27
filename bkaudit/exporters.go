package bkaudit

// Exporter - Interface for Exporter
type Exporter interface {
	Export(queue Queue)
	Validate() bool
}

// LoggerExporter - Build in Exporter
type LoggerExporter struct {
	Logger interface {
		Info(arg ...interface{})
	}
}

// Export - Export Audit Event to Log
func (e *LoggerExporter) Export(queue Queue) {
	for event := range queue {
		// get string data
		data, err := event.String()
		if err != nil {
			logger.Error("export event failed: ", err)
			return
		}
		// Directly Export to EventLog
		e.Logger.Info(data)
	}
}

// Validate - Validate Exporter
func (e *LoggerExporter) Validate() bool {
	if e.Logger == nil {
		logger.Error("logger of exporter unset")
		return false
	}
	return true
}
