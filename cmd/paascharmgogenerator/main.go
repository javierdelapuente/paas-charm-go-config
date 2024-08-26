package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

func run(charmcraftDir string, packageName string, outputFile string) (err error) {
	yamlFile, err := os.ReadFile(path.Join(charmcraftDir, CharmcraftFileName))
	if err != nil {
		return
	}

	charmcraftConfig, err := ParseCharmcraftYaml(yamlFile)
	if err != nil {
		return
	}

	goStructs, err := GenerateGoStructs(packageName, charmcraftConfig)
	if err != nil {
		return
	}

	// TODO CHECK FILE EXISTS. This is not very correct...
	// What if directories do not exist? build them?
	// Do we want to override the file?
	if _, err := os.Stat(outputFile); err == nil {
		return fmt.Errorf("Output file already %s exists", outputFile)
	}
	// WRITE FILE
	// You can also write it to a file as a whole?
	err = os.WriteFile(outputFile, goStructs, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	log.Println("let's go")
	charmcraftDir := flag.String("charmcraft-dir", ".", "charmcraft.yaml file location")
	packageName := flag.String("package", "charmconfig", "name of the generated package")
	outputFile := flag.String("output-file", "charmconfig/charmconfig.go", "generated file name")
	flag.Parse()

	log.Printf("charmcraftDir: %s\n", *charmcraftDir)
	log.Printf("package: %s\n", *packageName)
	log.Printf("output file: %s\n", *outputFile)

	err := run(*charmcraftDir, *packageName, *outputFile)
	if err != nil {
		log.Fatalf("Error generating go configuration. Error: %v", err)
	}
}
