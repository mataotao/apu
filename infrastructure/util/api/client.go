package api

import (
	"time"

	resty "github.com/go-resty/resty/v2"
)

var Client *resty.Client

func init() {
	Client = resty.New()
	Client.SetTimeout(5 * time.Minute)
}
