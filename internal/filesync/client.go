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

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{tlsCert},
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	httpClient = &http.Client{Transport: transport}

	// If we got this far, no errors occurred
	return nil
}

func FileSyncCommit() error {
	if httpClient == nil {
		return errors.New("HTTP client not initialized")
	}

	// Perform `file-sync commit`
	req, err := http.NewRequest("POST", "https://localhost:8140/file-sync/v1/commit", bytes.NewBuffer([]byte(`{"commit-all": true}`)))
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("file-sync commit returned error: %v", body)
	}

	// Perform `file-sync force-sync`
	resp, err = httpClient.Post("https://localhost:8140/file-sync/v1/force-sync", "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("file-sync commit returned error: %v", body)
	}

	// Seems like it all worked
	return nil
}

func FileSyncForceSync() error {
	if httpClient == nil {
		return errors.New("HTTP client not initialized")
	}

	return nil
}
