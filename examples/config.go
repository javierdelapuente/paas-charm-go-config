package examples

type CharmConfig struct {
	ConfigOptions ConfigOptions
	ProxyConfig   ProxyConfig
	Integrations  Integrations
}

type ConfigOptions struct {
	Port        int     `env:"APP_PORT"`
	MetricsPort *int    `env:"APP_METRICS_PORT"`
	MetricsPath *string `env:"APP_METRICS_PATH"`
	SecretKey   *string `env:"APP_SECRET_KEY"`
	UserConfigOptions
}

type UserConfigOptions struct {
	// string | int | float | boolean | secret
	MyUserString  *string  `env:"APP_MY_USER_STRING"`
	MyUserInt     *int     `env:"APP_MY_USER_INT"`
	MyUserFloat   *float64 `env:"APP_MY_USER_FLOAT"`
	MyUserBoolean *bool    `env:"APP_MY_USER_BOOLEAN"`
}

type ProxyConfig struct {
	HTTPProxy  *string `env:"HTTP_PROXY"`
	HTTPSProxy *string `env:"HTTPS_PROXY"`
	// envKeyValSeparator default is `,` so it should work?
	NoProxy []string `env:"NO_PROXY"`
}

type Integrations struct {
	// mandatory integration
	PostgreSQL PostgreSQLIntegration `envPrefix:"APP_POSTGRESQL_"`
	// optional integration.
	// Grr this does not work https://github.com/caarlos0/env/issues/217
	// So MongoDB *DatabaseIntegration will not work :(
	MongoDB MongoDBIntegration `envPrefix:"APP_MONGODB_"`
}

type PostgreSQLIntegration struct {
	DatabaseIntegration
}

// Hardcoded to the value in charmcraft.yaml
// Maybe hardcode it?
func (s3 PostgreSQLIntegration) IsOptional() bool {
	return false
}

type MongoDBIntegration struct {
	DatabaseIntegration
}

// Hardcoded to the value in charmcraft.yaml
// Maybe hardcode it?
func (s3 MongoDBIntegration) IsOptional() bool {
	return true
}

func (di DatabaseIntegration) IsActive() bool {
	if di.ConnectString == "" {
		return false
	}
	return true
}

type DatabaseIntegration struct {
	ConnectString string  `env:"DB_CONNECT_STRING"`
	Scheme        string  `env:"DB_SCHEME"`
	NetLoc        string  `env:"DB_NETLOC"`
	Path          string  `env:"DB_PATH"`
	Params        string  `env:"DB_PARAMS"`
	Query         string  `env:"DB_QUERY"`
	Fragment      string  `env:"DB_FRAGMENT"`
	Username      *string `env:"DB_USERNAME"`
	Hostname      *string `env:"DB_HOSTNAME"`
	Port          *int    `env:"DB_PORT"`
}

type S3Integration struct {
	AccessKey    string  `env:"APP_S3_ACCES1S_KEY"`
	SecretKey    string  `env:"APP_S3_SECRET_KEY"`
	Region       *string `env:"APP_S3_REGION"`
	StorageClass *string `env:"APP_S3_STORAGE_CLASS"`
	Bucket       string  `env:"APP_S3_BUCKET"`
	Endpoint     *string `env:"APP_S3_ENDPOING"`
	Path         *string `env:"APP_S3_PATH"`
	ApiVersion   *string `env:"APP_S3_API_VERSION"`
	// TODO For this variables and similar ones,
	// should we document/restrict/whatever the possible values?
	UriStyle        *string `env:"APP_S3_URI_STYLE"`
	AddressingStyle *string `env:"APP_S3_ADDRESSING_STYLE"`
	// TODO THIS IS A JSON, WHAT DO I DO IN HERE?
	Attributes *string `env:"APP_S3_ATTRIBUTES"`
	// TODO THIS IS A JSON, WHAT DO I DO IN HERE?
	TLSCAChain *string `env:"APP_S3_TLS_CA_CHAIN"`
}

func (s3 S3Integration) IsActive() bool {
	if s3.AccessKey == "" {
		return false
	}
	return true
}

// Hardcoded to the value in charmcraft.yaml
func (_ S3Integration) IsOptional() bool {
	return true
}

type SAMLIntegration struct {
	// Review mandatory/optional things
	EntityID                string  `env:"SAML_ENTITY_ID"`
	MetadataURL             *string `env:"SAML_METADATA_URL"`
	SingleSignOnRedirectURL string  `env:"SAML_SINGLE_SIGN_ON_REDIRECT_URL"`
	SigningCertificat       string  `env:"SAML_SIGNING_CERTIFICATE"`
}

func (si SAMLIntegration) IsActive() bool {
	if si.EntityID == "" {
		return false
	}
	return true
}

// Hardcoded to the value in charmcraft.yaml
func (_ SAMLIntegration) IsOptional() bool {
	return true
}
