package migrate

import "embed"

//go:embed *.sql
var MigrateFS embed.FS

