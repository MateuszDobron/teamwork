// Package customerimporter reads from a CSV file and returns a sorted (data
// structure of your choice) of email domains along with the number of customers
// with e-mail addresses for each domain. This should be able to be ran from the
// CLI and output the sorted domains to the terminal or to a file. Any errors
// should be logged (or handled). Performance matters (this is only ~3k lines,
// but could be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DomainCounts struct {
	DomainMap map[string]uint64
}

func NewDomainCounts() DomainCounts {
	return DomainCounts{
		DomainMap: make(map[string]uint64),
	}
}

func (dc DomainCounts) sortKeys() []string {
	keys := make([]string, 0, len(dc.DomainMap))
	for k := range dc.DomainMap {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func (dc DomainCounts) CsvDomainCounts(csvWriter *csv.Writer) error {
	for _, key := range dc.sortKeys() {
		pair := []string{key, strconv.FormatUint(dc.DomainMap[key], 10)}
		if err := csvWriter.Write(pair); err != nil {
			return err
		}
	}
	return nil
}

func (dc DomainCounts) PrintDomainCounts() {
	fmt.Println("domain,number_of_customers")
	for _, key := range dc.sortKeys() {
		fmt.Printf("%s,%v\n", key, dc.DomainMap[key])
	}
}

type CustomerImporter struct {
	path *string
}

// NewCustomerImporter returns a new CustomerImporter that reads from file at specified path.
func NewCustomerImporter(filePath *string) *CustomerImporter {
	return &CustomerImporter{
		path: filePath,
	}
}

// ImportDomainData reads and returns sorted customer domain data from CSV file.
func (ci CustomerImporter) ImportDomainData() (DomainCounts, error) {
	file, err := os.Open(*ci.path)
	if err != nil {
		return DomainCounts{}, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	domainCounts := NewDomainCounts()

	// skip first line with headers
	line, readErr := csvReader.Read()
	if readErr != nil {
		fmt.Println(line, readErr)
		return DomainCounts{}, readErr
	}
	for line, readErr := csvReader.Read(); readErr != io.EOF; line, readErr = csvReader.Read() {
		if readErr != nil {
			return DomainCounts{}, readErr
		}
		email, domain, found := strings.Cut(line[2], "@")
		if email == "" || !found {
			return DomainCounts{}, fmt.Errorf("error invalid email address: %s", line[2])
		}
		domainCounts.DomainMap[domain] += 1
	}
	return domainCounts, nil
}
