package github

type GithubProfile struct {
	Login               string  `json:"login"`               //: "anishjain94",
	Id                  int64   `json:"id"`                  //: 25664697,
	Node_id             string  `json:"node_id"`             //: "MDQ6VXNlcjI1NjY0Njk3",
	Avatar_url          string  `json:"avatar_url"`          //: "https://avatars.githubusercontent.com/u/25664697?v=4",
	Gravatar_id         string  `json:"gravatar_id"`         //: "",
	Url                 string  `json:"url"`                 //: "https://api.github.com/users/anishjain94",
	Html_url            string  `json:"html_url"`            //: "https://github.com/anishjain94",
	Followers_url       string  `json:"followers_url"`       //: "https://api.github.com/users/anishjain94/followers",
	Following_url       string  `json:"following_url"`       //: "https://api.github.com/users/anishjain94/following{/other_user}",
	Gists_url           string  `json:"gists_url"`           //: "https://api.github.com/users/anishjain94/gists{/gist_id}",
	Starred_url         string  `json:"starred_url"`         //: "https://api.github.com/users/anishjain94/starred{/owner}{/repo}",
	Subscriptions_url   string  `json:"subscriptions_url"`   //: "https://api.github.com/users/anishjain94/subscriptions",
	Organizations_url   string  `json:"organizations_url"`   //: "https://api.github.com/users/anishjain94/orgs",
	Repos_url           string  `json:"repos_url"`           //: "https://api.github.com/users/anishjain94/repos",
	Events_url          string  `json:"events_url"`          //: "https://api.github.com/users/anishjain94/events{/privacy}",
	Received_events_url string  `json:"received_events_url"` //: "https://api.github.com/users/anishjain94/received_events",
	Type                string  `json:"type"`                //: "User",
	Site_admin          *bool   `json:"site_admin"`          //: false,
	Name                *string `json:"name"`                //: "Anish Jain",
	Company             *string `json:"company"`             //: "@pice",
	Blog                *string `json:"blog"`                //: "pice.one",
	Location            *string `json:"location"`            //: "bangalore",
	Email               *string `json:"email"`               //: null,
	Hireable            bool    `json:"hireable"`            //: true,
	Bio                 string  `json:"bio"`                 //: "All things Backend@Pice",
	Twitter_username    string  `json:"twitter_username"`    //: "Jainanish94",
	Public_repos        int64   `json:"public_repos"`        //: 8,
	Public_gists        int64   `json:"public_gists"`        //: 0,
	Followers           int64   `json:"followers"`           //: 11,
	Following           int64   `json:"following"`           //: 15,
	Created_at          string  `json:"created_at"`          //: "2017-02-09T13:56:46Z",
	Updated_at          string  `json:"updated_at"`          //: "2023-01-28T12:43:45Z"
}

type GithubRepo struct {
	Id          string `json:"id"`
	NodeId      string `json:"node_id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Private     string `json:"private"`
	Html_url    string `json:"html_url"`
	Description string `json:"description"`
	Fork        string `json:"fork"`
	Url         string `json:"url"`
}
