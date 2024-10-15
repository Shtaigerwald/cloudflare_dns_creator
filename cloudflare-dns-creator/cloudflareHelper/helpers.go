package cloudflareHelper

import (
    "context"
    "fmt"
    "os"
    "github.com/cloudflare/cloudflare-go"
    "log"
    "time"
    "strings"
)

var (
	email, emailExists = os.LookupEnv("CF_API_EMAIL")
    key, keyExists = os.LookupEnv("CF_API_KEY")
    api *cloudflare.API
    err error
    clouflareDomain, domainExists = os.LookupEnv("BASE_DOMAIN")
    domainIP, ipExists = os.LookupEnv("DOMAIN_IP")
)

func init() {
    api, err = cloudflare.New(key, email)
    if err != nil {
        log.Fatal(err)
    }
}

func GetDNSRecordsName() []string {
    listDNSRecords := GetDNSRecords()
    var listHostsName []string
    for _, dnsRecord := range listDNSRecords {
        listHostsName = append(listHostsName, dnsRecord.Name)
    }
    return listHostsName
}

func GetDNSRecords() []cloudflare.DNSRecord {
    zoneID := FindZoneID(clouflareDomain)
    //zoneID := "3edda56f8757a1f5203ecfb593834aeb"
    listDNSRecords, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
    if err != nil {
        log.Fatalf("Failed to list zones: %v", err)
    }
    return listDNSRecords
}

func FindZoneID(domainName string) string {
    zoneID, err := api.ZoneIDByName(domainName)
    if err != nil {
        log.Fatalf("Failed find zone by name:%s, err: %v", domainName, err)
    }
    return zoneID
}

func FindHostsToBeDeleted() []string {
    listDnsRecords := GetDNSRecords()
    currentTime := time.Now()
    var listHostsTobeDeleted []string
    for _, dnsRecord := range listDnsRecords {
        if strings.Contains(dnsRecord.Name, "mr-") && strings.Contains(dnsRecord.Name, "testing"){
            duration := currentTime.Sub(dnsRecord.CreatedOn)
            if duration.Hours() > 15*24 {
                fmt.Printf("Will be deleted: \nDNS_RECORD_ID:%s, DNS_RECORD_NAME: %s\n", dnsRecord.ID, dnsRecord.Name)
                listHostsTobeDeleted = append(listHostsTobeDeleted, dnsRecord.ID)
            }
        }
    }
    if len(listHostsTobeDeleted) == 0 {
        fmt.Println("Cloudflare_delete_old_records: Nothing to be deleted")
    }
    return listHostsTobeDeleted
}

func DeleteDNSRecord(listDnsRecordID []string) {
    ctx := context.Background()
    zoneID := FindZoneID(clouflareDomain)
    resource := &cloudflare.ResourceContainer{
        Level:      cloudflare.ZoneRouteLevel,
        Identifier: zoneID,
    }
    for _, dnsID := range listDnsRecordID {
        err := api.DeleteDNSRecord(ctx, resource, dnsID)
        if err != nil{
            log.Printf("err: %v", err)
        }
    }
}

func CreateDNSRecords(recordName string) {
    zoneID := FindZoneID(clouflareDomain)
    Btrue := true
    _, err := api.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID),cloudflare.CreateDNSRecordParams{
        Type:    "A",
        Name:    recordName,
        Content: fmt.Sprint(domainIP),
        Proxied: &Btrue,
        })
    if err != nil {
        if strings.Contains(err.Error(), "Record already exists") {
            fmt.Printf("Record already exists: %v\n", clouflareDomain)
        }
        log.Fatal(err)
    }else {
        fmt.Printf("Record created: %v\n", recordName)
    }
}
