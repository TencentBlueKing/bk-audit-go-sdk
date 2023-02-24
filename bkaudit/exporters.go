package bkaudit

// BaseExporter - Interface for Exporter
type BaseExporter interface {
	Export(queue BaseQueue)
}

// Exporter - Build in Exporter
type Exporter struct{}

// Export - Export Audit Event to Log
func (e *Exporter) Export(queue BaseQueue) {
	for event := range queue {
		// get string data
		data, err := event.String()
		if err != nil {
			RuntimeLog.Error("export event failed: ", err)
			return
		}
		// Directly Export to EventLog
		EventLog.Info(data)
	}
}
