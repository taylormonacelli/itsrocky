package daggerverse

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func GetAuthor(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	x := strings.TrimPrefix(path, "https://")

	fullURL := "https://" + x

	u, err := url.Parse(fullURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	p := strings.TrimPrefix(u.Path, "/")
	z := strings.Index(p, "/")
	author := p[:z]

	return author, nil
}

func getHost(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	x := strings.TrimPrefix(path, "https://")

	fullURL := "https://" + x

	u, err := url.Parse(fullURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	return u.Host, nil
}

func GetAuthorRepoURL(path string) (string, error) {
	author, err := GetAuthor(path)
	if err != nil {
		return "", fmt.Errorf("error getting author: %v", err)
	}

	host, err := getHost(path)
	if err != nil {
		return "", fmt.Errorf("error getting author: %v", err)
	}

	y := "https://" + host + "/" + author

	return y, nil
}

func (c CustomizedRepositoryInfo) String() (string, error) {
	jsonBytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %v", err)
	}
	return string(jsonBytes), nil
}

func (cr CustomizedRepositoryInfos) String() string {
	// Use json.MarshalIndent for the entire slice
	jsonBytes, err := json.MarshalIndent(cr, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v\n", err)
	}
	return string(jsonBytes)
}
