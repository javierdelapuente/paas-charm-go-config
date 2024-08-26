package examples_test

import (
	"fmt"
	"os"
	"testing"

	"reflect"

	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/google/go-cmp/cmp"
	"github.com/javierdelapuente/paascharmgogenerator/examples"
)

func TestConfigVariableLoading(t *testing.T) {
	var tests = []struct {
		inputEnvVars map[string]string
		expected     examples.CharmConfig
	}{
		// Empty case. Should not be that way
		{map[string]string{}, examples.CharmConfig{}},
		// Simpler (real) example
		{
			map[string]string{
				"APP_BASE_URL":     "http://go-k8s.testing:8080",
				"APP_METRICS_PATH": "/metrics",
				"APP_METRICS_PORT": "8080",
				"APP_PORT":         "8080",
				"APP_SECRET_KEY":   "_6vLis5ukuN8p03AMQClOdhDPBsJ2YvjbLK_-9l8BVHmttsAzXPZiUcHKhRLb8-KUzpDReFEO8yhwzIVQ0OA0Q",
				"NO_PROXY":         "127.0.0.1,localhost,::1",
				"no_proxy":         "127.0.0.1,localhost,::1",
			},
			func() examples.CharmConfig {
				metricsPort := 8080
				metricsPath := "/metrics"
				secretKey := "_6vLis5ukuN8p03AMQClOdhDPBsJ2YvjbLK_-9l8BVHmttsAzXPZiUcHKhRLb8-KUzpDReFEO8yhwzIVQ0OA0Q"
				return examples.CharmConfig{
					ConfigOptions: examples.ConfigOptions{
						Port:        8080,
						MetricsPort: &metricsPort,
						MetricsPath: &metricsPath,
						SecretKey:   &secretKey,
					},
					ProxyConfig: examples.ProxyConfig{
						NoProxy: []string{"127.0.0.1", "localhost", "::1"},
					},
				}
			}(),
		},
		// Big example
		{
			map[string]string{
				"APP_BASE_URL":                     "http://go-k8s.testing:8080",
				"APP_METRICS_PATH":                 "/metrics",
				"APP_METRICS_PORT":                 "8080",
				"APP_PORT":                         "8080",
				"APP_POSTGRESQL_DB_CONNECT_STRING": "postgresql://relation_id_4:MTqX51qDce6o4fpb@postgresql-k8s-primary.testing.svc.cluster.local:5432/go-k8s",
				"APP_POSTGRESQL_DB_FRAGMENT":       "",
				"APP_POSTGRESQL_DB_HOSTNAME":       "postgresql-k8s-primary.testing.svc.cluster.local",
				"APP_POSTGRESQL_DB_NAME":           "go-k8s",
				"APP_POSTGRESQL_DB_NETLOC":         "relation_id_4:MTqX51qDce6o4fpb@postgresql-k8s-primary.testing.svc.cluster.local:5432",
				"APP_POSTGRESQL_DB_PARAMS":         "",
				"APP_POSTGRESQL_DB_PASSWORD":       "MTqX51qDce6o4fpb",
				"APP_POSTGRESQL_DB_PATH":           "/go-k8s",
				"APP_POSTGRESQL_DB_PORT":           "5432",
				"APP_POSTGRESQL_DB_QUERY":          "",
				"APP_POSTGRESQL_DB_SCHEME":         "postgresql",
				"APP_POSTGRESQL_DB_USERNAME":       "relation_id_4",
				"APP_MONGODB_DB_CONNECT_STRING":    "mongodb://mongodb.example.com/",
				"APP_SECRET_KEY":                   "_6vLis5ukuN8p03AMQClOdhDPBsJ2YvjbLK_-9l8BVHmttsAzXPZiUcHKhRLb8-KUzpDReFEO8yhwzIVQ0OA0Q",
				"APP_USER-DEFINED-CONFIG":          "newvalue",
				"APP_USER_DEFINED_CONFIG":          "newvalue",
				"NO_PROXY":                         "127.0.0.1,localhost,::1",
				"no_proxy":                         "127.0.0.1,localhost,::1",
			},
			func() examples.CharmConfig {
				metricsPort := 8080
				metricsPath := "/metrics"
				secretKey := "_6vLis5ukuN8p03AMQClOdhDPBsJ2YvjbLK_-9l8BVHmttsAzXPZiUcHKhRLb8-KUzpDReFEO8yhwzIVQ0OA0Q"
				postgreSQLUsername := "relation_id_4"
				postgreSQLHostname := "postgresql-k8s-primary.testing.svc.cluster.local"
				postgreSQLPort := 5432
				return examples.CharmConfig{
					ConfigOptions: examples.ConfigOptions{
						Port:        8080,
						MetricsPort: &metricsPort,
						MetricsPath: &metricsPath,
						SecretKey:   &secretKey,
					},
					ProxyConfig: examples.ProxyConfig{
						NoProxy: []string{"127.0.0.1", "localhost", "::1"},
					},
					Integrations: examples.Integrations{
						PostgreSQL: examples.PostgreSQLIntegration{
							examples.DatabaseIntegration{
								ConnectString: "postgresql://relation_id_4:MTqX51qDce6o4fpb@postgresql-k8s-primary.testing.svc.cluster.local:5432/go-k8s",
								Scheme:        "postgresql",
								NetLoc:        "relation_id_4:MTqX51qDce6o4fpb@postgresql-k8s-primary.testing.svc.cluster.local:5432",
								Path:          "/go-k8s",
								Username:      &postgreSQLUsername,
								Hostname:      &postgreSQLHostname,
								Port:          &postgreSQLPort,
							},
						},
						MongoDB: examples.MongoDBIntegration{
							examples.DatabaseIntegration{ConnectString: "mongodb://mongodb.example.com/"},
						},
					},
				}
			}(),
		},
	}
	previousEnvVars := os.Environ()
	defer func() {
		os.Clearenv()
		for _, keyVal := range previousEnvVars {
			split := strings.SplitN(keyVal, "=", 2)
			key, val := split[0], split[1]
			os.Setenv(key, val)
		}
	}()

	for ix, tt := range tests {
		testname := fmt.Sprintf("%d", ix)
		t.Run(testname, func(t *testing.T) {
			os.Clearenv()
			for key, val := range tt.inputEnvVars {
				os.Setenv(key, val)
			}

			var cfg examples.CharmConfig
			err := env.Parse(&cfg)
			if err != nil {
				t.Errorf("Error parsing config %s", err)
			}
			// Another option is to use https://pkg.go.dev/github.com/google/go-cmp/cmp#example-Diff-Testing

			if !(reflect.DeepEqual(cfg, tt.expected)) {
				// t.Logf("Actual: %#v\n", cfg)
				// t.Logf("Expected: %#v\n", tt.expected)
				t.Logf("Diff: \n%s", cmp.Diff(cfg, tt.expected))
				t.Errorf("Structs is not correctly generated. Actual: %#v, Expected: %#v, \n", cfg, tt.expected)

			}
		})

	}
}

func TestDatabaseIntegrationActive(t *testing.T) {
	var tests = []struct {
		inputEnvVars map[string]string
		expected     bool
	}{
		{
			map[string]string{
				"APP_POSTGRESQL_DB_CONNECT_STRING": "postgresql://postgresql.example.com",
			},
			true,
		},
		{
			map[string]string{},
			false,
		},
	}
	for ix, tt := range tests {
		testname := fmt.Sprintf("%d", ix)
		t.Run(testname, func(t *testing.T) {
			var cfg examples.CharmConfig
			err := env.ParseWithOptions(&cfg, env.Options{Environment: tt.inputEnvVars})
			if err != nil {
				t.Errorf("Error parsing config %s", err)
			}
			if cfg.Integrations.PostgreSQL.IsActive() != tt.expected {
				t.Errorf("actual active: %t, expected: %t", cfg.Integrations.PostgreSQL.IsActive(), tt.expected)
			}
			if cfg.Integrations.PostgreSQL.IsOptional() {
				t.Errorf("wrong")
			}
			if !cfg.Integrations.MongoDB.IsOptional() {
				t.Errorf("wrong")
			}
		})

	}
}
