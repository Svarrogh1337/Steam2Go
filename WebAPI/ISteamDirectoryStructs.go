package Steam2Go

type GetCMListForConnectResponseV1 struct {
	Response ResponseV1 `json:"response"`
}

type ResponseV1 struct {
	Serverlist []ServerlistV1 `json:"serverlist"`
	Success    bool           `json:"success"`
	Message    string         `json:"message"`
}

type ServerlistV1 struct {
	Endpoint       string  `json:"endpoint"`
	LegacyEndpoint string  `json:"legacy_endpoint"`
	Type           string  `json:"type"`
	Dc             string  `json:"dc"`
	Realm          string  `json:"realm"`
	Load           int     `json:"load"`
	WtdLoad        float64 `json:"wtd_load"`
}

type GetCMListForConnectOptions struct {
	Cellid   int    `json:"cellid"`
	Cmtype   string `json:"cmtype"`
	Realm    string `json:"realm"`
	Maxcount int    `json:"maxcount"`
}
