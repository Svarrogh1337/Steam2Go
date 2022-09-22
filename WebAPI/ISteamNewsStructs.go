package Steam2Go

type GetAppNewsResponseV1 struct {
	Appnews AppNewsV1 `json:"appnews"`
}

type AppNewsV1 struct {
	Appid     int         `json:"appid"`
	Newsitems NewsItemsV1 `json:"newsitems"`
	Count     int         `json:"count"`
}

type NewsItemsV1 struct {
	Newsitem []NewsItem `json:"newsitem"`
}

type NewsItem struct {
	Gid           string `json:"gid"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	IsExternalURL bool   `json:"is_external_url"`
	Author        string `json:"author"`
	Contents      string `json:"contents"`
	Feedlabel     string `json:"feedlabel"`
	Date          int    `json:"date"`
	Feedname      string `json:"feedname"`
	FeedType      int    `json:"feed_type"`
	Appid         int    `json:"appid"`
}

type GetAppNewsResponseV2 struct {
	Appnews AppNewsV2 `json:"appnews"`
}

type AppNewsV2 struct {
	Appid     int           `json:"appid"`
	Newsitems []NewsItemsV2 `json:"newsitems"`
	Count     int           `json:"count"`
}

type NewsItemsV2 struct {
	Gid           string   `json:"gid"`
	Title         string   `json:"title"`
	URL           string   `json:"url"`
	IsExternalURL bool     `json:"is_external_url"`
	Author        string   `json:"author"`
	Contents      string   `json:"contents"`
	Feedlabel     string   `json:"feedlabel"`
	Date          int      `json:"date"`
	Feedname      string   `json:"feedname"`
	FeedType      int      `json:"feed_type"`
	Appid         int      `json:"appid"`
	Tags          []string `json:"tags,omitempty"`
}

type GetAppNewsOptions struct {
	Maxlength int    `json:"maxlength"`
	Enddate   int64  `json:"enddate"`
	Count     int    `json:"count"`
	Tags      string `json:"tags,omitempty"`
}
