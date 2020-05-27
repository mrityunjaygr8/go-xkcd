package model

import (
	"encoding/json"
	"fmt"
)

type ComicResponse struct {
	Num        int    `json:"num"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	SafeTitle  string `json:"safe_title"`
	Link       string `json:"link"`
	News       string `json:"news"`
	Transcript string `json:"transcript"`
}

type Comic struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}


func (cr ComicResponse) FormattedDate() string {
	return fmt.Sprintf("%s-%s-%s", cr.Day, cr.Month, cr.Year)
}

func (cr ComicResponse) Comic() Comic {
	return Comic{
		Title: cr.Title,
		Number: cr.Num,
		Date: cr.FormattedDate(),
		Description: cr.Alt,
		Image: cr.Img,
	}
}

func (c Comic) PrettyPrint() string {
	p := fmt.Sprintf(
		"Title: %s\nComic No.: %d\nDate: %s\nDescription: %s\nImage: %s\n",
		c.Title, c.Number, c.Date, c.Description, c.Image)

	return p
}

func (c Comic) JSON() string {
	cJSON, err := json.Marshal(c)
	if err != nil {
		return ""
	}

	return string(cJSON)
}
