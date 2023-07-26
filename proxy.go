package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

func ScrapProxy() []string {
	output := []string{}
	c := colly.NewCollector(colly.Async(true))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scarping Proxies from : ", r.URL)
	})

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr.prx_tr", func(_ int, h *colly.HTMLElement) {
			ip := h.ChildText("td.t_ip")
			port := h.ChildText("td.t_port")
			// Type := h.ChildText("td.t_type")

			output = append(output, ip+":"+port)
		})
	})

	c.Visit("https://proxylist.to/")
	c.Wait()
	return output
}



func main() {
	//creating the proxyURL
	fmt.Println("Hey , wasup...")
	fmt.Println("scraping started....")

	fmt.Println(" -------------------------------------------------- ")
	fmt.Printf("Starting time : %v\n", time.Now().Format(time.RFC850))

	proxies := ScrapProxy()
	for _, pro := range proxies {
		proxyStr := "http://" + pro
		proxyURL, _ := url.Parse(proxyStr)

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
		timeout := time.Duration(30 * time.Second)
		client := &http.Client{
			Transport: transport,
			Timeout:   timeout,
		}

		response, err := client.Get("https://api.ipify.org?format=json")
		if err != nil {
			// fmt.Println("Response is False : %s", err.Error())
			fmt.Printf("[-] Proxy not found : %s\n", proxyStr)
			continue
		}
		// data_bytes, _ := ioutil.ReadAll(response.Body)
		if response.StatusCode == 200 {
			fmt.Printf("[+] Proxy Found : %s\n", proxyStr)
			// fmt.Println(string(data_bytes))
		}
		time.Sleep(3 * time.Second)

	}
	fmt.Printf("Ending time : %v", time.Now().Format(time.RFC850))

}
