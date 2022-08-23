package poker

import (
	"io"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "1234")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("123"))

	file.Seek(0, 0)
	newFileContent, _ := io.ReadAll(file)

	got := string(newFileContent)
	want := "123"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
