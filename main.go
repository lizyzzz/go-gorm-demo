package main

import (
	gormdemo "go-gorm-demo/gorm-demo"
	sqldemo "go-gorm-demo/sql-demo"
)

type User struct {
	Username string
	Password string
}

func main() {
	sqldemo.SqlDemo()
	// gormdemo.CreateTable()
	// gormdemo.InsertTable()
	// gormdemo.QueryTable()
	// gormdemo.UpdateTable()
	// gormdemo.DeleteTable()
	// gormdemo.TestHook()
	// gormdemo.PrepareData()
	// gormdemo.AdvancedQuery()

	// gormdemo.One2MoreCreateTable()
	// gormdemo.One2MoreInsertTable()
	// gormdemo.One2MoreQuery()
	// gormdemo.One2MoreDelete()

	// gormdemo.One2OneCreateTable()
	// gormdemo.One2OneInsertTable()
	// gormdemo.One2OneQuery()
	// gormdemo.One2OneDelete()

	// gormdemo.More2MoreCreateTable()
	// gormdemo.More2MoreInsertTable()
	// gormdemo.More2MoreQuery()
	// gormdemo.More2MoreUpdate()
	// gormdemo.More2MoreDelete()
	// gormdemo.DefineConnTable()
	// gormdemo.UserDefinedConnTableOP()

	// gormdemo.QueryJoinTable()
	// gormdemo.CreateUserDataTable()

	gormdemo.CreateTxTable()
	gormdemo.TxExample()
}
