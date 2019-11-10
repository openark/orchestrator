package db

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/github/orchestrator/go/config"
)

func init() {
	config.Config.MySQLOrchestratorUseMutualTLS = false
}

func TestSetupMySQLTopologyTLS_returnsUri(t *testing.T) {
	if new_uri, err := SetupMySQLTopologyTLS(createUri(t, "topology_returns_uri")); err != nil {
		t.Fatal(err)
	} else if !strings.Contains(new_uri, "&tls=topology") {
		t.Fatalf("New URI does not contain tls configuration.")
	}
}

func TestSetupMySQLOrchestratorTLS_returnsUri(t *testing.T) {
	if new_uri, err := SetupMySQLOrchestratorTLS(createUri(t, "orchestrator_returns_uri")); err != nil {
		t.Fatal(err)
	} else if !strings.Contains(new_uri, "&tls=orchestrator") {
		t.Fatalf("New URI does not contain tls configuration.")
	}
}

func TestSetupMySQLOrchestratorTLS_withoutMutualTLS(t *testing.T) {
	original_uri := createUri(t, "orchestrator_without_mutual_tls")
	if _, err := SetupMySQLOrchestratorTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); len(tlsConfig.Certificates) != 0 {
		t.Fatalf("A client key pair should not have been added.")
	}
}

func TestSetupMySQLTopologyTLS_withoutMutualTLS(t *testing.T) {
	original_uri := createUri(t, "topology_without_mutual_tls")
	if _, err := SetupMySQLTopologyTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); len(tlsConfig.Certificates) != 0 {
		t.Fatalf("A client key pair should not have been added.")
	}
}

func TestSetupMySQLOrchestratorTLS_withMutualTLS(t *testing.T) {
	config.Config.MySQLOrchestratorUseMutualTLS = true
	config.Config.MySQLOrchestratorSSLCertFile = writeFakeFile(t, pemCertificate)
	config.Config.MySQLOrchestratorSSLPrivateKeyFile = writeFakeFile(t, pemPrivateKey)

	original_uri := createUri(t, "orchestrator_with_mutual_tls")
	if _, err := SetupMySQLOrchestratorTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); len(tlsConfig.Certificates) != 1 {
		t.Fatalf("A client key pair should have been added but was not.")
	}
}

func TestSetupMySQLTopologyTLS_withMutualTLS(t *testing.T) {
	config.Config.MySQLTopologyUseMutualTLS = true
	config.Config.MySQLTopologySSLCertFile = writeFakeFile(t, pemCertificate)
	config.Config.MySQLTopologySSLPrivateKeyFile = writeFakeFile(t, pemPrivateKey)

	original_uri := createUri(t, "topology_with_mutual_tls")
	if _, err := SetupMySQLTopologyTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); len(tlsConfig.Certificates) != 1 {
		t.Fatalf("A client key pair should have been added but was not.")
	}
}

func TestSetupMySQLOrchestratorTLS_skipVerify(t *testing.T) {
	config.Config.MySQLOrchestratorSSLSkipVerify = true

	original_uri := createUri(t, "orchestrator_with_skip_verify")
	if _, err := SetupMySQLOrchestratorTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); !tlsConfig.InsecureSkipVerify {
		t.Fatalf("Certificate verification should have been disabled.")
	}
}

func TestSetupMySQLTopologyTLS_skipVerify(t *testing.T) {
	config.Config.MySQLTopologySSLSkipVerify = true

	original_uri := createUri(t, "topology_with_skip_verify")
	if _, err := SetupMySQLTopologyTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); !tlsConfig.InsecureSkipVerify {
		t.Fatalf("Certificate verification should have been disabled.")
	}
}

func TestSetupMySQLOrchestratorTLS_allowTLS10(t *testing.T) {
	original_uri := createUri(t, "orchestrator_with_tls_10")
	if _, err := SetupMySQLOrchestratorTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); tlsConfig.MinVersion != tls.VersionTLS10 {
		t.Fatalf("Minimum TLS version should be set to TLSv1.0")
	}
}

func TestSetupMySQLTopologyTLS_allowTLS10(t *testing.T) {
	original_uri := createUri(t, "topology_with_tls_10")
	if _, err := SetupMySQLTopologyTLS(original_uri); err != nil {
		t.Fatal(err)
	}

	if tlsConfig := getTlsConfig(t, original_uri); tlsConfig.MinVersion != tls.VersionTLS10 {
		t.Fatalf("Minimum TLS version should be set to TLSv1.0")
	}
}

func getTlsConfig(t *testing.T, database string) *tls.Config {
	t.Helper()

	tlsConfig, found := tlsConfigCache.Get(database)
	if !found {
		t.Fatalf("A TLS configuration for %s should have been added to the TLS cache but was not.", database)
	}
	return tlsConfig.(*tls.Config)
}

func createUri(t *testing.T, database string) string {
	t.Helper()
	uri := fmt.Sprintf("user:password@tcp(localhost:5554)/%s?timeout=10", database)
	if _, found := tlsConfigCache.Get(uri); found {
		t.Fatalf("This TLS configuration for %s has already been created. The calling test probably forgot to use a unique URI.", database)
	}
	return uri
}

func writeFakeFile(t *testing.T, content string) string {
	t.Helper()

	f, err := ioutil.TempFile("", "tls_test")
	if err != nil {
		return ""
	}
	ioutil.WriteFile(f.Name(), []byte(content), 0644)
	return f.Name()
}

const pemCertificate = `-----BEGIN CERTIFICATE-----
MIIDtTCCAp2gAwIBAgIJAOxKC7FsJelrMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTcwODEwMTQ0MjM3WhcNMTgwODEwMTQ0MjM3WjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEA12vHV3gYy5zd1lujA7prEhCSkAszE6E37mViWhLQ63CuedZfyYaTAHQK
HYDZi4K1MNAySUfZRMcICSSsxlRIz6mzXrFsowaJgwx4cbMDIvXE03KstuXoTYJh
+xmXB+5yEVEtIyP2DvPqfCmwCZb3k94Y/VY1nAQDxIxciXrAxT9zT1oYd0YWr2yp
J2mgsfnY4c3zg7W5WgvOTmYz7Ey7GJjpUjGdayx+P1CilKzSWH1xZuVQFNLSHvcH
WXkEoCMVc0tW5mO5eEO1aNHo9MSjPF386l1rq+pz5OwjqCEZq2b1YxesyLnbF+8+
iYGfYmFaDLFwG7zVDwialuI4TzIIOQIDAQABo4GnMIGkMB0GA1UdDgQWBBQ1ubGx
Yvn3wN5VXyoR0lOD7ARzVTB1BgNVHSMEbjBsgBQ1ubGxYvn3wN5VXyoR0lOD7ARz
VaFJpEcwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgTClNvbWUtU3RhdGUxITAfBgNV
BAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZIIJAOxKC7FsJelrMAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQEFBQADggEBALmm4Zw/4jLKDJciUGUYOcr5Xe9TP/Cs
afH7IWvaFUDfV3W6yAm9jgNfIy9aDLpuu2CdEb+0qL2hdmGLV7IM3y62Ve0UTdGV
BGsm1zMmIguew2wGbAwGr5LmIcUseatVUKAAAfDrBNwotEAdM8kmGekUZfOM+J9D
FoNQ62C0buRHGugtu6zWAcZNOe6CI7HdhaAdxZlgn8y7dfJQMacoK0NcWeUVQwii
6D4mgaqUGM2O+WcquD1vEMuBPYVcKhi43019E0+6LI5QB6w80bARY8K7tkTdRD7U
y1/C7iIqyuBVL45OdSabb37TfGlHZIPIwLaGw3i4Mr0+F0jQT8rZtTQ=
-----END CERTIFICATE-----`

const pemPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA12vHV3gYy5zd1lujA7prEhCSkAszE6E37mViWhLQ63CuedZf
yYaTAHQKHYDZi4K1MNAySUfZRMcICSSsxlRIz6mzXrFsowaJgwx4cbMDIvXE03Ks
tuXoTYJh+xmXB+5yEVEtIyP2DvPqfCmwCZb3k94Y/VY1nAQDxIxciXrAxT9zT1oY
d0YWr2ypJ2mgsfnY4c3zg7W5WgvOTmYz7Ey7GJjpUjGdayx+P1CilKzSWH1xZuVQ
FNLSHvcHWXkEoCMVc0tW5mO5eEO1aNHo9MSjPF386l1rq+pz5OwjqCEZq2b1Yxes
yLnbF+8+iYGfYmFaDLFwG7zVDwialuI4TzIIOQIDAQABAoIBAHLf4pleTbqmmBWr
IC7oxhgIBmAR2Nbq7eyO2/e0ePxURnZqPwI0ZUekmZBKGbgvp3e0TlyNl+r5R+u4
RvosD/fNQv2IF6qH3eSoTcIz98Q40xD+4eNWjp5mnOFOMB/mo6VgaHWIw7oNkElN
4bX7b2LG2QSfaE8eRPQW9XHKp+mGhYFbxgPYxUmlIXuYZF61hVwxysDA6DP3LOi8
yUL6E64x6NqN9xtg/VoN+f6N0MOvsr4yb5+uvni1LVRFI7tNqIN4Y6P6trgKfnRR
EpZeAUu8scqyxE4NeqnnjK/wBuXxaeh3e9mN1V2SzT629c1InmmQasZ5slcCJQB+
38cswgECgYEA+esaLKwHXT4+sOqMYemi7TrhxtNC2f5OAGUiSRVmTnum2gl4wOB+
h5oLZAuG5nBEIoqbMEbI35vfuHqIe390IJtPdQlz4TGDsPufYj/gnnBBFy/c8f+n
f/CdRDRYrpnpKGwvUntLRB2pFbe2hlqqq+4YUqiHauJMOCJnPbOo1lECgYEA3KnF
VOXyY0fKD45G7ttfAcpw8ZI2gY99sCRwtBQGsbO61bvw5sl/3j7AmYosz+n6f7hb
uHmitIuPv4z3r1yfVysh80tTGIM3wDkpr3fLYRxpVOZU4hgxMQV9yyaSA/Hfqn48
vIK/NC4bERqpofNNdrIqNaGWkd87ZycvpRfa0WkCgYBztbVVr4RtWG9gLAg5IRot
KhD0pEWUdpiYuDpqifznI3r6Al6lNot+rwTNGkUoFhyFvZTigjNozFuFpz3fqAAV
RLNCJdFAF1O4spd1vst5r9GDMcbjSJG9u6KkvHO+y0XXUFeMoccUT4NEqd1ZUUsp
9T/PrXWdOA9AAjW4rKDkMQKBgQC9R4NVR8mbD8Frhoeh69qbFqO7E8hdalBN/3QN
hAAZ/imNnSEPVliwsvNSwQufbPzLAcDrhKrkY7JyhOERM0oa44zDvSESLbxszpvL
P97c9hoEEW9OYaIQgr1cvUES0S8ieBZxPVX11HazPUO0/5a68ijyyCD4D5xM53gf
DU9NwQKBgQCmVthQi65xcc4mgCIwXtBZWXeaPv5x0dLEXIC5EoN6eXLK9iW//7cE
hhawtJtl+J6laB+TkEGQsyhc4v85WcywdisyR7LR7CUqFYJMKeE/VtTVKnYbfq54
rHoQS9YotByBwPtRx0V93gkc+KWBOGmSBBxKj7lrBkYkcWAiRfpJjg==
-----END RSA PRIVATE KEY-----`
