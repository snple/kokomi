package util

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
)

func LoadClientCert(caFile, certFile, keyFile, serverName string, insecureSkipVerify bool) (*tls.Config, error) {
	pool := x509.NewCertPool()

	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	if ok := pool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("pool.AppendCertsFromPEM err")
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	cfg := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		ServerName:         serverName,
		RootCAs:            pool,
		InsecureSkipVerify: insecureSkipVerify,
	}

	return cfg, nil
}

func LoadServerCert(caFile, certFile, keyFile string) (*tls.Config, error) {
	pool := x509.NewCertPool()

	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	if ok := pool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("pool.AppendCertsFromPEM err")
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    pool,
	}

	return cfg, nil
}
