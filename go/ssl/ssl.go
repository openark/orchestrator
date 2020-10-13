package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	nethttp "net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/howeyc/gopass"
	"github.com/openark/golib/log"
	"github.com/openark/orchestrator/go/config"
)

type IdentifierType int

const (
	IT_CommonName IdentifierType = iota
	IT_OrganizationalUnit
)

var cipherSuites = []uint16{
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	tls.TLS_RSA_WITH_AES_128_CBC_SHA,
	tls.TLS_RSA_WITH_AES_256_CBC_SHA,
}

// Determine if a string element is in a string array
func HasString(elem string, arr []string) bool {
	for _, s := range arr {
		if s == elem {
			return true
		}
	}
	return false
}

// NewTLSConfig returns an initialized TLS configuration suitable for client
// authentication. If caFile is non-empty, it will be loaded.
func NewTLSConfig(caFile string, verifyCert bool) (*tls.Config, error) {
	var c tls.Config

	// Set to TLS 1.2 as a minimum.  This is overridden for mysql communication
	c.MinVersion = tls.VersionTLS12
	// Remove insecure ciphers from the list
	c.CipherSuites = cipherSuites
	c.PreferServerCipherSuites = true

	if verifyCert {
		log.Info("verifyCert requested, client certificates will be verified")
		c.ClientAuth = tls.VerifyClientCertIfGiven
	}
	caPool, err := ReadCAFile(caFile)
	if err != nil {
		return &c, err
	}
	c.ClientCAs = caPool
	c.BuildNameToCertificate()
	return &c, nil
}

// Returns CA certificate. If caFile is non-empty, it will be loaded.
func ReadCAFile(caFile string) (*x509.CertPool, error) {
	var caCertPool *x509.CertPool
	if caFile != "" {
		data, err := ioutil.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		caCertPool = x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(data) {
			return nil, errors.New("No certificates parsed")
		}
		log.Info("Read in CA file:", caFile)
	}
	return caCertPool, nil
}

// Verify that the OU of the presented client certificate matches the list
// of Valid OUs
func Verify(r *nethttp.Request, validIdentitifiers []string, it IdentifierType) error {
	if config.Config.SSLSkipVerify {
		log.Debug("TLS Verify: skipping verification")
		return nil
	}
	if HasString(r.URL.String(), config.Config.SSLSkipVerifyEndpoints) {
		log.Debug("TLS Verify: non-verifiable endpoint:", r.URL.String())
		return nil
	}
	if r.TLS == nil {
		return errors.New("TLS Verify: No TLS")
	}
	for _, chain := range r.TLS.VerifiedChains {
		s := []string{}
		if it == IT_CommonName {
			s = append(s, chain[0].Subject.CommonName)
		} else {
			s = chain[0].Subject.OrganizationalUnit
		}
		log.Debug("TLS Verify: all identifiers:", strings.Join(s, " "))
		for _, id := range s {
			log.Debug("TLS Verify: Client presented identifier:", id)
			if HasString(id, validIdentitifiers) {
				log.Debug("TLS Verify: Found valid identifier:", id)
				return nil
			}
		}
	}
	log.Error("TLS Verify: No valid client identitifiers found")
	return errors.New("TLS Verify: Invalid client identity")
}

// TODO: make this testable?
func VerifyClient(validIdentifiers []string, it IdentifierType) martini.Handler {
	return func(res nethttp.ResponseWriter, req *nethttp.Request, c martini.Context) {
		log.Debug("TLS: Verifying client certificate")
		if err := Verify(req, validIdentifiers, it); err != nil {
			nethttp.Error(res, err.Error(), nethttp.StatusUnauthorized)
		}
	}
}

// AppendKeyPair loads the given TLS key pair and appends it to
// tlsConfig.Certificates.
func AppendKeyPair(tlsConfig *tls.Config, certFile string, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}
	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
	return nil
}

// Read in a keypair where the key is password protected
func AppendKeyPairWithPassword(tlsConfig *tls.Config, certFile string, keyFile string, pemPass []byte) error {

	// Certificates aren't usually password protected, but we're kicking the password
	// along just in case.  It won't be used if the file isn't encrypted
	certData, err := ReadPEMData(certFile, pemPass)
	if err != nil {
		return err
	}
	keyData, err := ReadPEMData(keyFile, pemPass)
	if err != nil {
		return err
	}
	cert, err := tls.X509KeyPair(certData, keyData)
	if err != nil {
		return err
	}
	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
	return nil
}

// Read a PEM file and ask for a password to decrypt it if needed
func ReadPEMData(pemFile string, pemPass []byte) ([]byte, error) {
	pemData, err := ioutil.ReadFile(pemFile)
	if err != nil {
		return pemData, err
	}

	// We should really just get the pem.Block back here, if there's other
	// junk on the end, warn about it.
	pemBlock, rest := pem.Decode(pemData)
	if len(rest) > 0 {
		log.Warning("Didn't parse all of", pemFile)
	}

	if x509.IsEncryptedPEMBlock(pemBlock) {
		// Decrypt and get the ASN.1 DER bytes here
		pemData, err = x509.DecryptPEMBlock(pemBlock, pemPass)
		if err != nil {
			return pemData, err
		} else {
			log.Info("Decrypted", pemFile, "successfully")
		}
		// Shove the decrypted DER bytes into a new pem Block with blank headers
		var newBlock pem.Block
		newBlock.Type = pemBlock.Type
		newBlock.Bytes = pemData
		// This is now like reading in an uncrypted key from a file and stuffing it
		// into a byte stream
		pemData = pem.EncodeToMemory(&newBlock)
	}
	return pemData, nil
}

// Print a password prompt on the terminal and collect a password
func GetPEMPassword(pemFile string) []byte {
	fmt.Printf("Password for %s: ", pemFile)
	pass, err := gopass.GetPasswd()
	if err != nil {
		// We'll error with an incorrect password at DecryptPEMBlock
		return []byte("")
	}
	return pass
}

// Determine if PEM file is encrypted
func IsEncryptedPEM(pemFile string) bool {
	pemData, err := ioutil.ReadFile(pemFile)
	if err != nil {
		return false
	}
	pemBlock, _ := pem.Decode(pemData)
	if len(pemBlock.Bytes) == 0 {
		return false
	}
	return x509.IsEncryptedPEMBlock(pemBlock)
}

// ListenAndServeTLS acts identically to http.ListenAndServeTLS, except that it
// expects TLS configuration.
// TODO: refactor so this is testable?
func ListenAndServeTLS(addr string, handler nethttp.Handler, tlsConfig *tls.Config) error {
	if addr == "" {
		// On unix Listen calls getaddrinfo to parse the port, so named ports are fine as long
		// as they exist in /etc/services
		addr = ":https"
	}
	l, err := tls.Listen("tcp", addr, tlsConfig)
	if err != nil {
		return err
	}
	return nethttp.Serve(l, handler)
}

func NewTLSTransport() (*nethttp.Transport, error) {
	var transport *nethttp.Transport
	tlsConfig, err := NewTLSConfig(config.Config.SSLCAFile, false)
	if err != nil {
		log.Errorf("Can't create TLS configuration for reverse proxy connection: %s", err)
		return transport, err
	}

	if config.Config.SSLCertFile != "" || config.Config.SSLPrivateKeyFile != "" {
		if err = AppendKeyPair(tlsConfig, config.Config.SSLCertFile, config.Config.SSLPrivateKeyFile); err != nil {
			log.Errorf("Can't setup TLS key pairs for reverse proxy connection: %s", err)
			return transport, err
		}
	}

	verifyPeerCn := func(certificates [][]byte, _ [][]*x509.Certificate) error {
		certs := make([]*x509.Certificate, len(certificates))
		for i, asn1Data := range certificates {
			cert, err := x509.ParseCertificate(asn1Data)
			if err != nil {
				log.Errorf("tls: failed to parse certificate from server: %v", err)
				return err
			}
			certs[i] = cert
		}

		opts := x509.VerifyOptions{
			Roots:         tlsConfig.ClientCAs,
			Intermediates: x509.NewCertPool(),
			KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		}
		for _, cert := range certs[1:] {
			opts.Intermediates.AddCert(cert)
		}

		var err error
		for _, cn := range config.Config.SSLValidCNs {
			opts.DNSName = cn
			_, err = certs[0].Verify(opts)
			if err == nil {
				log.Debugf("tls: accepted server certificate with CN %s", cn)
				return nil
			}
		}
		log.Debugf("tls: rejected server certificate")
		return err
	}

	tlsConfig.ClientAuth = tls.RequireAnyClientCert
	tlsConfig.VerifyPeerCertificate = verifyPeerCn
	tlsConfig.InsecureSkipVerify = true

	transport = &nethttp.Transport{
		DialTLS: func(network, addr string) (net.Conn, error) {
			conn, err := tls.Dial(network, addr, tlsConfig)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	return transport, nil
}
