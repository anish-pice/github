package infra

import "github.com/go-resty/resty/v2"

var (
	GithubClient *resty.Client
)

func InitClient() {

	GithubClient = resty.New().SetBaseURL("https://api.github.com/users").EnableTrace()
}
