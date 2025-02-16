package wasm

import (
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestEvent(t *testing.T) {

	t.Parallel()

	wasmTmpDir, server := startServer(t)

	err := run("tinygo build -o " + wasmTmpDir + "/event.wasm -target wasm testdata/event.go")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := chromectx(5 * time.Second)
	defer cancel()

	var log1, log2 string
	err = chromedp.Run(ctx,
		chromedp.Navigate(server.URL+"/run?file=event.wasm"),
		chromedp.WaitVisible("#log"),
		chromedp.Sleep(time.Second),
		chromedp.InnerHTML("#log", &log1),
		waitLog(`1
4`),
		chromedp.Click("#testbtn"),
		chromedp.Sleep(time.Second),
		chromedp.InnerHTML("#log", &log2),
		waitLog(`1
4
2
3
true`),
	)
	t.Logf("log1: %s", log1)
	t.Logf("log2: %s", log2)
	if err != nil {
		t.Fatal(err)
	}

}
