package storage

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func RunMigrations(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
            version VARCHAR(255) PRIMARY KEY,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	files, err := filepath.Glob("internal/migrations/*.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %v", err)
	}

	sort.Strings(files)

	for _, file := range files {
		filename := filepath.Base(file)

		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE version = $1", filename).Scan(&count)
		if err != nil {
			return fmt.Errorf("failed to check migration: %v", err)
		}

		if count > 0 {
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file, err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to apply migration %s: %v", file, err)
		}

		_, err = db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", filename)
		if err != nil {
			return fmt.Errorf("failed to record migration %s: %v", file, err)
		}

		fmt.Printf("✅ Applied migration: %s\n", filename)
	}

	return nil
}
