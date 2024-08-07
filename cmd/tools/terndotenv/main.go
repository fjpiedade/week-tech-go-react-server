// package main

// import (
// 	"os/exec"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		panic(err)
// 	}

// 	cmd := exec.Command(
// 		"tern",
// 		"migrate",
// 		"--migrations",
// 		"./internal/store/pgstore/migrations",
// 		"--config",
// 		"./internal/store/pgstore/migrations/tern.config",
// 	)

// 	if err := cmd.Run(); err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		//panic(err)
		log.Fatal("Error loading .env file")
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running tern migrate:", err)
		// Consider adding more specific error handling here
		return
	}

	fmt.Println("Tern migration completed successfully!")
}
