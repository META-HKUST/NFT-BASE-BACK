package utils

// HTTPError example
type Error struct {
	Code    string `json:"code" example:"USER_NAME_EXIST"`
	Message string `json:"message" example:"user name exist"`
}
