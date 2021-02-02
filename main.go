//go:generate go run assets-generator.go

package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/emptyhua/shiori/internal/cmd"
	"github.com/sirupsen/logrus"

	// Database driver
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	// Add this to prevent it removed by go mod tidy
	_ "github.com/shurcooL/vfsgen"
)

func main() {
	go func() {
		http.ListenAndServe("127.0.0.1:9101", nil)
	}()

	err := cmd.ShioriCmd().Execute()
	if err != nil {
		logrus.Fatalln(err)
	}
}
