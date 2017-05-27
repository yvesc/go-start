package main

import (
	"fmt"
	"net/http"
	
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go_sqlite3"

	"github.com/qor/qor"
	"github.com/qor/admin"
)

// Create another GORM-backend model
type Product struct {
  gorm.Model
  Name        string
  Description string
}


func main() {
  DB, _ := gorm.Open("sqlite3", "demo.db")
  DB.AutoMigrate(&User{}, &Product{})

  // Initalize
  Admin := admin.New(&qor.Config{DB: DB})

  // Create resources from GORM-backend model
  Admin.AddResource(&User{})
  Admin.AddResource(&Product{})

  // Register route
  mux := http.NewServeMux()
  // amount to /admin, so visit `/admin` to view the admin interface
  Admin.MountTo("/admin", mux)

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", mux)
}

