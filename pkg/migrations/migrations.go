package migrations

import (
	"sync"

	migrate "github.com/rubenv/sql-migrate"
)

type dbmigrations struct {
	m          sync.Mutex
	migrations []*migrate.Migration
}

var instance = &dbmigrations{
	m:          sync.Mutex{},
	migrations: make([]*migrate.Migration, 0),
}

func (m *dbmigrations) add(migration *migrate.Migration) {
	m.m.Lock()
	m.migrations = append(m.migrations, migration)
	m.m.Unlock()
}

func GetAll() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: instance.migrations,
	}
}
