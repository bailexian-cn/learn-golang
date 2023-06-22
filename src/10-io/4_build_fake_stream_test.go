package __io

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestBuildFakeReader(t *testing.T) {
	writer := bytes.Buffer{}
	reader := strings.NewReader("haha")
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("read err")
		}
		writer.Write(buf[:n])
	}
	if writer.String() != "haha" {
		t.Error(fmt.Sprintf("except str is haha, but is %s", writer.String()))
	}
}







