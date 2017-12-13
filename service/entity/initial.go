package entity

import (
	
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// ORM engine
var engine *xorm.Engine

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {

	var err error
	// create engine
	engine, err = xorm.NewEngine("sqlite3", "./agenda.db")
	checkErr(err)

	engine.SetMapper(core.GonicMapper{})

	// sync the struct changes to database
	checkErr(engine.Sync2(new(User), new(Met)))
}
