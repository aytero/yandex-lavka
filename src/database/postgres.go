package database

import (
    "fmt"
    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/jmoiron/sqlx"
)

const (
    defaultMaxConns = 5
    //    _defaultMaxPoolSize  = 1
    //    _defaultConnAttempts = 10
    //    _defaultConnTimeout  = time.Second
)

// todo rename to postgres

type Database struct {
    //maxPoolSize  int
    //connAttempts int
    //connTimeout  time.Duration
    maxConns int
    Conn     *sqlx.DB
}

const (
    host     = "db"
    port     = 5432
    user     = "postgres"
    password = "password"
    dbname   = "ydx_db"
)

func New(url string) (*Database, error) {
    // todo
    url = fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    //c, err := sqlx.Connect("pgx", psqlInfo)
    c, err := sqlx.Connect("pgx", url)
    if err != nil {
        return nil, err
    }
    db := &Database{
        Conn:     c,
        maxConns: defaultMaxConns,
    }

    //	for _, opt := range opts {
    //		opt(db)
    //	}
    db.Conn.SetMaxIdleConns(db.maxConns)
    db.Conn.SetMaxOpenConns(db.maxConns)

    // todo migrations

    return db, nil
}

func (db *Database) Close() {
    if db.Conn != nil {
        db.Conn.Close()
    }
}

/*
func New() *gorm.DB {
	db, err := gorm.Open("postgres", "./realworld.db")
	if err != nil {
		fmt.Println("repository err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

//func TestDB() *gorm.DB {
//	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
//	if err != nil {
//		fmt.Println("storage err: ", err)
//	}
//	db.DB().SetMaxIdleConns(3)
//	db.LogMode(false)
//	return db
//}

//func DropTestDB() error {
//	if err := os.Remove("./../realworld_test.db"); err != nil {
//		return err
//	}
//	return nil
//}

//DO: err check
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.OrderDto{},
		&model.CourierDto{},
		//&model.Article{},
		//&model.Comment{},
		//&model.Tag{},
	)
}

*/
