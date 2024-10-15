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
    clouflareDomain, domainExists = os.LookupEnv("BASE_DOMAIN")
    domainIP, ipExists = os.LookupEnv("DOMAIN_IP")
)

// Тестируем функцию GetDNSRecordsName
func TestGetDNSRecordsName_cloudflare(t *testing.T) {
    fmt.Printf("Get dns records from domain '%s':\n", clouflareDomain)
	lists_dns_records_name := cloudflareHelper.GetDNSRecordsName()
    for _, name := range lists_dns_records_name{
        fmt.Println(name)
    }
}

func TestFindZoneID_clouflare(t *testing.T) {
    fmt.Printf("Domain: '%s' - ", clouflareDomain)
	zoneID := cloudflareHelper.FindZoneID(clouflareDomain)
    fmt.Printf("Zone ID:%s\n",zoneID)
}

func TestFindHostsToBeDeleted_cloudflare(t *testing.T) {
    fmt.Printf("Get dns records from domain '%s':\n", clouflareDomain)
    lists_dns_records_tobe_deleted := cloudflareHelper.FindHostsToBeDeleted()
    for _, name := range lists_dns_records_tobe_deleted{
        fmt.Println(name)
    }
}

func TestRequestToTraefik(t *testing.T) {
    fmt.Println("Get json from traefik:")
    jsonfromtraefik := traefikHelper.RequestToTraefik("/api/http/routers?per_page=999999999999999999")
    fmt.Println(jsonfromtraefik)
}

func TestGetHttpRoutes_traefik(t *testing.T) {
    fmt.Println("Get routes from Traefik:")
    list_traefik_routes := traefikHelper.GetHttpRoutes()
    for _, name := range list_traefik_routes{
      fmt.Println(name)
    }
}
