package main

import (
	"context"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/meraiku/db_tasks/internal/helpers"
)

func main() {
	godotenv.Load()
	ctx := context.TODO()

	if err := helpers.ResetDB(ctx); err != nil {
		log.Fatalf("failed to reset DB: %v", err)
	}

	log.Println("DB reset successfully")

	if err := helpers.InsertData(ctx); err != nil {
		log.Fatalf("failed to insert data in DB: %v", err)
	}

	log.Println("Data inserted successfully")

	cmds := exec.CommandContext(ctx, "make", "tests")

	out, err := cmds.Output()
	if err != nil {
		log.Println(string(out))
		log.Fatal("Tests failed :(")
	}

	log.Println(string(out))
}
