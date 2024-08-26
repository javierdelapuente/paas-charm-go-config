package main_test

import (
	"testing"

	"reflect"

	"github.com/google/go-cmp/cmp"
	generator "github.com/javierdelapuente/paas-charm-go-generator/cmd/paasgoconfiggenerator"
)

func TestParseCharmcraftYaml(t *testing.T) {
	yamlData := `
name: go-app
type: charm
base: ubuntu@24.04
platforms:
  amd64:
  arm64:
  armhf:
  ppc64el:
  riscv64:
  s390x:
summary: A very short one-line summary of the Go application.
description: |
  A comprehensive overview of your Go application.
extensions:
  - go-framework
config:
  options:
    user-defined-str:
      type: string
      default: "hello"
      description: user-defined-str Description
    user-defined-int:
      type: int
      default: 100
      description: user-defined-int Description
    user-defined-bool:
      type: bool
      description: user-defined-bool Description
requires:
  mysql:
    interface: mysql_client
    optional: true
    limit: 1
  s3:
    interface: s3
    optional: false

parts: {0-git: {plugin: nil, build-packages: [git]}}
`
	charmcraft, err := generator.ParseCharmcraftYaml([]byte(yamlData))

	if err != nil {
		t.Errorf("Error parsing data")
	}

	expected := generator.CharmcraftYamlConfig{}
	expected.Config.Options = map[string]generator.ConfigOption{
		"user-defined-str":  {Type: "string", Default: "hello", Description: "user-defined-str Description"},
		"user-defined-int":  {Type: "int", Default: 100, Description: "user-defined-int Description"},
		"user-defined-bool": {Type: "bool", Default: nil, Description: "user-defined-bool Description"},
	}
	expected.Requires = map[string]generator.IntegrationConfig{
		"mysql": {Interface: "mysql_client", Optional: true},
		"s3":    {Interface: "s3", Optional: false},
	}

	if !(reflect.DeepEqual(charmcraft, expected)) {
		t.Logf("Diff: \n%s", cmp.Diff(charmcraft, expected))
		t.Errorf("Structs is not correctly generated. Actual: %#v, Expected: %#v, \n", charmcraft, expected)
	}
}
