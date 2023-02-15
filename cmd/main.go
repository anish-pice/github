package main

import (
	"github.com/anishjain94/githubApi/infra"
	"github.com/anishjain94/githubApi/modules/github"
)

func main() {
	infra.InitClient()
	userName := "jainanish94"
	github.StoreData(userName)
}
