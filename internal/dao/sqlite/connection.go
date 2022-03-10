package sqlite

import (
	"fmt"
	"log"
	"path"

	"github.com/afreeberg/k8s-homebase/internal/utilities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func get() *gorm.DB {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	// db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", config.Database)), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", path.Join(config.Datapath, config.Database))), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
