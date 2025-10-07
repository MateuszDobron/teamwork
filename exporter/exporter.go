package exporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/MateuszDobron/teamwork/customerimporter"
)

type CustomerExporter struct {
	outputPath string
}

// NewCustomerExporter returns a new CustomerExporter that writes customer domain data to specified file.
func NewCustomerExporter(outputPath string) CustomerExporter {
	return CustomerExporter{
		outputPath: outputPath,
	}
}

// ExportData writes sorted customer domain data to a CSV file. If file already exists, it will
// be truncated.
func (ex CustomerExporter) ExportData(data customerimporter.DomainCounts) error {
	if len(data.DomainMap) == 0 {
		return fmt.Errorf("provided data is empty 0 length")
	}
	outputFile, err := os.Create(ex.outputPath)
	if err != nil {
		return fmt.Errorf("creating new file for saving: %v at path: %s", err, ex.outputPath)
	}
	defer outputFile.Close()
	return exportCsv(data, outputFile)
}

func exportCsv(data customerimporter.DomainCounts, output io.Writer) (err error) {
	headers := []string{"domain", "number_of_customers"}
	csvWriter := csv.NewWriter(output)
	defer func() {
		csvWriter.Flush()
		if ferr := csvWriter.Error(); err == nil && ferr != nil {
			err = ferr
		}
	}()
	if err = csvWriter.Write(headers); err != nil {
		return
	}
	if err = data.CSVDomainCounts(csvWriter); err != nil {
		return
	}
	return
}
