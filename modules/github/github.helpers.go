package github

import (
	"encoding/json"
	"fmt"

	"github.com/anishjain94/githubApi/infra"
)

func StoreData(userName string) {

	userName = "/" + userName
	gitProfileResponse := GithubProfile{}
	_, err := infra.GithubClient.R().
		SetResult(&gitProfileResponse).
		Get(userName)

	if err != nil {
		print(err) //Handle errors properly & Check for rate limitation
	}

	data, _ := json.MarshalIndent(gitProfileResponse, "", "    ")
	CreateFile("../logs/profile.log", string(data))

	gitRepoResponse := []GithubRepo{}
	_, err = infra.GithubClient.R().
		SetResult(&gitRepoResponse).
		Get(userName + "/repos")

	if err != nil {
		fmt.Println(err) //Handle errors properly & Check for rate limitation
	}
	data, _ = json.MarshalIndent(gitRepoResponse, "", "    ")
	CreateFile("../logs/repo.log", string(data))
}
