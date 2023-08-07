package autoconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sqlc-dev/sqlc/internal/config"
	"gopkg.in/yaml.v3"
)

const dbDir = "db"

func GenerateConfig(stderr io.Writer, dir, filename, engine string) error {
	configFile := readConfigFileName(dir, filename)
	dir = filepath.Join(dir, dbDir)

	schemaFile, err := findSchemaSQL(dir)
	if err != nil {
		return errors.New("error find schema.sql file: file does not exist")
	}

	queries, err := findQuerySQL(dir)
	if err != nil {
		return errors.New("error find query file")
	}

	cfg, err := makeConfig(engine, schemaFile, queries)
	if err != nil {
		return errors.New("error make config")
	}

	if strings.Contains(configFile, "json") {
		return writeJson(cfg, configFile)
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
			rel, _ := filepath.Rel(dir, path)
			ret = filepath.Join(dbDir, rel)
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
	ret := []string{}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Base(path) == "schema.sql" {
			return nil
		}

		if filepath.Ext(path) == ".sql" {
			rel, err := filepath.Rel(dir, path)
			if err != nil {
				return nil
			}

			ret = append(ret, filepath.Join(dbDir, rel))
		}

		return nil
	})

	if len(ret) <= 0 {
		return nil, fmt.Errorf("not found .sql file")
	}

	return ret, nil
}

func makeConfig(engien, schemaFile string, queries []string) (Config, error) {
	cfg := Config{
		Version: "2",
	}

	for _, query := range queries {
		sql := SQL{
			Engine:  config.Engine(engien),
			Schema:  config.Paths{schemaFile},
			Queries: config.Paths{query},
			Gen: SQLGen{
				Go: &SQLGo{
					Out:                 filepath.Join("internal", query),
					EmitInterface:       true,
					EmitExportedQueries: true,
					EmitJSONTags:        true,
					EmitExactTableNames: true,
				},
			},
		}

		cfg.SQL = append(cfg.SQL, sql)
	}

	return cfg, nil
}

func writeYAML(cfg Config, filename string) error {
	d, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, d, 0644)
}

func writeJson(cfg Config, filename string) error {
	d, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, d, 0644)
}
