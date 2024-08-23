package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"log"
	"text/template"

	"strings"

	"unicode"

	"gopkg.in/yaml.v3"
)

const CharmcraftFileName string = "charmcraft.yaml"

// type ConfigOption struct {
// 	Type        string
// 	Description string
//}

type ConfigOption struct {
	Type        string
	Default     interface{}
	Description string
}

type IntegrationConfig struct {
	Interface string
	Optional  bool
}

type CharmcraftYamlConfig struct {
	Config struct {
		// Options map[string]map[interface{}]interface{} // OK
		// Options map[string]ConfigOption // OK
		Options map[string]ConfigOption
	}
	Requires map[string]IntegrationConfig
}

func ParseCharmcraftYaml(yamlData []byte) (charmcraft CharmcraftYamlConfig, err error) {
	err = yaml.Unmarshal(yamlData, &charmcraft)
	if err != nil {
		log.Printf("Error unmarshalling: %v\n", err)
		return
	}

	return
}

type templateConfig struct {
	PackageName string
	Charmcraft  CharmcraftYamlConfig
}

//go:embed config.template
var GoTemplate string

func GenerateGoStructs(packageName string, charmcraftConfig CharmcraftYamlConfig) (goStructs []byte, err error) {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"makeGoVariable": makeGoVariable, "makeEnvVariable": makeEnvVariable, "makeGoType": makeGoType,
		"isDatabaseIntegration": isDatabaseIntegration, "makeIntegrationName": makeIntegrationName,
	}).Parse(GoTemplate)
	fmt.Printf("templates: %v \n", tmpl.DefinedTemplates())
	if err != nil {
		fmt.Printf("Error Parsing Template file")
		return
	}

	templateConfig := templateConfig{
		PackageName: packageName,
		Charmcraft:  charmcraftConfig,
	}
	var buf bytes.Buffer

	err = tmpl.Execute(&buf, templateConfig)
	if err != nil {
		log.Printf("Error Executing Template.\n")
		return
	}

	srcFormatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("Error Go Code Formatting.\n%s\n", buf.String())
		return
	}

	return srcFormatted, nil
}

func makeEnvVariable(name string, prefix string) (result string) {
	result = prefix + name
	result = strings.ReplaceAll(result, "-", "_")
	result = strings.ToUpper(result)
	return result
}

func makeGoVariable(name string) (result string) {
	parts := strings.Split(name, "-")

	for _, part := range parts {
		partRunes := []rune(part)
		// TODO WHAT HAPPNES WITH A THING LIKE user--other?
		// and if there are two, one user-other and other user--other?
		// this will crash.
		if len(partRunes) > 0 {
			partRunes[0] = unicode.ToUpper(partRunes[0])
			result += string(partRunes)
		}
	}
	return
}

func makeGoType(configOption ConfigOption) (result string) {
	// https://juju.is/docs/sdk/charmcraft-yaml#heading--config
	fmt.Printf(" IN FUNCTION ")
	switch configOption.Type {
	case "string":
		result = "string"
	case "int":
		result = "int"
	case "boolean":
		result = "bool"
	case "bool":
		result = "bool"
	case "secret":
		// TODO IS THIS CORRECT?
		result = "string"
	default:
		log.Printf("Invalid type for a config option of type: %s. Returning string.\n", configOption.Type)
		result = "string"
	}

	if configOption.Default == nil {
		result = "*" + result
	}

	return
}

func makeIntegrationName(name string) (result string) {
	switch name {
	case "redis":
		result = "Redis"
	case "mysql":
		result = "MySQL"
	case "postgresql":
		result = "PostgreSQL"
	case "mongodb":
		result = "MongoDB"
	default:
		log.Printf("Invalid INTEGRATION NAME?")
		// TODO CRASH, SHOULD NOT GET HERE?
	}
	return
}

func isDatabaseIntegration(name string) bool {
	// TODO a cleaner/less ugly way?
	switch name {
	case "redis":
		return true
	case "mysql":
		return true
	case "postgresql":
		return true
	case "mongodb":
		return true
	}
	return false
}
