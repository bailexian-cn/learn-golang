package upload

import (
	"io"
	"net/http"
)

func Upload(uploadUrl string, reader io.Reader) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, uploadUrl, reader)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
