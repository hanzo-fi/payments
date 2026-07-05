package bunconnect

import (
	"context"
	"testing"

	"github.com/spf13/pflag"
)

// TestSQLiteRoundTrip proves the sqlite dialect+driver seam is functional:
// open -> create table -> insert -> select back, all through bun.
func TestSQLiteRoundTrip(t *testing.T) {
	db, err := OpenSQLiteDB("file:roundtrip?mode=memory&cache=shared&_pragma=foreign_keys(1)")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	if _, err := db.NewCreateTable().Model((*connectorRow)(nil)).IfNotExists().Exec(ctx); err != nil {
		t.Fatalf("create table: %v", err)
	}
	if _, err := db.NewInsert().Model(&connectorRow{Name: "stripe", Enabled: true}).Exec(ctx); err != nil {
		t.Fatalf("insert: %v", err)
	}
	var got connectorRow
	if err := db.NewSelect().Model(&got).Where("name = ?", "stripe").Scan(ctx); err != nil {
		t.Fatalf("select: %v", err)
	}
	if got.Name != "stripe" || !got.Enabled {
		t.Fatalf("unexpected row: %+v", got)
	}
	if name := db.Dialect().Name().String(); name != "sqlite" {
		t.Fatalf("expected sqlite dialect, got %q", name)
	}
}

type connectorRow struct {
	ID      int64  `bun:"id,pk,autoincrement"`
	Name    string `bun:"name,notnull"`
	Enabled bool   `bun:"enabled,notnull"`
}

func TestFromFlagsDefaultsAndValidation(t *testing.T) {
	empty := pflag.NewFlagSet("empty", pflag.ContinueOnError)
	d, dsn, err := FromFlags(empty)
	if err != nil || d != DriverSQLite || dsn != defaultSQLiteDSN {
		t.Fatalf("unregistered: got (%q,%q,%v)", d, dsn, err)
	}

	fs := pflag.NewFlagSet("fs", pflag.ContinueOnError)
	AddFlags(fs)
	_ = fs.Set(StorageDriverFlag, "postgres")
	if d, _, err := FromFlags(fs); err != nil || d != DriverPostgres {
		t.Fatalf("postgres: got (%q,%v)", d, err)
	}

	fs2 := pflag.NewFlagSet("fs2", pflag.ContinueOnError)
	AddFlags(fs2)
	_ = fs2.Set(StorageDriverFlag, "mysql")
	if _, _, err := FromFlags(fs2); err == nil {
		t.Fatalf("expected error for invalid driver")
	}
}
