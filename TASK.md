# Тестове завдання

## Потрібно реалізувати REST API з двома ендпоінтами:
Завантаження *.csv файлу, парсинг його і збереження результатів парсингу в базу даних.
Фільтрація і вивантаження попередньо збережених даних в JSON форматі в респонсі.

### Вимоги до фільтрів:
- пошук по transaction_id
- пошук по terminal_id (можливість вказати декілька одночасно id)
- пошук по status (accepted/declined)
- пошук по payment_type (cash/card)
- пошук по date_post по періодам (from/to), наприклад: from 2022-08-12, to 2022-09-01 повинен повернути всі транзакції за вказаний період
- пошук по частково вказаному payment_narrative

## Технічні вимоги:
- База даних повинна бути реляційна: MySQL, PostgrSQL, тощо
- Документація API повинна бути присутня (Swagger, OpenAPI чи просто в README.md)
- Документація до запуску і використання проекту (в README.md)
- Використання docker та docker-compose
- Можна використовувати будь-які бібліотеки чи фреймворки доступні в опенсорсі.

## Буде перевагою:
- Юніт та/або інтеграційні тести
- Передбачити можливість завантажувати файл великого розміру (від 1гб) при умові, що ресурс виданий сервісу буде обмежений в 100мб ОЗУ
- Зробити третій ендпоінт, котрий буде вивантажувати дані не в JSON, а у вигляді CSV файлу.
