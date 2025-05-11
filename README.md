# url-shortify

### Этот проект сокращать огромные ссылки, делая их маленькими. Также он собирает аналитику по посещаемости ссылок.

## Как пользоваться?

Пишем эти команды в терминале:

```sh
git clone https://github.com/TheJustRusik/url-shortify
cd url-shortify
docker compose up --d
```

И у вас поднимется 3 сервиса:
1. shortener - он занимается сокращением и редиректом ссылок
2. analytics - собирает и отображает аналитику по ссылкам
3. db - база данных PostgreSQL

Swagger первого микросервиса (http://localhost:8080/swagger/) в нём есть 2 эндпоинта:
1. Создаёт сокращенную ссылку
2. Делает редирект с сокращенной ссылки на основную (пример: http://localhost:8080/Ee2m6X)

Аналитика второго микросервиса:
http://localhost:8081/stats/Ee2m6X -> покажет json в котором видно кто, когда, откуда, куда перешёл

# Демо
## Создать ссылку:
```sh
curl -X 'POST' \
  'https://url-short.kenuki.dev/shorten' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "url": "https://github.com/TheJustRusik/url-shortify/blob/main/shortener-service/cmd/main.go"
}'
```
В ответ получите 6 символов
## Перейти по ссылке:
https://url-short.kenuki.dev/f65CAt просто в браузере заходите сюда
## Аналитика
https://url-short.kenuki.dev/stats/f65CAt тоже открываем в браузере
