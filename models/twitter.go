package models

type Tweet struct {
	ID         int64  `json:"id"`
	Text       string `json:"text"`
	Timestamp  int64  `json:"timestamp_ms"`
	ScreenName string `json:"screen_name"`
}
