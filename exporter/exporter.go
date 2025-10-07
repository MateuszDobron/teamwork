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
		return fmt.Errorf("error provided data is empty 0 length")
	}
	outputFile, err := os.Create(ex.outputPath)
	if err != nil {
		return fmt.Errorf("error creating new file for saving: %v", err)
	}
	defer outputFile.Close()
	return exportCsv(data, outputFile)
}

func exportCsv(data customerimporter.DomainCounts, output io.Writer) error {
	headers := []string{"domain", "number_of_customers"}
	csvWriter := csv.NewWriter(output)
	defer func() error {
		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			return err
		}
		return nil
	}()
	if err := csvWriter.Write(headers); err != nil {
		return err
	}
	return data.CsvDomainCounts(csvWriter)
}
