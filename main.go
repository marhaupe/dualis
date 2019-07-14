package main

import (
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
)

func init() {
	flag.StringVar(&password, "p", "", "password")
	flag.StringVar(&username, "u", "", "username")
	flag.Parse()
	if password == "" || username == "" {
		fmt.Println("Please supply both -u (username) and -p (password)")
		os.Exit(1)
	}
}

func main() {
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
		chromedp.SendKeys("#field_user", username),
		chromedp.SendKeys("#field_pass", password),
		chromedp.Click("#logIn_btn", chromedp.ByID),
		chromedp.WaitReady("#link000307 > a", chromedp.ByQuery),
		chromedp.Click("#link000307 > a", chromedp.ByQuery),
		chromedp.WaitReady("#contentSpacer_IE > div > table", chromedp.ByQuery),
		chromedp.Screenshot("#contentSpacer_IE > div > table", &result, chromedp.ByQuery),
	); err != nil {
		panic(err)
	}

	err := ioutil.WriteFile("grades.png", []byte(result), 0644)
	if err != nil {
		fmt.Println("Error while saving screenshot of your grades")
		os.Exit(2)
	}

	fmt.Println("Saved screenshot of your grades")
	open.Run("grades.png")

	time.Sleep(time.Second * 4)
	os.Remove("grades.png")
}
