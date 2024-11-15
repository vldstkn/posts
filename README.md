## О проекте
Это простое CRUD-приложение, в нем нет ничего сложного. Пользователь может пройти регистрацию и создавать или получать посты. 
Суть проекта заключалась в ознакомлении с Golang и его стандартными библиотеками.

## Что использовалось
Я использовал роутер от [**chi**](https://github.com/go-chi/chi), он показался мне наиболее привлекательным, потому что хорошо совместим со 
стандартной библиотекой "net/http". По той же причине выбрал библиотеку [**sqlx**](https://github.com/jmoiron/sqlx), для запросов к базе данных.
В качестве базы данных выбрал **Postgres**. Для валидации данных использовал [**validator**](https://github.com/go-playground/validator)

## Подробнее о проекте
В проекте реализована авторизация с использованием access и refresh токенами.
Особое внимание старался уделить принципам чистой архитектуры, надеюсь, у меня получилось.
Использовал утилиту [**air**](https://github.com/air-verse/air) для live-reload.
Миграции проводил с помощью библиотеки [**goose**](https://github.com/pressly/goose).
Приложение не покрывал никакими тестами, поскольку не хотелось долго останавливаться на таком небольшом и простом проекте, 
ведь его задача была совсем другой.

## Структура папок
1. cmd - файлы с точками входа в приложение
2. configs - конфигурационные файлы (local.yaml)
3. internal - вся внутренняя логика приложения
    * app - файлы для сборки приложения
    * config - обработка конфигов (пункт 2)
    * di - интерфейсы сервисов и репозиториев
    * domain - модели приложения, бизнес-сущности (user, post)
    * repository - слой репозитория
    * service - сервисный слой
    * transport - обработка внешний запросов (http)
4. migrations - файл миграции базы данных
5. pkg - вся экспортируемая/переиспользуемая логика
    * db - логика подключения к базе данных
    * jwt - парсинг и создание jwt
    * logger - настройка логгера
    * req - обработка тела запроса
    * res - создание ответа
