package message

type Open struct {
	SId          string   `json:"sid"`
	Upgrades     []string `json:"upgrades"`
	PingInterval int      `json:"pingInterval"`
	PingTimeout  int      `json:"pingTimeout"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
