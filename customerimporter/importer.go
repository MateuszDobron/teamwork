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
	"strings"
)

type CustomerImporter struct {
	path string
}

// NewCustomerImporter returns a new CustomerImporter that reads from file at specified path.
func NewCustomerImporter(filePath string) CustomerImporter {
	return CustomerImporter{
		path: filePath,
	}
}

// ImportDomainData reads and returns sorted customer domain data from CSV file.
func (ci CustomerImporter) ImportDomainData() (DomainCounts, error) {
	file, err := os.Open(ci.path)
	if err != nil {
		return DomainCounts{}, fmt.Errorf("improper path: %s caused: %v", ci.path, err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	// ReuseRecord avoids allocating a new []string for every row.
	// The returned slice is overwritten on each Read.
	csvReader.ReuseRecord = true

	domainCounts := NewDomainCounts()

	// skip first line with headers
	line, readErr := csvReader.Read()
	if readErr != nil {
		return DomainCounts{}, fmt.Errorf("invalid line: %s caused %v", line, readErr)
	}
	for line, readErr := csvReader.Read(); readErr != io.EOF; line, readErr = csvReader.Read() {
		if readErr != nil {
			return DomainCounts{}, readErr
		}
		domain, ok := extractDomain(line[2])
		if !ok {
			return DomainCounts{}, fmt.Errorf("invalid email address: %s", line[2])
		}
		domainCounts.DomainMap[domain] += 1
	}
	return domainCounts, nil
}

// extractDomain returns the domain part of an email.
// It avoids allocating an extra substring for the not domain part.
// Returns "" and false if invalid.
func extractDomain(mail string) (string, bool) {
	i := strings.IndexByte(mail, '@')
	// i == -1 no '@'; i == 0 empty local; i == len(field)-1 empty domain
	if i <= 0 || i+1 >= len(mail) {
		return "", false
	}
	dom := mail[i+1:]
	return dom, true
}
