package main

import (
	"flag"
	"os"

	"log/slog"

	"github.com/MateuszDobron/teamwork/customerimporter"
	"github.com/MateuszDobron/teamwork/exporter"
)

type Options struct {
	path    *string
	outFile *string
}

func readOptions() *Options {
	opts := &Options{}
	// default value is not nil
	opts.path = flag.String("path", "./customers.csv", "Path to the file with customer data")
	opts.outFile = flag.String("out", "", "Optional: output file path. If empty program will output results to the terminal")
	flag.Parse()
	return opts
}

func main() {
	opts := readOptions()
	importer := customerimporter.NewCustomerImporter(*opts.path)
	data, err := importer.ImportDomainData()
	if err != nil {
		slog.Error("error importing customer data: ", slog.Any("err", err))
		os.Exit(1)
	}
	if *opts.outFile == "" {
		data.PrintDomainCounts()
	} else {
		exporter := exporter.NewCustomerExporter(*opts.outFile)
		if saveErr := exporter.ExportData(data); saveErr != nil {
			slog.Error("error saving domain data: ", slog.Any("err", saveErr))
			os.Exit(1)
		}
	}
}
