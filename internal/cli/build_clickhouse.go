// +build clickhouse

package cli

import (
	_ "github.com/orkusinc/migrate/database/clickhouse"
	_ "github.com/kshvakov/clickhouse"
)
