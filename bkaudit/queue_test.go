package bkaudit

import (
	"os"
	"testing"
)

type fileExporter struct{}

func (e *fileExporter) Export(queue BaseQueue) {
	file, _ := os.OpenFile("audit.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer func() { _ = file.Close() }()
	for event := range queue {
		data, _ := event.String()
		_, _ = file.Write([]byte(data + "\n"))
	}
}

func BenchmarkExport(b *testing.B) {
	client, _ := InitEventClient("", "", &Formatter{}, []BaseExporter{&fileExporter{}}, 0, nil)
	b.ResetTimer()
	runTest(client, b.N)

}

type noExporter struct{}

func (e *noExporter) Export(queue BaseQueue) {
	for event := range queue {
		_, _ = event.String()
	}
}

func BenchmarkNoExport(b *testing.B) {
	client, _ := InitEventClient("", "", &Formatter{}, []BaseExporter{&noExporter{}}, 0, nil)
	b.ResetTimer()
	runTest(client, b.N)
}

func runTest(client *EventClient, times int) {
	for i := 0; i < times; i++ {
		client.AddEvent(
			&AuditAction{ActionID: "test"},
			&AuditResource{},
			&AuditInstance{},
			&AuditContext{Username: "admin"},
			"",
			"",
			0,
			0,
			0,
			"",
			map[string]any{},
		)
	}
}
