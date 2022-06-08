package request


//响应参数

type DSResponse struct {
	StatusCode int `json:"status_code"`
	StatusMsg string `json:"status_msg"`
}

type GetFeedResponse struct {
	DSResponse
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time,omitempty"`
}

//视频
type Video struct {
	Id  int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`	//视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`					//视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

//用户信息
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

//评论
type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}