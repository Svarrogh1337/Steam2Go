package Steam2Go

type GetAppListResponseV1 struct {
	Applist ApplistV1 `json:"applist"`
}

type ApplistV1 struct {
	Apps AppsV1 `json:"apps"`
}

type AppsV1 struct {
	App []AppV1 `json:"app"`
}

type AppV1 struct {
	Appid int    `json:"appid"`
	Name  string `json:"name"`
}
