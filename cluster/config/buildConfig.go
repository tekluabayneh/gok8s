package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"

	Types "github.com/tekluabayneh/gok8s/types"
)

func NewHTTPClient(cfg *Types.ApiKubeconfig) (*http.Client, error) {
	tlsCfg := &tls.Config{}

	if cfg.InsecureSkipTLSVerify {
		tlsCfg.InsecureSkipVerify = true
	} else if cfg.CertificateAuthorityData != "" {
		ca, err := base64.StdEncoding.DecodeString(cfg.CertificateAuthorityData)
		if err != nil {
			return nil, fmt.Errorf("decoding CA data: %w", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(ca) {
			return nil, fmt.Errorf("no valid certs found in CA data")
		}
		tlsCfg.RootCAs = pool
	}

	if cfg.ClientCertificateData != "" && cfg.ClientKeyData != "" {
		certPEM, err := base64.StdEncoding.DecodeString(cfg.ClientCertificateData)
		if err != nil {
			return nil, fmt.Errorf("decoding client cert: %w", err)
		}
		keyPEM, err := base64.StdEncoding.DecodeString(cfg.ClientKeyData)
		if err != nil {
			return nil, fmt.Errorf("decoding client key: %w", err)
		}
		pair, err := tls.X509KeyPair(certPEM, keyPEM)
		if err != nil {
			return nil, fmt.Errorf("loading client key pair: %w", err)
		}
		tlsCfg.Certificates = []tls.Certificate{pair}
	}

	// 3. Build the client
	return &http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsCfg},
	}, nil
}
