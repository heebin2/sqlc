package genyaml

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sqlc-dev/sqlc/internal/config"
)

func GenerateYAML(stderr io.Writer, dir, filename, engine string) error {
	configFile := readConfigFileName(dir, filename)

	schemaFile, err := findSchemaSQL(dir)
	if err != nil {
		return errors.New("error find schema.sql file: file does not exist")
	}

	fmt.Println("schema file : ", schemaFile)

	queries, err := findQuerySQL(dir)
	if err != nil {
		return errors.New("error find query file")
	}

	cfg, err := makeConfig(schemaFile, queries)
	if err != nil {
		return errors.New("error make config")
	}

	return writeYAML(cfg, configFile)
}

func readConfigFileName(dir, filename string) string {
	if filename != "" {
		return filepath.Join(dir, filename)
	}

	return filepath.Join(dir, "sqlc.yaml")
}

func findSchemaSQL(dir string) (string, error) {
	ret := ""
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Base(path) == "schema.sql" {
			ret = path
			return errors.New("find")
		}

		return nil
	})

	if ret == "" {
		return "", errors.New("not found")
	}

	return ret, nil
}

func findQuerySQL(dir string) ([]string, error) {

	return nil, nil
}

func makeConfig(schemaFile string, queries []string) (config.Config, error) {
	return config.Config{}, nil
}

func writeYAML(cfg config.Config, filename string) error {

	return nil
}
