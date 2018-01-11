package sanitize

import (
	"path/filepath"
	"testing"
)

func TestIsSane(t *testing.T) {
	var tc = []struct {
		path string
		sane bool
	}{
		{"../../../", false},
		{"/etc/passwd", false},
		{"../etc/passwd/", false},
		{"testdir/cat/catpicture", true},
		{"./testdir/cat/catpicture", true},
	}

	for _, tt := range tc {
		rootPath, err := filepath.Abs("testdir/")
		if err != nil {
			t.Fatal(err)
		}
		sane, err := IsSane(rootPath, tt.path)
		if err != nil {
			t.Fatal(err)
		}
		if sane != tt.sane {
			t.Errorf("Sane is %v, should be %v for path %v", sane, tt.sane, tt.path)
		}
	}
}

func TestIsSaneAbsOnly(t *testing.T) {
	rootPath := "testdata/"
	_, err := IsSane(rootPath, "./")
	if err == nil {
		t.Error("Did not get error for supplying a relative root path")
	}
}
