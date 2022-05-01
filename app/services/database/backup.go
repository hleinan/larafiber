package database

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func Backup() {
	fmt.Println("Starting to backup")
	c := cron.New()
	backupPath := "./database/backup/"
	checkOrCreateFolder(backupPath)

	backupDailyPath := "./database/backup/daily/"
	checkOrCreateFolder(backupDailyPath)

	c.AddFunc("0 * * * *", func() {
		time := time.Now()
		weekday := time.Weekday()
		if time.Hour() == 0 {
			path := strings.Join([]string{backupPath, strconv.Itoa(int(weekday))}, "")
			checkOrCreateFolder(path)

			file := strings.Join([]string{path, "/db.sqlite"}, "")
			backupDatabase(file)
		}
		path := strings.Join([]string{backupDailyPath, strconv.Itoa(time.Hour())}, "")
		checkOrCreateFolder(path)

		file := strings.Join([]string{path, "/db.sqlite"}, "")
		backupDatabase(file)
	})
	c.Start()
}

func backupDatabase(dst string) {
	fmt.Println(dst)
	src := "./database/db.sqlite"
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		fmt.Println(err)
	}

	source, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
	}
	defer destination.Close()
	io.Copy(destination, source)
}

func checkOrCreateFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Does not exist")
		err := createFolder(path)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func createFolder(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}