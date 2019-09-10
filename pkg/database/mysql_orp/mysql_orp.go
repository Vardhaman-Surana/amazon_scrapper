package mysql_orp

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
	"github.com/vds/amazon_scrapper/pkg/getenv"
	"github.com/vds/amazon_scrapper/pkg/migrations"
	"github.com/vds/amazon_scrapper/pkg/models"
	"gopkg.in/gorp.v1"
	"sync"
)

var (
	globalDB  *gorp.DbMap
	once     sync.Once
)
func NewDBmap()(*gorp.DbMap, error){
	var initErr error
	once.Do(func() {
		env,err:=getenv.GetDBEnv()
		if err!=nil{
			initErr = err
			globalDB= nil
		}else{
			dbMap, err := DBForURL(env.URL)
			initErr = err
			globalDB= dbMap
		}
	})
	return globalDB,initErr
}

func DBForURL(url string) (*gorp.DbMap, error){
	fmt.Printf("Creating DB with url %s", url)

	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("Error creating DB:", err)
		fmt.Println("To verify, db is:", db)
		return nil,err
	} /*else {
		err = db.Ping()
		if err != nil {
			time.Sleep(5 * time.Second)
			fmt.Print("trying to reconnect")
			return DBForURL(url)
		}
	}*/
	_, err = migrate.Exec(db, "mysql", migrations.GetAll(), migrate.Up)
	if err != nil {
		return nil, err
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	dbMap.AddTableWithName(models.Product{},models.ProductTableName).SetKeys(true, "ID")

	return dbMap,nil
}



