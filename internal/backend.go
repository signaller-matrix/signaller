package internal

import (
	"github.com/nxshock/signaller/internal/models"
	"github.com/nxshock/signaller/internal/models/sync"
)

type Backend interface {
	Register(username, password, device string) (token string, error *models.ApiError)
	Login(username, password, device string) (token string, err *models.ApiError)
	Logout(token string) *models.ApiError
	Sync(token string, request sync.SyncRequest) (response *sync.SyncReply, err *models.ApiError)
}
