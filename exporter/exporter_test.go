package exporter

import (
	"fmt"
	"testing"

	"github.com/MateuszDobron/teamwork/customerimporter"
)

func TestExportData(t *testing.T) {
	path := "./testdata/test_output.csv"
	dc := customerimporter.NewDomainCounts()
	dc.DomainMap = map[string]uint64{
		"livejournal.com": 12,
		"microsoft.com":   22,
		"newsvine.com":    15,
		"pinteres.uk":     10,
		"yandex.ru":       43,
	}
	exporter := NewCustomerExporter(path)

	err := exporter.ExportData(dc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExportInvalidPath(t *testing.T) {
	path := ""
	exporter := NewCustomerExporter(path)

	err := exporter.ExportData(customerimporter.DomainCounts{})
	if err == nil {
		t.Fatal(err)
	}
	t.Log(err)
}

func TestExportEmptyData(t *testing.T) {
	path := "./test_output.csv"
	exporter := NewCustomerExporter(path)

	err := exporter.ExportData(customerimporter.NewDomainCounts())
	if err == nil {
		t.Fatal(err)
	}
	t.Log(err)
}

func BenchmarkImportDomainData(b *testing.B) {
	b.StopTimer()
	dir := b.TempDir()
	path := fmt.Sprintf("%s/test_output.csv", dir)
	dataPath := "../customerimporter/benchmark10k.csv"
	importer := customerimporter.NewCustomerImporter(dataPath)
	data, err := importer.ImportDomainData()
	if err != nil {
		b.Error(err)
	}
	exporter := NewCustomerExporter(path)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := exporter.ExportData(data); err != nil {
			b.Fatal(err)
		}
	}
}
