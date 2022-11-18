package requests

import "net/http"

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (c *UpdateUserRequest) Bind(_ *http.Request) error { return nil }
