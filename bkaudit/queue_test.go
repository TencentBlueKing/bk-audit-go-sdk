package bkaudit

import (
	"os"
	"sync"
	"testing"
)

type fileExporter struct {
	file *os.File
}

func (e *fileExporter) Export(queue Queue, wg *sync.WaitGroup) {
	defer wg.Done()
	e.file, _ = os.OpenFile("audit.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer func() { _ = e.file.Close() }()
	for event := range queue {
		data, _ := event.String()
		_, _ = e.file.Write([]byte(data + "\n"))
	}
}

func (e *fileExporter) Validate() bool {
	_, err := e.file.Stat()
	return err != nil
}

func BenchmarkExport(b *testing.B) {
	client, _ := InitEventClient("", "", &EventFormatter{}, []Exporter{&fileExporter{}}, 0, nil)
	b.ResetTimer()
	runTest(client, b.N)
	client.Exit()
}

type noExporter struct{}

func (e *noExporter) Export(queue Queue, wg *sync.WaitGroup) {
	defer wg.Done()
	for event := range queue {
		_, _ = event.String()
	}
}

func (e *noExporter) Validate() bool {
	return true
}

func BenchmarkNoExport(b *testing.B) {
	client, _ := InitEventClient("", "", &EventFormatter{}, []Exporter{&noExporter{}}, 0, nil)
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
