package pkg

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"

	"golang.org/x/net/html/charset"
)

func Request(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func XMLUnmarshalUnUTF(content []byte, target any) error {
	r := bytes.NewReader(content)
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel

	if err := d.Decode(&target); err != nil {
		return err
	}
	return nil
}
