package traefikHelper

import (
    "strings"
)


func extractDomain(input string) string {
    // Разделяем строку по точкам
    parts := strings.Split(input, ".")
    // Проверяем, что домен содержит минимум 4 части
    if strings.Contains(input, "mr-") && strings.Contains(input, "testing"){
        if len(parts) >= 5 {
            // Возвращаем последние четыре части домена
            return strings.Join(parts[len(parts)-5:], ".")
        }
    }else{
        if len(parts) >= 4 {
            // Возвращаем последние четыре части домена
            return strings.Join(parts[len(parts)-4:], ".")
        }
    }
    // Если домен меньше 4 частей, возвращаем оригинальный домен
    return input
}
func addUnique(domainList *[]string, domain string) {
    // Извлекаем основной домен (последние 4 части)
    baseDomain := extractDomain(domain)

    // Генерируем шаблон с * для поддоменов
    wildcardDomain := "*." + baseDomain

    // Проверяем, существует ли уже один из доменов в списке
    if !contains(*domainList, baseDomain) && !contains(*domainList, wildcardDomain) {
        // Если домена нет, добавляем оба варианта в список
        *domainList = append(*domainList, baseDomain)
        if strings.Contains(baseDomain, "development") || strings.Contains(baseDomain, "testing"){
            *domainList = append(*domainList, wildcardDomain)
        }
    }
}
func contains(s []string, str string) bool {
    for _, v := range s {
        if v == str {
            return true
        }
    }
    return false
}
