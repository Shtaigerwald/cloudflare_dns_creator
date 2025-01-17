package main

import (
    "cloudflare-dns-creator/traefikHelper"
    "fmt"
    "cloudflare-dns-creator/cloudflareHelper"
    "time"

)

func uniqueElements(list1, list2 []string) []string {
	// Создаем map для второго списка
	set := make(map[string]struct{})
	for _, item := range list2 {
		set[item] = struct{}{}
	}

	// Проходим по первому списку и добавляем только уникальные элементы
	var unique []string
	for _, item := range list1 {
		if _, found := set[item]; !found {
			unique = append(unique, item)
		}
	}

	return unique
}


func taskEvery200ms() {
	for {
        traefikHosts := traefikHelper.GetHttpRoutes()
        cloudflareHosts := cloudflareHelper.GetDNSRecordsName()
        uniqueHosts := uniqueElements(traefikHosts, cloudflareHosts)
        if len(uniqueHosts) != 0 {
            fmt.Println("Записи будут добавлены:", uniqueHosts)
            for _, uniqueHost := range uniqueHosts{
                fmt.Printf("Запись будет добавлена: %s\n", uniqueHost)
                cloudflareHelper.CreateDNSRecords(uniqueHost)
            }
        }
		time.Sleep(200 * time.Millisecond)
	}
}

func taskDaily() {
	for {
		now := time.Now()
		// Вычисляем время до следующего полуночи
		nextMidnight := now.Add(time.Hour * 24)
		nextMidnight = time.Date(nextMidnight.Year(), nextMidnight.Month(), nextMidnight.Day(), 0, 0, 0, 0, nextMidnight.Location())
		durationUntilNextMidnight := nextMidnight.Sub(now)
        dnsRecordsToBeDeleted := cloudflareHelper.FindHostsToBeDeleted()
        cloudflareHelper.DeleteDNSRecord(dnsRecordsToBeDeleted)
		time.Sleep(durationUntilNextMidnight)
	}
}

func main() {
    fmt.Println("Run...")
	// Запускаем первую горутину для выполнения задачи каждые 0.2 секунды
	go taskEvery200ms()
	// Запускаем вторую горутину для выполнения задачи раз в сутки
	go taskDaily()
	// Чтобы программа не завершилась сразу после запуска горутин,
	// ждем завершения выполнения через блокирующий канал
	select {}
}
