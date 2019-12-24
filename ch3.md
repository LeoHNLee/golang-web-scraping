# Chapter3: Web Scraping Etiquette

1. robots.txt

    - robots.txt parser

    ```shell
    go get github.com/temoto/robotstxt

    go get github.com/gregjones/httpcache
    go get github.com/peterbourgon/diskv
    ```

1. user-agent

    - example: Mozilla/5.0 (X11; Linux x86_64; rv:57.0) Gecko/20000101 Firefox/57.0
    - cURL: curl/@@
    - go: Go-http-client/@@
    - Java: Apache-HttpClient/@@
    - Googlebot(for image): Googlebot-Image/@@

1. How to use caching

    ```
    // Response
    HTTP/1.1 200 OK
    Accept-Ranges: bytes
    Cache-Control: max-age=604800
    Content-Type: text/html; charset=UTF-8
    Date: Mon, 29 Oct 2018 13:31:23 GMT
    Etag: "1541025663"
    Expires: Mon, 05 Nov 2018 13:31:23 GMT
    Last-Modified: Fri, 09 Aug 2013 23:54:35 GMT
    Server: ECS (dca/53DB)
    ...
    ```

1. example

    ```go
    package main

    import (
        "fmt"
        "net/http"
        "time"
    )

    func main() {
        var lastRequestTime time.Time

        maximumNumberOfRequests := 5

        pageDelay := 5 * time.Second

        for i := 0; i < maximumNumberOfRequests; i++ {
            elapsedTime := time.Now().Sub(lastRequestTime)
            fmt.Printf("Elapsed Time: %.2f (s)\n", elapsedTime.Seconds())
            if elapsedTime < pageDelay {
                var timeDiff time.duration = pageDelay - elapsedTime
                fmt.Printf("Sleeping for %.2f (s)\n)", timeDiff.Seconds())
                time.Sleep(pageDelay - elapsedTime)
            }

            println("GET example.com/index.html")
            _, err := http.Get("http://www.example.com/index.html")
            if err != nil {
                panic(err)
            }
            lastRequestTime = time.Now()
        }
    }
    ```

    ```go
    package main
    import (
        "nte/http"

        "github.com/temoto/robotstxt"
    )
    func main() {
        resp, err := http.Get("http://www.packtpub.com/robots.txt")
        if err != nil {
            panic(err)
        }
        data, err := robotstxt.FromResponse(resp)
        if err != nil {
            panic(err)
        }

        // Look for the definition in the robots.txt file that matches the default Go User-Agent string
        grp := data.FindGroup("Go-http-client/1.1")
        if grp != nil {
            testUrls := []string{
                // These paths are all permissable
                "/all",
                "/all?search=Go",
                "/bundles",

                // These paths are not
                "/contact/",
                "/search/",
                "/user/password/",
            }
            for _, url := range testUrls {
                print("checking "+ url+"...")

                // Test the path against the User-Agent group
                if grp.Test(url) == true {
                    println("OK")
                } else {
                    println("X")
                }
            }
        }
    }
    ```

    ```go
    // multiple web-site
    lastRequestMap := map[string]time.Time{
        "example.com": time.Time{},
        "packtpub.com": time.Time{},
    }
    ```

    ```go
    package main

    import (
        "io/ioutil"

        "github.com/gregjones/httpcache"
        "github.com/gregjones/httpcache/diskcache"
    )

    func main() {
        storage := diskcache.New(".cache")
        
    }
    ```