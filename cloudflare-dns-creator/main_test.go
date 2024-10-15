package main

import (
    "cloudflare-dns-creator/cloudflareHelper"
	"testing"
	"fmt"
    "os"
    "cloudflare-dns-creator/traefikHelper"
)
var(
    email, emailExists = os.LookupEnv("CF_API_EMAIL")
    key, keyExists = os.LookupEnv("CF_API_KEY")
    clouflare_domain, domainExists = os.LookupEnv("BASE_DOMAIN")
    domain_ip, ipExists = os.LookupEnv("DOMAIN_IP")
)

// Тестируем функцию Get_dns_records_name
func TestGet_dns_records_name_cloudflare(t *testing.T) {
    fmt.Printf("Get dns records from domain '%s':\n", clouflare_domain)
	lists_dns_records_name := cloudflareHelper.Get_dns_records_name()
    for _, name := range lists_dns_records_name{
        fmt.Println(name)
    }
}

func TestFind_zone_id_clouflare(t *testing.T) {
    fmt.Printf("Domain: '%s' - ", clouflare_domain)
	zoneID := cloudflareHelper.Find_zone_id(clouflare_domain)
    fmt.Printf("Zone ID:%s\n",zoneID)
}

func TestFind_list_hosts_tobe_deleted_cloudflare(t *testing.T){
    fmt.Printf("Get dns records from domain '%s':\n", clouflare_domain)
    lists_dns_records_tobe_deleted := cloudflareHelper.Find_list_hosts_tobe_deleted()
    for _, name := range lists_dns_records_tobe_deleted{
        fmt.Println(name)
    }
}

func TestRequest_to_traefik(t *testing.T){
    fmt.Println("Get json from traefik:")
    jsonfromtraefik := traefikHelper.Request_to_traefik("/api/http/routers?per_page=999999999999999999")
    fmt.Println(jsonfromtraefik)
}

func TestGet_http_routes_traefik(t *testing.T){
    fmt.Println("Get routes from Traefik:")
    list_traefik_routes := traefikHelper.Get_http_routes()
    for _, name := range list_traefik_routes{
      fmt.Println(name)
    }
}
