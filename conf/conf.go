package conf

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"os/exec"

	"github.com/p4gefau1t/trojan-go/common"
	utls "github.com/refraction-networking/utls"
)

type RunType string

const (
	Client  RunType = "client"
	Server  RunType = "server"
	NAT     RunType = "nat"
	Forward RunType = "forward"
	Relay   RunType = "relay"
)

type DNSType string

const (
	UDP DNSType = "udp"
	DOH DNSType = "https"
	DOT DNSType = "dot"
	TCP DNSType = "tcp"
)

type TLSConfig struct {
	Verify               bool     `json:"verify"`
	VerifyHostName       bool     `json:"verify_hostname"`
	CertPath             string   `json:"cert"`
	KeyPath              string   `json:"key"`
	ClientCertPath       []string `json:"client_cert"`
	KeyPassword          string   `json:"key_password"`
	Cipher               string   `json:"cipher"`
	CipherTLS13          string   `json:"cipher_tls13"`
	PreferServerCipher   bool     `json:"prefer_server_cipher"`
	SNI                  string   `json:"sni"`
	HTTPResponseFileName string   `json:"plain_http_response"`
	FallbackHost         string   `json:"fallback_addr"`
	FallbackPort         int      `json:"fallback_port"`
	ReuseSession         bool     `json:"reuse_session"`
	ALPN                 []string `json:"alpn"`
	Curves               string   `json:"curves"`
	Fingerprint          string   `json:"fingerprint"`
	KeyLogPath           string   `json:"key_log"`

	ClientHelloID    *utls.ClientHelloID
	FallbackAddress  *common.Address
	CertPool         *x509.CertPool
	ClientCertPool   *x509.CertPool
	KeyPair          []tls.Certificate
	HTTPResponse     []byte
	CipherSuites     []uint16
	CipherSuiteTLS13 []uint16
	SessionTicket    bool
	CurvePreferences []tls.CurveID
	KeyLogger        io.Writer
}

type TCPConfig struct {
	PreferIPV4   bool `json:"prefer_ipv4"`
	KeepAlive    bool `json:"keep_alive"`
	FastOpen     bool `json:"fast_open"`
	FastOpenQLen int  `json:"fast_open_qlen"`
	ReusePort    bool `json:"reuse_port"`
	NoDelay      bool `json:"no_delay"`
}

type MuxConfig struct {
	Enabled     bool `json:"enabled"`
	IdleTimeout int  `json:"idle_timeout"`
	Concurrency int  `json:"concurrency"`
}

type MySQLConfig struct {
	Enabled    bool   `json:"enabled"`
	ServerHost string `json:"server_addr"`
	ServerPort int    `json:"server_port"`
	Database   string `json:"database"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CheckRate  int    `json:"check_rate"`
}

type RedisConfig struct {
	Enabled    bool   `json:"enabled"`
	ServerHost string `json:"server_addr"`
	ServerPort int    `json:"server_port"`
	Password   string `json:"password"`
}

type ForwardProxyConfig struct {
	Enabled   bool   `json:"enabled"`
	ProxyHost string `json:"proxy_addr"`
	ProxyPort int    `json:"proxy_port"`
	Username  string `json:"username"`
	Password  string `json:"password"`

	ProxyAddress *common.Address
}

type CompressionConfig struct {
	Enabled bool `json:"enabled"`
}

type RouterConfig struct {
	Enabled         bool     `json:"enabled"`
	Bypass          []string `json:"bypass"`
	Proxy           []string `json:"proxy"`
	Block           []string `json:"block"`
	DomainStrategy  string   `json:"domain_strategy"`
	DefaultPolicy   string   `json:"default_policy"`
	GeoIPFilename   string   `json:"geoip"`
	GeoSiteFilename string   `json:"geosite"`

	BypassList []byte
	ProxyList  []byte
	BlockList  []byte

	GeoIP          []byte
	BypassIPCode   []string
	ProxyIPCode    []string
	BlockIPCode    []string
	GeoSite        []byte
	BypassSiteCode []string
	ProxySiteCode  []string
	BlockSiteCode  []string
}

type WebsocketConfig struct {
	Enabled             bool      `json:"enabled"`
	HostName            string    `json:"hostname"`
	Path                string    `json:"path"`
	ObfuscationPassword string    `json:"obfuscation_password"`
	DoubleTLS           bool      `json:"double_tls"`
	TLS                 TLSConfig `json:"ssl"`

	ObfuscationKey []byte
}

type APIConfig struct {
	Enabled bool      `json:"enabled"`
	APIHost string    `json:"api_addr"`
	APIPort int       `json:"api_port"`
	APITLS  bool      `json:"api_tls"`
	TLS     TLSConfig `json:"ssl"`

	APIAddress *common.Address
}

type TransportPluginConfig struct {
	Enabled      bool     `json:"enabled"`
	Type         string   `json:"type"`
	Command      string   `json:"command"`
	PluginOption string   `json:"plugin_option"`
	Arg          []string `json:"arg"`
	Env          []string `json:"env"`

	Cmd *exec.Cmd
}

type GlobalConfig struct {
	RunType          RunType               `json:"run_type"`
	LogLevel         int                   `json:"log_level"`
	LogFile          string                `json:"log_file"`
	LocalHost        string                `json:"local_addr"`
	LocalPort        int                   `json:"local_port"`
	TargetHost       string                `json:"target_addr"`
	TargetPort       int                   `json:"target_port"`
	RemoteHost       string                `json:"remote_addr"`
	RemotePort       int                   `json:"remote_port"`
	BufferSize       int                   `json:"buffer_size"`
	DisableHTTPCheck bool                  `json:"disable_http_check"`
	Passwords        []string              `json:"password"`
	DNS              []string              `json:"dns"`
	TLS              TLSConfig             `json:"ssl"`
	TCP              TCPConfig             `json:"tcp"`
	MySQL            MySQLConfig           `json:"mysql"`
	Redis            RedisConfig           `json:"redis"`
	Mux              MuxConfig             `json:"mux"`
	Router           RouterConfig          `json:"router"`
	Websocket        WebsocketConfig       `json:"websocket"`
	API              APIConfig             `json:"api"`
	ForwardProxy     ForwardProxyConfig    `json:"forward_proxy"`
	Compression      CompressionConfig     `json:"compression"`
	TransportPlugin  TransportPluginConfig `json:"transport_plugin"`

	LocalAddress  *common.Address
	RemoteAddress *common.Address
	TargetAddress *common.Address
	Hash          map[string]string
}
