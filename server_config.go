package gold

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// ServerConfig holds a list of configuration parameters for the server
type ServerConfig struct {
	// PortHTTP contains the HTTPS listening port number in format ":80"
	ListenHTTP string

	// PortHTTPS contains the HTTPS listening port number in format ":443"
	ListenHTTPS string

	// TLSCert holds the server certificate eg. cert.pem
	TLSCert string

	// TLSKey holds the server key eg. key.pem
	TLSKey string

	// Root points to the folder that will be used as root for data
	DataRoot string

	// Vhosts enables the use of virtual hosts (i.e. user.example.org)
	Vhosts bool

	// Insecure enables insecure (HTTP) operation mode only
	Insecure bool

	// NoHTTP allows to enable or disable redirects from HTTP to HTTPS
	NoHTTP bool

	// Debug (display or hide stdout logging)
	Debug bool

	// CookieAge contains the validity duration for cookies (in hours)
	CookieAge int64

	// TokenAge contains the validity duration for recovery tokens (in minutes)
	TokenAge int64

	// TokenAge contains the validity duration for recovery tokens (in minutes)
	AuthProtocol string

	// METASuffix sets the default suffix for meta files (e.g. ,meta or .meta)
	MetaSuffix string

	// ACLSuffix sets the default suffix for ACL files (e.g. ,acl or .acl)
	ACLSuffix string

	// DataApp sets the default app for viewing RDF resources
	DataApp string

	// DirApp points to the app for browsing the data space
	DirApp string

	// SignUpApp points to the app used for creating new accounts
	SignUpApp string

	// DirIndex contains the default index file name
	DirIndex []string

	// DiskLimit is the maximum total disk (in bytes) to be allocated to a given user
	DiskLimit int

	// SMTPConfig holds the settings for the remote SMTP user/server
	SMTPConfig EmailConfig
}

// NewServerConfig creates a new config object
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		CookieAge:    8736, // hours (1 year)
		TokenAge:     5,
		AuthProtocol: "webid",
		MetaSuffix:   ".meta",
		ACLSuffix:    ".acl",
		DataApp:      "tabulator",
		DirIndex:     []string{"index.html", "index.htm"},
		DirApp:       "http://linkeddata.github.io/warp/#list/",
		SignUpApp:    "https://solid.github.io/solid-signup/?domain=",
		DiskLimit:    100000000, // 100MB
		DataRoot:     serverDefaultRoot(),
	}
}

// LoadJSONFile loads server configuration
func (c *ServerConfig) LoadJSONFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &c)
}

func serverDefaultRoot() string {
	serverRoot, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if !strings.HasSuffix(serverRoot, "/") {
		serverRoot += "/"
	}
	return serverRoot
}
