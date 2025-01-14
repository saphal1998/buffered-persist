package persist_test

import (
	"bytes"
	"crypto/rand"
	"os"
	"testing"

	"github.com/saphal1998/persist"
)

func TestSaveAndLoadFile(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "testsavefile")
	defer os.RemoveAll(file.Name())
	if err != nil {
		t.Errorf("Cannot create temporary file")
	}

	data := make([]byte, 50000000)
	rand.Read(data)
	persist.Save(data, file.Name())

	contents, err := persist.Load(file.Name())
	if err != nil {
		t.Errorf("Cannot read from file")
	}

	if len(contents) != len(data) || !bytes.Equal(data, contents) {
		t.Errorf("expected=%v, got=%v", data, contents)

	}
}
