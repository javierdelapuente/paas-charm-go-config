Example environment variables:

- generic charm
APP_BASE_URL

- FRAMEWORK CONFIG
APP_PORT
APP_METRICS_PORT
APP_METRICS_PATH
APP_LOG_LEVEL

- PROXY
HTTP_PROXY
HTTPS_PROXY
NO_PROXY

- DATABASES
with base_name in: REDIS, MYSQL, POSTGRESQL, MONGODB
APP_{base_name}_DB_SCHEME    default empty
APP_{base_name}_DB_NETLOC    default empty
APP_{base_name}_DB_PATH      default empty
APP_{base_name}_DB_PARAMS    default empty
APP_{base_name}_DB_QUERY     default empty
APP_{base_name}_DB_FRAGMENT  default empty
APP_{base_name}_DB_USERNAME  not mandatory
APP_{base_name}_DB_PASSWORD  not mandatory
APP_{base_name}_DB_HOSTNAME  not mandatory
APP_{base_name}_DB_PORT      not mandatory

- S3
S3_ACCESS_KEY  mandatory
S3_SECRET_KEY  mandatory
S3_REGION      this or uri mandatory?
S3_STORAGE_CLASS
S3_BUCKET      mandatory
S3_ENDPOINT
S3_PATH
S3_API_VERSION
S3_URI_STYLE
S3_ADDRESSING_STYLE
S3_ATTRIBUTES    this is a json, what can be in there?
S3_TLS_CA_CHAIN  this is a json, what can be in there?

- SAML
SAML_ENTITY_ID
SAML_METADATA_URL    optional?
SAML_SINGLE_SIGN_ON_REDIRECT_URL
SAML_SIGNING_CERTIFICATE


# Real examples:
## 1
            APP_BASE_URL: http://go-k8s.testing:8080
            APP_METRICS_PATH: /metrics
            APP_METRICS_PORT: "8080"
            APP_PORT: "8080"
            APP_SECRET_KEY: _6vLis5ukuN8p03AMQClOdhDPBsJ2YvjbLK_-9l8BVHmttsAzXPZiUcHKhRLb8-KUzpDReFEO8yhwzIVQ0OA0Q
            NO_PROXY: 127.0.0.1,localhost,::1
            no_proxy: 127.0.0.1,localhost,::1

## 2
            APP_BASE_URL: http://go-k8s.testing:8080
            APP_METRICS_PATH: /metrics
            APP_METRICS_PORT: "8080"
            APP_PORT: "8080"
            APP_POSTGRESQL_DB_CONNECT_STRING: postgresql://relation_id_4:MTqX51qDce6o4fpb@postgresql-k8s-primary.testing.svc.cluster.local:5432/go-k8s
            APP_POSTGRESQL_DB_FRAGMENT: ""
            APP_POSTGRESQL_DB_HOSTNAME: postgresql-k8s-primary.testing.svc.cluster.local
            APP_POSTGRESQL_DB_NAME: go-k8s
            APP_POSTGRESQL_DB_NETLOC: relation_id_4:MTqX51qDce6o4fpb@postgresql-k8s-primary.testing.svc.cluster.local:5432
            APP_POSTGRESQL_DB_PARAMS: ""
            APP_POSTGRESQL_DB_PASSWORD: MTqX51qDce6o4fpb
            APP_POSTGRESQL_DB_PATH: /go-k8s
            APP_POSTGRESQL_DB_PORT: "5432"
            APP_POSTGRESQL_DB_QUERY: ""
            APP_POSTGRESQL_DB_SCHEME: postgresql
            APP_POSTGRESQL_DB_USERNAME: relation_id_4
            APP_SECRET_KEY: _6vLis5ukuN8p03AMQClOdhDPBsJ2YvjbLK_-9l8BVHmttsAzXPZiUcHKhRLb8-KUzpDReFEO8yhwzIVQ0OA0Q
            APP_USER-DEFINED-CONFIG: newvalue
            APP_USER_DEFINED_CONFIG: newvalue
            NO_PROXY: 127.0.0.1,localhost,::1
            no_proxy: 127.0.0.1,localhost,::1

