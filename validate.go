package validateGo

import (
	"errors"
)

func TwoSum(a, b int) int {
	return a + b
}

type createAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

type updateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

func validateUpdate(u createAdRequest) error {
	if u.Title == "" || len(u.Title) > 100 {
		return errors.New("Title incorrect")
	}
	if u.Text == "" || len(u.Text) > 500 {
		return errors.New("Text incorrect")
	}
	return nil
}

func validateCreate(u updateAdRequest) error {
	if u.Title == "" || len(u.Title) > 100 {
		return errors.New("Title incorrect")
	}
	if u.Text == "" || len(u.Text) > 500 {
		return errors.New("Text incorrect")
	}
	return nil
}
