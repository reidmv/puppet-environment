package filesync

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	httpClient *http.Client
)

func InitializeHttpClient(certFile, keyFile string) error {
	tlsCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	tlsConfig := tls.Config{
		Certificates:       []tls.Certificate{tlsCert},
		InsecureSkipVerify: true,
	}

	transport := http.Transport{TLSClientConfig: &tlsConfig}
	httpClient = &http.Client{Transport: &transport}

	// If we got this far, no errors occurred
	return nil
}

func FileSyncCommit() error {
	if httpClient == nil {
		return errors.New("HTTP client not initialized")
	}

	req, err := http.NewRequest("GET", "https://localhost:8140/file-sync/v1/commit", bytes.NewBuffer([]byte(`{"commit-all": true}`)))
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(body)
	return nil
}

func FileSyncForceSync() error {
	if httpClient == nil {
		return errors.New("HTTP client not initialized")
	}

	return nil
}
