# 2248_Notebook

Примеры запросов:

- Получение записки по uuid:
`curl --location 'http://localhost:8080/notebook/get?noteId='`

- Создание новой записки:
`curl --location --request POST 'http://localhost:8080/notebook/post' --header 'Content-Type: application/json' --data '{"message": "Hello world!"}'`

- Изменение существующей записки:
`curl --location --request PUT 'http://localhost:8080/notebook/put' --header 'Content-Type: application/json' --data '{"id": "","message": "Hello again"}'`