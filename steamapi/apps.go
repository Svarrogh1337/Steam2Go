package steamapi

import (
	"context"
	"fmt"
)

// GetAppListResponseV1 stores the data from GetAppList ( version 1 ) method  that is part of ISteamApps interface.
type GetAppListResponseV1 struct {
	Applist struct {
		Apps struct {
			App []struct {
				Appid int    `json:"appid"`
				Name  string `json:"name"`
			} `json:"app"`
		} `json:"apps"`
	} `json:"applist"`
}

// GetAppListResponseV2 stores the data from GetAppList ( version 2 ) method  that is part of ISteamApps interface.
type GetAppListResponseV2 struct {
	Applist struct {
		Apps []struct {
			Appid int    `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}

// GetSDRConfigResponseV1 stores the data from GetSDRConfig ( version 1 ) method  that is part of ISteamApps interface.
type GetSDRConfigResponseV1 struct {
	Revision   int      `json:"revision"`
	Certs      []string `json:"certs"`
	P2PShareIP struct {
		Cn      int `json:"cn"`
		Default int `json:"default"`
		Ru      int `json:"ru"`
	} `json:"p2p_share_ip"`
	Pops struct {
		Ams struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"ams"`
		Atl struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"atl"`
		Bom struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"bom"`
		Dfw struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"dfw"`
		Dxb struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"dxb"`
		Eat struct {
			Aliases  []string  `json:"aliases"`
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Tier     int       `json:"tier"`
		} `json:"eat"`
		Eze struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"eze"`
		Fra struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"fra"`
		Gru struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"gru"`
		Hkg struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"hkg"`
		Iad struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"iad"`
		Jnb struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"jnb"`
		Lax struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"lax"`
		Lhr struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"lhr"`
		Lim struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"lim"`
		Maa struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"maa"`
		Mad struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"mad"`
		Mdf1 struct {
			Desc           string    `json:"desc"`
			Geo            []float64 `json:"geo"`
			Partners       int       `json:"partners"`
			RelayPublicKey string    `json:"relay_public_key"`
			Relays         []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"mdf1"`
		Ord struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"ord"`
		Par struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"par"`
		Scl struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"scl"`
		Sea struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"sea"`
		Seo struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"seo"`
		Sgp struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"sgp"`
		Sto struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"sto"`
		Sto2 struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"sto2"`
		Syd struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"syd"`
		Tyo struct {
			Aliases  []string  `json:"aliases"`
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"tyo"`
		Tyo1 struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"tyo1"`
		Vie struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"vie"`
		Waw struct {
			Desc     string    `json:"desc"`
			Geo      []float64 `json:"geo"`
			Partners int       `json:"partners"`
			Relays   []struct {
				Ipv4      string `json:"ipv4"`
				PortRange []int  `json:"port_range"`
			} `json:"relays"`
			Tier int `json:"tier"`
		} `json:"waw"`
	} `json:"pops"`
	RelayPublicKey string          `json:"relay_public_key"`
	RevokedKeys    []string        `json:"revoked_keys"`
	TypicalPings   [][]interface{} `json:"typical_pings"`
	Success        int             `json:"success"`
}

// GetServersAtAddressResponseV1 stores the data from GetServersAtAddress ( version 1 ) method  that is part of ISteamApps interface.
type GetServersAtAddressResponseV1 struct {
	Response struct {
		Success bool `json:"success"`
		Servers []struct {
			Addr     string `json:"addr"`
			Gmsindex int    `json:"gmsindex"`
			Steamid  string `json:"steamid"`
			Appid    int    `json:"appid"`
			Gamedir  string `json:"gamedir"`
			Region   int    `json:"region"`
			Secure   bool   `json:"secure"`
			Lan      bool   `json:"lan"`
			Gameport int    `json:"gameport"`
			Specport int    `json:"specport"`
		} `json:"servers"`
	} `json:"response"`
}

// UpToDateCheckResponseV1 stores the data from UpToDateCheck ( version 1 ) method  that is part of ISteamApps interface.
// If the specific app is up-to-date, RequiredVersion will be set to 0 and Message will be ""
type UpToDateCheckResponseV1 struct {
	Response struct {
		Success           bool   `json:"success"`
		UpToDate          bool   `json:"up_to_date"`
		VersionIsListable bool   `json:"version_is_listable"`
		RequiredVersion   int    `json:"required_version,omitempty"`
		Message           string `json:"message,omitempty"`
	} `json:"response"`
}

// GetAppListV1 gets the complete list of public apps.
// This call has no additional parameters.
// This version is no longer officially supported. It will continue to be usable, but it's highly recommended that you use the latest version.
// Response - GetAppListResponseV1
func (c *Client) GetAppListV1(ctx context.Context) (*GetAppListResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetAppListResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamApps/GetAppList/v%d?key=%s", c.baseURL, apiVersion, c.apiKey)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

// GetAppListV2 gets the complete list of public apps.
// This call has no additional parameters.
// This call has previous versions which are no longer officially supported. They will continue to be usable, but it's highly recommended that you use the latest version.
// Change history:
// Version 2 - Removed the redundant "app" field.
// Response - GetAppListResponseV2
func (c *Client) GetAppListV2(ctx context.Context) (*GetAppListResponseV2, error) {
	const (
		apiVersion = 2
	)
	resp := GetAppListResponseV2{}
	apiURL := fmt.Sprintf("%s/ISteamApps/GetAppList/v%d?key=%s", c.baseURL, apiVersion, c.apiKey)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

// GetSDRConfigV1 gets the complete list of Steam Datagram Relay (SDR) servers for specified app.
// This call has mandatory parameter appid.
// This call has no additional parameters.
// Response - GetSDRConfigResponseV1
func (c *Client) GetSDRConfigV1(ctx context.Context, appid int) (*GetSDRConfigResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetSDRConfigResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamApps/GetSDRConfig/v%d?key=%s&appid=%d", c.baseURL, apiVersion, c.apiKey, appid)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

// GetServersAtAddressV1 gets the details of server/servers running at specified IPv4:port/IPv4
// This call has mandatory parameter addr.
// This call has no additional parameters.
// Response - GetSDRConfigResponseV1
func (c *Client) GetServersAtAddressV1(ctx context.Context, addr string) (*GetServersAtAddressResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetServersAtAddressResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamApps/GetServersAtAddress/v%d?key=%s&addr=%s", c.baseURL, apiVersion, c.apiKey, addr)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

// UpToDateCheckV1 gets the details of server/servers running at specified IPv4:port/IPv4
// This call has 2 mandatory parameters - appid and version.
// This call has no additional parameters.
// Response - GetSDRConfigResponseV1
// If the specific app is up-to-date, RequiredVersion will be set to 0 and Message will be ""
func (c *Client) UpToDateCheckV1(ctx context.Context, appid int, version int) (*UpToDateCheckResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := UpToDateCheckResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamApps/UpToDateCheck/v%d?key=%s&appid=%d&version=%d", c.baseURL, apiVersion, c.apiKey, appid, version)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
