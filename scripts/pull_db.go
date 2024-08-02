// configuration.go
package main

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var dsn = os.Getenv("POSTGRES_CONN_STRING")

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/pg/dbquery", // output directory, default value is ./query
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		ModelPkgPath:      "dbmodels",
	})

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection err: %v\n", err)
	}

	g.UseDB(db)

	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if strings.HasPrefix(tableName, "_") {
			return ""
		}
		return tableName
	})

	g.ApplyBasic(
		g.GenerateModelAs("users", "User"),
		g.GenerateModelAs("admins", "Admin"),
		g.GenerateModelAs("broadcast_messages", "BroadcastMessage"),
		g.GenerateModelAs("buttons", "Button"),
		g.GenerateModelAs("directus_files", "File"),
	)

	// Execute the generator
	g.Execute()
}
