package database

import (
    "errors"
    "fmt"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
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
        host, port, user, password, dbname) // sslmode
    // try to connect if failed recover panic and try few other methods
    c, err := sqlx.Connect("pgx", url)
    if err != nil {
        return nil, err
    }
    db := &Database{
        Conn:     c,
        maxConns: defaultMaxConns,
    }

    driver, err := postgres.WithInstance(c.DB, &postgres.Config{})
    if err != nil {
        return nil, fmt.Errorf("migration driver: %w", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://database/migration/.",
        "postgres", driver)
    if err != nil {
        return nil, fmt.Errorf("migration file: %w", err)
    }

    err = m.Up()
    if err != nil && !errors.Is(err, migrate.ErrNoChange) {
        return nil, fmt.Errorf("migration up: %w", err)
    }

    /*
    	arguments := []string{}
    	if len(args) > 3 {
    		arguments = append(arguments, args[3:]...)
    	}

    	if err := goose.Run(command, db, *dir, arguments...); err != nil {
    		log.Fatalf("goose %v: %v", command, err)
    	}
    */

    /// todo migrate end

    return db, nil
}

func (db *Database) Close() {
    if db.Conn != nil {
        db.Conn.Close()
    }
}
