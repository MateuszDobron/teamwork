package customerimporter

import (
	"encoding/csv"
	"fmt"
	"slices"
	"strconv"
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

func (dc DomainCounts) CSVDomainCounts(csvWriter *csv.Writer) error {
	pair := make([]string, 2)
	for _, key := range dc.sortKeys() {
		pair[0] = key
		pair[1] = strconv.FormatUint(dc.DomainMap[key], 10)
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
