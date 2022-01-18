package main

import (
	"github.com/et-nik/otus-highload/internal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	internal.Run()
}
