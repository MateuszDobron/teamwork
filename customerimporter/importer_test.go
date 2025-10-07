package customerimporter

import "testing"

func TestImportData(t *testing.T) {
	path := "./testdata/test_data.csv"
	importer := NewCustomerImporter(path)

	_, err := importer.ImportDomainData()
	if err != nil {
		t.Error(err)
	}
}

func TestImportDataSort(t *testing.T) {
	sortedDomains := []string{"360.cn", "acquirethisname.com", "blogtalkradio.com", "chicagotribune.com", "cnet.com", "cyberchimps.com", "github.io", "hubpages.com", "rediff.com", "statcounter.com"}
	path := "./testdata/test_data.csv"
	importer := NewCustomerImporter(path)
	data, err := importer.ImportDomainData()
	if err != nil {
		t.Error(err)
	}
	for i, v := range data.sortKeys() {
		if v != sortedDomains[i] {
			t.Errorf("data not sorted properly. mismatch:\nhave: %v\nwant: %v", v, sortedDomains[i])
		}
	}
}

func TestImportInvalidPath(t *testing.T) {
	path := ""
	importer := NewCustomerImporter(path)

	_, err := importer.ImportDomainData()
	if err == nil {
		t.Error("invalid path error not caught")
	}
}

func TestImportInvalidData(t *testing.T) {
	path := "./testdata/test_invalid_data.csv"
	importer := NewCustomerImporter(path)

	_, err := importer.ImportDomainData()
	if err == nil {
		t.Error("invalid data not caught")
	}
}

func TestExtractDomain(t *testing.T) {
	tests := []struct {
		in      string
		wantDom string
		wantOK  bool
	}{
		{"user@example.com", "example.com", true},
		{"a@b", "b", true},
		{"@example.com", "", false},    // empty local part
		{"user@", "", false},           // empty domain
		{"userexample.com", "", false}, // no '@'
		{"", "", false},                // empty string
	}

	for _, tt := range tests {
		got, ok := extractDomain(tt.in)
		if got != tt.wantDom || ok != tt.wantOK {
			t.Errorf("extractDomain(%q) = (%q,%v), want (%q,%v)",
				tt.in, got, ok, tt.wantDom, tt.wantOK)
		}
	}
}

func BenchmarkImportDomainData(b *testing.B) {
	b.StopTimer()
	path := "./testdata/benchmark10k.csv"
	importer := NewCustomerImporter(path)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := importer.ImportDomainData(); err != nil {
			b.Error(err)
		}
	}
}
