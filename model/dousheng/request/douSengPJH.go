package request


//请求入参
type GetFeed struct {
	LatestTime string `json:"latest_time"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token 	string `json:"token"` // 用户登录状态下设置
}