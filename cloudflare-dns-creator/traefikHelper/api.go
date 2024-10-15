package traefikHelper

import (
    "fmt"
    "os"
    "github.com/go-resty/resty/v2"
    "encoding/json"
    "regexp"
)

var (
    baseDomain  = os.Getenv("BASE_DOMAIN")
    traefikAuth = os.Getenv("TRAEFIK_AUTH")

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

func GetHttpRoutes() []string {
    routes := RequestToTraefik("/api/http/routers?per_page=999999999999999999")
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
            addUniqueDomainInList(&uniqueDomains, host)
        }
    }
    return uniqueDomains
}

func RequestToTraefik(route string) string {
    traefikApiUrl := fmt.Sprintf("https://traefik.%s%s", baseDomain, route)
    client := resty.New()
    if traefikAuth == "basic"{
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
        panic(fmt.Sprintf("Unsupported auth %s", traefikAuth))
    }
}
