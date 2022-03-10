package factory

import (
	"log"

	"github.com/afreeberg/k8s-homebase/internal/dao/interfaces"
	"github.com/afreeberg/k8s-homebase/internal/dao/psql"
	"github.com/afreeberg/k8s-homebase/internal/dao/sqlite"
)

func UserDaoFactory(e string) interfaces.UserDao {
	var i interfaces.UserDao
	switch e {
	case "postgres":
		i = psql.UserImplPsql{}
	case "sqlite":
		i = sqlite.UserImplSqlite{}
	default:
		log.Fatalf("No Engine, or Engine not implemented! Val: %s\n", e)
		return nil
	}
	return i
}
