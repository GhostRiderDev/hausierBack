package cfg

// Config defines all app config to load
type ApiConfig struct {
	ServerCfg ServerConfig `json:"server"`
}

// ServerConfig defines config to server (http/grpc)
type ServerConfig struct {
	Port uint16 `json:"port"`
}
