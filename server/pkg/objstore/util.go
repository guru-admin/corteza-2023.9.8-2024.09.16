package objstore

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/url"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/http"
)

func FromURL(fileURL string) (io.ReadCloser, error) {
	if u, err := url.ParseRequestURI(fileURL); err != nil {
		return nil, errors.WithStack(err)
	} else if u.Scheme != "https" {
		return nil, fmt.Errorf("Only HTTPS is supported for file uploads")
	}

	client, err := http.New(&http.Config{
		Timeout: 10,
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := client.Get(fileURL)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func FromMultipartFile(file *multipart.FileHeader) (io.ReadCloser, error) {
	reader, err := file.Open()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return reader, nil
}

func FromAny(file *multipart.FileHeader, url string) (io.ReadCloser, error) {
	if file != nil {
		return FromMultipartFile(file)
	}
	return FromURL(url)
}
