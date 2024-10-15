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
    clouflare_domain, domainExists = os.LookupEnv("BASE_DOMAIN")
    domain_ip, ipExists = os.LookupEnv("DOMAIN_IP")
)

func init() {
    api, err = cloudflare.New(key, email)
    if err != nil {
        log.Fatal(err)
    }
}

func Get_dns_records_name() []string {
    list_dns_records := Get_dns_records()
    var list_hosts_name []string
    for _, dns_record := range list_dns_records {
        list_hosts_name = append(list_hosts_name, dns_record.Name)
    }
    return list_hosts_name
}

func Get_dns_records() []cloudflare.DNSRecord {
    ctx := context.Background()
    zoneID := Find_zone_id(clouflare_domain)
    //zoneID := "3edda56f8757a1f5203ecfb593834aeb"
    list_dns_records, _, err := api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
    if err != nil {
        log.Fatalf("Failed to list zones: %v", err)
    }
    return list_dns_records
}

func Find_zone_id(domain_name string) string{
    zoneID, err := api.ZoneIDByName(domain_name)
    if err != nil {
        log.Fatalf("Failed find zone by name:%s, err: %v", domain_name, err)
    }
    return zoneID
}

func Find_list_hosts_tobe_deleted() []string {
    list_dns_records := Get_dns_records()
    currentTime := time.Now()
    var list_hosts_tobe_deleted []string
    for _, dns_record := range list_dns_records {
        if strings.Contains(dns_record.Name, "mr-") && strings.Contains(dns_record.Name, "testing"){
            duration := currentTime.Sub(dns_record.CreatedOn)
            if duration.Hours() > 15*24 {
                fmt.Printf("Will be deleted: \nDNS_RECORD_ID:%s, DNS_RECORD_NAME: %s\n", dns_record.ID, dns_record.Name)
                list_hosts_tobe_deleted = append(list_hosts_tobe_deleted, dns_record.ID)
            }
        }
    }
    if len(list_hosts_tobe_deleted) == 0 {
        fmt.Println("Cloudflare_delete_old_records: Nothing to be deleted")
    }
    return list_hosts_tobe_deleted
}

func Delete_dns_record(list_dns_record_id []string) {
    ctx := context.Background()
    zoneID := Find_zone_id(clouflare_domain)
    resource := &cloudflare.ResourceContainer{
        Level:      cloudflare.ZoneRouteLevel,
        Identifier: zoneID,
    }
    for _, dns_id := range list_dns_record_id{
        err := api.DeleteDNSRecord(ctx, resource, dns_id)
        if err != nil{
            log.Printf("err: %v", err)
        }
    }
}

func Create_dns_records(record_name string){
    zoneID := Find_zone_id(clouflare_domain)
    Btrue := true
    _, err := api.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID),cloudflare.CreateDNSRecordParams{
        Type:     "A",
        Name:     record_name,
        Content:  fmt.Sprint(domain_ip),
        Proxied: &Btrue,
        })
    if err != nil {
        if strings.Contains(err.Error(), "Record already exists") {
            fmt.Printf("Record already exists: %v\n", clouflare_domain)
        }
        log.Fatal(err)
    }else {
        fmt.Printf("Record created: %v\n", record_name)
    }
}
