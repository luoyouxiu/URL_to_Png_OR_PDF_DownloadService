/*
 * @Author: luoyouxiu 1291838675@qq.com
 * @Date: 2025-01-20 01:48:47
 * @LastEditors: luoyouxiu 1291838675@qq.com
 * @LastEditTime: 2025-01-20 03:41:13
 * @FilePath: \go\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	http.HandleFunc("/render", renderHandler)
	log.Println("Starting server on :52001...")
	log.Fatal(http.ListenAndServe(":52001", nil))
}

func renderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rawURL := r.URL.Query().Get("url")
	if rawURL == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	format := r.URL.Query().Get("format")
	if format == "" {
		format = "pdf" // 默认格式为pdf
	}

	// URL encode the parameters
	decodedURL, err := url.QueryUnescape(rawURL)
	if err != nil {
		http.Error(w, "Invalid URL encoding", http.StatusBadRequest)
		return
	}

	// Create context
	// Set up browser options
	var opts []chromedp.ExecAllocatorOption
	if runtime.GOOS == "windows" {
		// Windows 使用 Edge
		opts = append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.ExecPath("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"),
		)
	} else {
		// Linux 使用 Chromium
		opts = append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.ExecPath("/usr/bin/chromium-browser"),
		)
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	var buf []byte
	var contentType string
	var filename string

	switch format {
	case "png":
		// Capture screenshot
		if err := chromedp.Run(ctx, captureScreenshot(decodedURL, &buf)); err != nil {
			http.Error(w, "Failed to capture screenshot", http.StatusInternalServerError)
			return
		}
		contentType = "image/png"
		filename = "screenshot.png"
	default:
		// Capture PDF
		if err := chromedp.Run(ctx, printToPDF(decodedURL, &buf)); err != nil {
			http.Error(w, "Failed to render PDF", http.StatusInternalServerError)
			return
		}
		contentType = "application/pdf"
		filename = "output.pdf"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	if _, err := io.Copy(w, bytes.NewReader(buf)); err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
		return
	}
}

func captureScreenshot(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(5 * time.Second), // Wait for page to load
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Scroll to bottom of page to ensure all content is loaded
			var height int64
			if err := chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight); document.body.scrollHeight`, &height).Do(ctx); err != nil {
				return err
			}
			chromedp.Sleep(2 * time.Second) // Wait for any lazy-loaded content
			// Capture full page screenshot with 100% quality
			err := chromedp.FullScreenshot(res, 100).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

func printToPDF(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(5 * time.Second), // Wait for page to load
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).   // A4 width in inches
				WithPaperHeight(11.69). // A4 height in inches
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
