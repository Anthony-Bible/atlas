// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package x_test

import (
	"context"
	"testing"
	_ "unsafe"

	"github.com/anthony-bible/atlas/cmd/atlas/x"
	"github.com/anthony-bible/atlas/schemahcl"
	"github.com/anthony-bible/atlas/sql/migrate"
	"github.com/anthony-bible/atlas/sql/sqlcheck"
	"github.com/anthony-bible/atlas/sql/sqlclient"
	_ "github.com/anthony-bible/atlas/sql/sqlite"
	_ "github.com/anthony-bible/atlas/sql/sqlite/sqlitecheck"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestLintLatest(t *testing.T) {
	ctx := context.Background()
	dev, err := sqlclient.Open(ctx, "sqlite://ci?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	dir := &migrate.MemDir{}
	require.NoError(t, dir.WriteFile("1.sql", []byte(`CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL);`)))
	az, err := sqlcheck.AnalyzerFor(dev.Name, &schemahcl.Resource{})
	require.NoError(t, err)
	report, err := lintLatest(ctx, dev, dir, 1, az)
	require.NoError(t, err)
	require.NotNil(t, report)
}

//go:linkname lintLatest github.com/anthony-bible/atlas/cmd/atlas/x.lintLatest
func lintLatest(context.Context, *sqlclient.Client, migrate.Dir, int, []sqlcheck.Analyzer) (report *x.SummaryReport, err error)
