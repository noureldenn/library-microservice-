package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BookInfo struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	AvailableCopies int    `json:"available_copies"`
}

type BookClient struct {
	BaseURL string
}

func NewBookClient(url string) *BookClient {
	return &BookClient{BaseURL: url}
}

func (bc *BookClient) GetBook(bookID uint) (*BookInfo, error) {
	url := fmt.Sprintf("%s/books/%d", bc.BaseURL, bookID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("book not found")
	}

	var book BookInfo
	if err := json.NewDecoder(resp.Body).Decode(&book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (bc *BookClient) DecreaseAvailable(bookID uint) error {
	url := fmt.Sprintf("%s/books/%d/decrease", bc.BaseURL, bookID)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to decrease available copies")
	}

	return nil
}
