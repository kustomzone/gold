package main

import (
	"crypto/tls"
)

var (
	tlsConfig = &tls.Config{
		ClientAuth: tls.RequestClientCert,
		NextProtos: []string{"http/1.1"},
	}

	tlsTestCert = []byte(`-----BEGIN CERTIFICATE-----
MIIB4TCCAUygAwIBAgIBADALBgkqhkiG9w0BAQUwEjEQMA4GA1UEChMHQWNtZSBD
bzAeFw0xNDAxMzAyMzUyMTlaFw0yNDAxMjgyMzUyMTlaMBIxEDAOBgNVBAoTB0Fj
bWUgQ28wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAMs8NmXX55GqvTRcIE2K
8ZoElA7xRuiIYPXFl6Zlt/xCYUzcxEEz2pKOX3jgYEzx4wG0hQ5bcNQMJWPftZ7K
6QBvDRWs8wVgrbeN8o9LelPDrPl40Zk96howpgek/nPd5AUt6y0/hV4CNVt07y+D
13BxZSEj1E8ZTwCwhQ9uGltPAgMBAAGjSzBJMA4GA1UdDwEB/wQEAwIAoDATBgNV
HSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMBQGA1UdEQQNMAuCCWxvY2Fs
aG9zdDALBgkqhkiG9w0BAQUDgYEAawZEY85RZAKrROH3t1xuGLI+MIWmiFH5Z/aQ
3kA/v5YHLlygjbgxedgFEe9TodiMk9M7kUTmAM6vS2qYf+apAj2QHFFyR8xc/BZ2
YHpBjeARoeg1ctbzCWeISB4BN7hOAQOojKcgaqbP49S5WG+ONfF6GuRE3oBJPJZf
1bRSET8=
-----END CERTIFICATE-----`)
	tlsTestKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDLPDZl1+eRqr00XCBNivGaBJQO8UboiGD1xZemZbf8QmFM3MRB
M9qSjl944GBM8eMBtIUOW3DUDCVj37WeyukAbw0VrPMFYK23jfKPS3pTw6z5eNGZ
PeoaMKYHpP5z3eQFLestP4VeAjVbdO8vg9dwcWUhI9RPGU8AsIUPbhpbTwIDAQAB
AoGAc00U25CzCvxf3V3K4dNLIIMqcJPIE9KTl7vjPn8E87PBOfchzJAbl/v4BD7f
w6eTj3sX5b5Q86x0ZgYcJxudNiLJK8XrrYqpe9yMoQ4PsN2mL77VtxwiiDrINnW+
eWX5eavIXFd1d6cNbudPy/vS4MpOAMid/g/m53tH8V/ZPUkCQQD7DGcW5ra05dK7
qpcj+TRQACe2VSgo78Li9DoifoU9vdx3pWWNxthdGfUlMXuAyl29sFXsxVE/ve47
k7jf/YSTAkEAzz5j+F28XwRkC+2HEQFTk+CBDsV3iNFNcRFeQiaYXwI6OCmQQXDA
pdmcjFqUzcKh7Wtx3G/Fz8hyifzr4/Xf1QJBAJgSjEP4H8b2zK93h7R32bN4VJYD
gZ9ClYhLLwgEIgwjfXBQlXLLd/b1qWUNU2XRr/Ue4v3ZDP2SvMQEGOI+PNcCQQCF
j3PmEKLhqXbAqSeusegnGTyTRHew2RLLl6Hjh/QS5uCWaVLqmbvOJtxZJ9dWc+Tf
masboX0eV9RZUYLEuySxAkBLfEizykRCZ1CYkIUtKsq6HOtj+ELPBVtVPMCx3O10
LMEOXuCrAMT/nApK629bgSlTU6P9PZd+05yRbHt4Ds1S
-----END RSA PRIVATE KEY-----`)
)
