package autoconf

import "github.com/sqlc-dev/sqlc/internal/config"

type SQLGo struct {
	PackageName               string `json:"package" yaml:"package"`
	EmitInterface             bool   `json:"emit_interface" yaml:"emit_interface"`
	EmitJSONTags              bool   `json:"emit_json_tags" yaml:"emit_json_tags"`
	EmitExportedQueries       bool   `json:"emit_exported_queries" yaml:"emit_exported_queries"`
	EmitMethodsWithDbArgument bool   `json:"emit_methods_with_db_argument" yaml:"emit_methods_with_db_argument"`
	Out                       string `json:"out" yaml:"out"`
}

type SQLGen struct {
	Go *SQLGo `json:"go,omitempty" yaml:"go"`
}

type SQL struct {
	Engine  config.Engine `json:"engine,omitempty" yaml:"engine"`
	Schema  config.Paths  `json:"schema" yaml:"schema"`
	Queries config.Paths  `json:"queries" yaml:"queries"`
	Gen     SQLGen        `json:"gen" yaml:"gen"`
}

type Config struct {
	Version string `json:"version" yaml:"version"`
	SQL     []SQL  `json:"sql" yaml:"sql"`
}
