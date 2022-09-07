package neeorm

import (
	"database/sql"
	"neeorm/dialect"
	"neeorm/log"
	"neeorm/session"
)

// Engine
// 交互后的收尾工作
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, er error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is value
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	// make sure to make sure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.ErrorF("dialect %s not Found", driver)
		return
	}
	e = &Engine{
		db:      db,
		dialect: dial,
	}

	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
