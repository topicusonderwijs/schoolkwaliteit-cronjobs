package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"skw.mijnschoolteam/mijnschoolteam/mysql"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("mysql", "user:password@/db_name?parseTime=true")
	if err != nil {
		log.Fatalf("Could not connect with mysql database")
		os.Exit(1)
	}

	queries := mysql.New(db)

	// Deactive all accounts
	accounts, err := queries.GetExpiredAccounts(ctx)
	if err != nil {
		log.Fatal("Could not get expired accounts")
	}

	var nFailed int
	for _, account := range accounts {
		err = queries.DeactivateExpiredAccountById(ctx, account.ID)
		if err != nil {
			log.Fatalf("Could not deactivate expired account %s", account.ID)
			nFailed++
		}
	}

	log.Printf("Finished deactivating accounts. %s failed", nFailed)
}
