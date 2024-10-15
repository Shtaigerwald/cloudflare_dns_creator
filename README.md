# Cloudflare dns creator

## Описание
Этот проект предназначен для взаимодействия с Traefik и Cloudflare API для добавления и удаления DNS записей. Приложение позволяет управлять HTTP маршрутами Traefik и использовать Cloudflare для управления DNS записями.

## Сборка и запуск
### Требования
* Docker
* Docker Compose (если нужно)
* Доступ к Traefik API
* Учетные данные Cloudflare
* Переменные окружения

## ENV_VARIABLES


`CLOUDFLARE_DOMAIN`-  Основной домен, к которому будет привязан Traefik. * Например, если указать example.com, Traefik будет доступен по адресу
traefik.example.com.

`IP_CONNECTED_TO_DOMEN`-  IP-адрес, к которому будут привязываться DNS-записи.

`TRAEFIK_AUTH`-  Способ авторизации в Traefik (например, BasicAuth).

`TRAEFIK_USER`-  Логин для доступа к Traefik API.

`TRAEFIK_PASSWORD`-  Пароль для доступа к Traefik API.

`CLOUDFLARE_API_EMAIL`-  Email, использующийся для аутентификации в Cloudflare.

`CLOUDFLARE_API_KEY`-  API ключ для доступа к Cloudflare.


## BuildСборка 
bash
```
docker build -t my-go-project .
```
## Run
```
docker run -d \
  -e BASE_DOMAIN="example.com" \
  -e DOMAIN_IP="192.168.1.1" \
  -e TRAEFIK_AUTH="basic" \
  -e TRAEFIK_USER="traefikUser" \
  -e TRAEFIK_PASSWORD="traefikPassword" \
  -e CF_API_EMAIL="your-email@example.com" \
  -e CF_API_KEY="your-cloudflare-api-key" \
  my-go-project
```


## Tests
### Run all tests
```
go test
```
### Run test TestGet_http_routes_traefik
```
go test -run TestGet_http_routes_traefik
```

## Structure

`main.go` - Основной файл программы.

`main_test.go` - Файл с тестами.

`Dockerfile` - Файл для сборки Docker контейнера.
