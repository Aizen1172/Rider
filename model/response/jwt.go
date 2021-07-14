package response

import "riderStyemProject/model"

type Login struct {
	ExpiresAt int64  `json:"expires_at"`
	Token     string `json:"token"`
	User      *model.User `json:"user"`
}
