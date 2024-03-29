package datastore

type Profile struct {
	Pid                  PidType `json:"pid"`
	Name                 string  `json:"name,omitempty"`
	PasswordHash         []byte  `json:"-"`
	Bio                  string  `json:"bio,omitempty"`
	Email                string  `json:"email,omitempty"`
	Joined               int64   `json:"joined,omitempty"`
	Location             string  `json:"location,omitempty"`
	Url                  string  `json:"url,omitempty"`
	ProfileImageUrl      string  `json:"profileimageurl,omitempty"`
	ProfileImageUrlHttps string  `json:"profileimageurlhttps,omitempty"`
	PossiblyCount        int     `json:"pcount"`
	MaybeCount           int     `json:"mcount"`
	FollowerCount        int     `json:"followercount"`
	FollowingCount       int     `json:"followingcount"`
	FeedCount            int     `json:"feedcount"`
	FeedType             string  `json:"feedtype"`
	FeedUrl              string  `json:"feedurl,omitempty"`
	ParentPid            PidType `json:"parentpid,omitempty"`
	ItemType             string  `json:"itemtype,omitempty"`
}

type ScoredProfile struct {
	Pid   string  `json:"pid"`
	Score float64 `json:"score"`
}

type FollowingProfile struct {
	Profile
	Reciprocal bool `json:"reciprocal"`
}

type BriefProfile struct {
	Pid                  PidType `json:"pid"`
	Name                 string  `json:"name,omitempty"`
	ProfileImageUrlHttps string  `json:"profileimageurlhttps,omitempty"`
}
