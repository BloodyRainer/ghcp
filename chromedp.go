package ghcp

import (
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	dp "github.com/chromedp/chromedp"
	"log"
	"context"
)

var cdp *chromedp.CDP
var ctx context.Context
var cancel context.CancelFunc

func initChromeDp() {

	var err error

	//create context
	ctx, cancel = context.WithCancel(context.Background())
	//defer cancel()

	runOpts := dp.WithRunnerOptions(
		runner.HeadlessPathPort("/usr/bin/google-chrome", 9222),
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("no-first-run", true),
		runner.Flag("no-default-browser-check", true),
	)

	cdp, err = dp.New(ctx, runOpts, dp.WithLog(log.Printf))
	if err != nil {
		log.Panic(err)
	}

}

func shutDownChromeDP() {

	defer cancel()

	err := cdp.Shutdown(ctx)
	if err != nil {
		log.Panic(err)
	}

	// dont use for headless chrome
	//err = c.Wait()
	//if err != nil {
	//	log.Panic(err)
	//}

	log.Println("ChromeDP finished!")
}

func cdpRun(action dp.Action) error {
	return cdp.Run(ctx, action)
}
