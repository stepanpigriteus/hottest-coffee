не хватает json ответов' 

PUT
localhost:8200/menu/ нет ответа
localhost:8200/menu/id Undefined Error, please check your method or endpoint correctness привести к json
localhost:8200/menu/id несуществующие id
localhost:8200/menu/  Undefined Error, please check your method or endpoint correctness привести к json
localhost:8200/menu/id успешно нет ответа
localhost:8200/orders/id успешно нет ответа
localhost:8200/orders/id неудачно нет ответа


POST
localhost:8200/menu/ успешно
localhost:8200/orders/ с некорректными данными  500

Delete
localhost:8200/menu/id успешное удаление
localhost:8200/orders/id успешно нет ответа
localhost:8200/orders/id не найден - нет ответа