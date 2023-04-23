package main

import (
	"fmt"
	"os"

	"vitess.io/vitess/go/vt/sqlparser"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <sql-file>", os.Args[0])
		os.Exit(0)
	}

	sqlfile := os.Args[1]
	data, err := os.ReadFile(sqlfile)
	if err != nil {
		panic(err)
	}

	stmt, err := sqlparser.Parse(string(data))
	if err != nil {
		panic(err)
	}

	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		fmt.Println(sqlparser.CanonicalString(stmt))
	default:
		fmt.Printf("%T\n", stmt)
		fmt.Println(sqlparser.String(stmt))
	}
}
