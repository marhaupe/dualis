package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/skratchdot/open-golang/open"
)

var (
	password string
	username string

	filename = "grades.png"
)

func init() {
	flag.StringVar(&password, "p", "", "password")
	flag.StringVar(&username, "u", "", "username")
	flag.Parse()
}

func main() {
	if password == "" || username == "" {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Username: ")
		username, _ = reader.ReadString('\n')
		username = username[:len(username)-1]

		fmt.Print("Password: ")
		password, _ = reader.ReadString('\n')
		password = password[:len(password)-1]
	}

	result, err := generateScreenshot()
	if err != nil {
		fmt.Println("error while taking screenshot of your grades", err)
	}

	err = ioutil.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		fmt.Println("error while saving screenshot of your grades", err)
		os.Exit(2)
	}

	open.Run(filename)

	time.Sleep(time.Second * 4)
	os.Remove(filename)
}

func generateScreenshot() ([]byte, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var result []byte
	if err := chromedp.Run(taskCtx,
		chromedp.Navigate("https://dualis.dhbw.de/scripts/mgrqispi.dll?APPNAME=CampusNet&PRGNAME=EXTERNALPAGES&ARGUMENTS=-N139906720206530,-N000307,"),

		// Login
		chromedp.WaitReady("#field_user", chromedp.ByQuery),
		chromedp.SendKeys("#field_user", username),
		chromedp.SendKeys("#field_pass", password),
		chromedp.Click("#logIn_btn", chromedp.ByID),

		// Press `Prüfungsergebnisse`
		chromedp.WaitReady("#link000307 > a", chromedp.ByQuery),
		chromedp.Click("#link000307 > a", chromedp.ByQuery),

		// Screenshot table
		chromedp.WaitReady("#contentSpacer_IE > div > table", chromedp.ByQuery),
		chromedp.Screenshot("#contentSpacer_IE > div > table", &result, chromedp.ByQuery),
	); err != nil {
		return nil, err
	}
	return result, nil
}
