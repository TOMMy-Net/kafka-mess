# kafka-mess

**Оглавление:**

* [Установка](#Установка)
* [API](#API)
* [Предложения](#Предложения)

Данные проект представляет собой сервис отправки и получения сообщений через Kafka

## Установка

Для утстановки потребуется [Docker](https://www.docker.com/)

Проект можно запустить введя команду:  `docker-compose up`

## API

1. (POST) **/api/message** json: {text:""} - Отправка сообщения на сервер и дальнейшая его обработка
2. (GET) **/** - Hello
3. (GET) **/api/message/stats**  - Получение статистики по сообщениям
4. (GET) **/api/message** - Получение сообщения

## Предложения

1. Добавить авторизацию по токену (jwt)
2. Доработать код взаимодействия с Kafka
