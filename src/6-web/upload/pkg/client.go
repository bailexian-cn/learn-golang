package pkg

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func Upload(uploadUrl string, reader io.Reader) error {
	client := &http.Client{}
	mr := multipart.NewReader(reader, "chart")

	req, err := http.NewRequest(http.MethodPut, uploadUrl, mr)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}