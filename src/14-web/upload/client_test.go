package upload

import (
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	file, err := os.Open("C:\\Users\\34552\\Downloads\\redis-1.1.15.tgz")
	if err != nil {
		t.Errorf(err.Error())
		_ = file.Close()
		return
	}
	err = Upload("http://localhost:8000/upload", file)
	if err != nil {
		t.Errorf(err.Error())
	}
	_ = file.Close()
}
