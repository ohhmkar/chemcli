package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "chem.db")
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS medicines (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		quantity INTEGER,
	    expiry_date TEXT,
	                                     
		unique(name, expiry_date)
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
