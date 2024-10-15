package traefikHelper

import (
    "fmt"
    "os"
    "github.com/go-resty/resty/v2"
    "encoding/json"
    "regexp"
)

var (
    base_domain = os.Getenv("BASE_DOMAIN")
    traefik_auth = os.Getenv("TRAEFIK_AUTH")

)
type TraefikConfig struct {
        EntryPoints  []string `json:"entryPoints"`
        Middlewares  []string `json:"middlewares,omitempty"` // Поле может отсутствовать, используем `omitempty`
        Service      string   `json:"service"`
        Rule         string   `json:"rule"`
        Priority     int      `json:"priority"`
        Status       string   `json:"status"`
        Using        []string `json:"using"`
        Name         string   `json:"name"`
        Provider     string   `json:"provider"`
    }
var configs []TraefikConfig

func Get_http_routes() []string {
    routes := Request_to_traefik("/api/http/routers?per_page=999999999999999999")
    err := json.Unmarshal([]byte(routes), &configs)
    if err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
    }
    var uniqueDomains []string
    for _, host := range configs {
        re := regexp.MustCompile(`Host\(` + "`" + `([^` + "`" + `]*)` + "`" + `\)`)
        match := re.FindStringSubmatch(host.Rule)
        if len(match) > 1 {
            host := match[1]
            addUnique(&uniqueDomains, host)
        }
    }
    return uniqueDomains
}

func Request_to_traefik(route string) string {
    traefikApiUrl := fmt.Sprintf("https://traefik.%s%s", base_domain, route)
    client := resty.New()
    if traefik_auth == "basic"{
        resp, err := client.R().
            SetBasicAuth(os.Getenv("TRAEFIK_USER"), os.Getenv("TRAEFIK_PASSWORD")).
            Get(traefikApiUrl)
        if err != nil {
            // handle error
            fmt.Println("Error:", err)
        }
        responseBody := string(resp.Body())
        return responseBody
    }else{
        panic(fmt.Sprintf("Unsupported auth %s", traefik_auth))
    }
}
