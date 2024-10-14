

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" // Assuming you're using PostgreSQL; adjust based on your DB
)

func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	// Begin a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Ensure commit or rollback happens
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	// First insert
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) VALUES ($1)", value1)
	if err != nil {
		return err
	}

	// Second insert (just an example; you can add more operations)
	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) VALUES ($1)", value2)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Assuming you have a PostgreSQL DB running
	connStr := "user=username dbname=mydb sslmode=disable" // Modify accordingly
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Example usage
	ctx := context.Background()
	err = DoSomeInserts(ctx, db, "value1", "value2")
	if err != nil {
		fmt.Println("Error during inserts:", err)
	} else {
		fmt.Println("Inserts successful")
	}
}
