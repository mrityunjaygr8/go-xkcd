package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
	"github.com/mrityunjaygr8/go-xkcd/model"
)

type ComicNumber int

type XKCDClient struct {
	client *http.Client
	baseURL string
}

const (
	BaseURL string = "https://xkcd.com"
	DefaultClientTimeout time.Duration = 30 * time.Second
	LatestComic ComicNumber = 0
)

func NewXKCDClient() *XKCDClient {
	return &XKCDClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

func (hc *XKCDClient) SetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

func (hc *XKCDClient) Fetch(n ComicNumber, save bool) (model.Comic, error) {
	resp, err := hc.client.Get(hc.buildURL(n))
	if err != nil {
		return model.Comic{}, err
	}
	defer resp.Body.Close()

	var comicResponse model.ComicResponse
	if err := json.NewDecoder(resp.Body).Decode(&comicResponse); err != nil {
		return model.Comic{}, err
	}

	if save {
		if err := hc.SaveToDisk(comicResponse.Img, "."); err != nil {
			fmt.Println("Failed to save image to disk!")
		}
	}

	return comicResponse.Comic(), nil
}

func(hc *XKCDClient) SaveToDisk(url, savePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	absSavePath, _ := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s", absSavePath, path.Base(url))

	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (hc *XKCDClient) buildURL(n ComicNumber) string {
	var finalURL string
	if n == LatestComic {
		finalURL = fmt.Sprintf("%s/info.0.json", hc.baseURL)
	} else {
		finalURL = fmt.Sprintf("%s/%d/info.0.json", hc.baseURL, n)
	}

	return finalURL
}
