# Тестовое задание для ИнфоТеКС

## Зависимости
Перед запуском проекта убедитесь, что у вас установлены следующие компоненты:
- Docker
- Docker Compose

## Развёртывание

### Запуск Swagger UI
Для запуска Swagger UI выполните следующие шаги:
```
docker-compose up swagger-ui
```
Перейдите по адресу [http://localhost:8080/swagger](http://localhost:8080/swagger) для просмотра Swagger UI.

### Запуск сервера и базы данных
Для разворачивания сервера и базы данных выполните следующие шаги:
```
docker-compose up postgres server
```
Обращаться к серверу можно по адресу http://localhost:3000/.

Теперь ваше приложение готово к использованию.