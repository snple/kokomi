package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigStruct struct {
	Debug       bool        `toml:"debug"`
	DeviceID    string      `toml:"device_id"`
	Secret      string      `toml:"secret"`
	DB          DB          `toml:"db"`
	NodeClient  GRPCClient  `toml:"node"`
	QuicClient  QuicClient  `toml:"quic"`
	EdgeService GRPCService `toml:"edge"`
	SlotService GRPCService `toml:"slot"`
	Sync        Sync        `toml:"sync"`
}

type DB struct {
	Debug bool   `toml:"debug"`
	File  string `toml:"file"`
}

type GRPCClient struct {
	Addr               string `toml:"addr"`
	TLS                bool   `toml:"tls"`
	CA                 string `toml:"ca"`
	Cert               string `toml:"cert"`
	Key                string `toml:"key"`
	ServerName         string `toml:"server_name"`
	InsecureSkipVerify bool   `toml:"insecure_skip_verify"`
}

type QuicClient struct {
	Enable             bool   `toml:"enable"`
	Addr               string `toml:"addr"`
	CA                 string `toml:"ca"`
	Cert               string `toml:"cert"`
	Key                string `toml:"key"`
	ServerName         string `toml:"server_name"`
	InsecureSkipVerify bool   `toml:"insecure_skip_verify"`
}

type QuicService struct {
	Enable bool   `toml:"enable"`
	Addr   string `toml:"addr"`
	CA     string `toml:"ca"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type GRPCService struct {
	Enable bool   `toml:"enable"`
	Addr   string `toml:"addr"`
	TLS    bool   `toml:"tls"`
	CA     string `toml:"ca"`
	Cert   string `toml:"cert"`
	Key    string `toml:"key"`
}

type Sync struct {
	LinkStatusTTL  int  `toml:"link_status_ttl"`
	TokenRefresh   int  `toml:"token_refresh"`
	SyncLinkStatus int  `toml:"sync_link_status"`
	SyncInterval   int  `toml:"sync_interval"`
	SyncRealtime   bool `toml:"sync_realtime"`
}

func DefaultConfig() ConfigStruct {
	return ConfigStruct{
		Debug: false,
		DB: DB{
			File: "store.db",
		},
		NodeClient: GRPCClient{
			Addr: "127.0.0.1:6007",
			TLS:  true,
			CA:   "certs/ca.crt",
			Cert: "certs/client.crt",
			Key:  "certs/client.key",
		},
		QuicClient: QuicClient{
			Addr: "127.0.0.1:6008",
			CA:   "certs/ca.crt",
			Cert: "certs/client.crt",
			Key:  "certs/client.key",
		},
		EdgeService: GRPCService{
			Addr: "127.0.0.1:6010",
			TLS:  true,
			CA:   "certs/ca.crt",
			Cert: "certs/server.crt",
			Key:  "certs/server.key",
		},
		Sync: Sync{
			LinkStatusTTL:  60 * 3,
			TokenRefresh:   60 * 30,
			SyncLinkStatus: 60,
			SyncInterval:   60,
			SyncRealtime:   false,
		},
	}
}

var Config = DefaultConfig()

func (c *ConfigStruct) Validate() error {
	if len(c.DeviceID) == 0 {
		return errors.New("DeviceID must be specified")
	}

	if len(c.Secret) == 0 {
		return errors.New("Secret must be specified")
	}

	{
		if len(c.NodeClient.Addr) == 0 {
			return errors.New("NodeClient.Addr must be specified")
		}

		if c.NodeClient.TLS {
			if len(c.NodeClient.CA) == 0 {
				return errors.New("NodeClient.CA must be specified")
			}

			if len(c.NodeClient.Cert) == 0 {
				return errors.New("NodeClient.Cert must be specified")
			}

			if len(c.NodeClient.Key) == 0 {
				return errors.New("NodeClient.Key must be specified")
			}
		}
	}

	{
		if len(c.QuicClient.Addr) == 0 {
			return errors.New("QuicClient.Addr must be specified")
		}

		if len(c.QuicClient.CA) == 0 {
			return errors.New("QuicClient.CA must be specified")
		}

		if len(c.QuicClient.Cert) == 0 {
			return errors.New("QuicClient.Cert must be specified")
		}

		if len(c.QuicClient.Key) == 0 {
			return errors.New("QuicClient.Key must be specified")
		}
	}

	if c.EdgeService.Enable {
		if len(c.EdgeService.Addr) == 0 {
			return errors.New("EdgeService.Addr must be specified")
		}

		if c.EdgeService.TLS {
			if len(c.EdgeService.CA) == 0 {
				return errors.New("EdgeService.CA must be specified")
			}

			if len(c.EdgeService.Cert) == 0 {
				return errors.New("EdgeService.Cert must be specified")
			}

			if len(c.EdgeService.Key) == 0 {
				return errors.New("EdgeService.Key must be specified")
			}
		}
	}

	return nil
}

func Parse() {
	var err error

	configFile := flag.String("c", "config.toml", "config file")
	flag.Parse()

	if _, err = toml.DecodeFile(*configFile, &Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = Config.Validate(); err != nil {
		fmt.Println("config:", err)
		os.Exit(1)
	}
}
