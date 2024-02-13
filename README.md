# Менеджер паролей GophKeeper

GophKeeper представляет собой клиент-серверную систему, 
позволяющую пользователю надёжно и безопасно хранить логины, пароли, 
текстовые и бинарные данные.

## Особенности клиентской части приложения
- Клиент поддерживает следующую бизнес-логику:
  - аутентификация и авторизация пользователей на удалённом сервере;
  - доступ к приватным данным по запросу;
  - отображение информации о версии и дате сборки бинарного файла клиентского приложения.
- Клиентское приложение предназначено для запуска через консоль (CLI).
- Приложение представляет собой "тонкий клиент" и не хранит приватные данные пользователя в локальном хранилище.
- Шифрование и расшифровка данных осуществляются на стороне клиентского приложения.
- Приложение хранит свои настройки, а также данные для авторизации и шифрования в конфигурационном файле **gophkeeper.yaml**.

## Особенности серверной части приложения
- Сервер поддерживает следующую бизнес-логику:
    - регистрация, аутентификация и авторизация пользователей;
    - хранение приватных данных в зашифрованном виде, а также их изменение и удаление;
    - синхронизация данных между несколькими авторизованными клиентами одного владельца;
    - передача приватных данных владельцу по запросу.
- В качестве СУБД использется **PostgreSQL**.
- Сервер получает и возвращает клиенту данные в зашифрованном виде. Ключ шифрования не хранится на сервере.
- Взаимодействие с сервером осуществляется по бинарному протоколу **gRPC**.

## Абстрактная схема взаимодействия с системой
Ниже описаны базовые сценарии взаимодействия пользователя с системой.
### Для нового пользователя:
1. Пользователь получает клиент под необходимую ему платформу.
1. Пользователь проходит процедуру первичной регистрации (команда `sign up`).
1. Пользователь добавляет в клиент новые данные (команды `password add`, `card add`, `text add`, `binary add`).
1. Клиент синхронизирует данные с сервером.
### Для существующего пользователя:
1. Пользователь получает клиент под необходимую ему платформу.
1. Пользователь проходит процедуру аутентификации (команда `sign in`).
1. Клиент синхронизирует данные с сервером (команда `record list`).
1. Пользователь запрашивает данные (команды `password show`, `card show`, `text show`, `binary show`).
1. Клиент отображает данные для пользователя.

## Список команд
`-a`, `--address` - глобальный флаг с адресом сервера.
1. `version` - отображение версии, даты и времени сборки клиентского приложения.
---

2. `sign up` - регистрация нового пользователя. Поддерживает следующие обязательные флаги:
   - `-u`, `--user` - имя учётной записи пользователя;
   - `-p`, `--password` - пароль пользователя.


3. `sign in` - вход в учётную запись существующего пользователя. Поддерживает следующие обязательные флаги:
    - `-u`, `--user` - имя учётной записи пользователя;
    - `-p`, `--password` - пароль пользователя.


4. `sign out` - выход из учётной записи авторизованного пользователя.
---

5. `record list` - вывод списка сохранённых записей с зашифрованными данными.


6. `record show` - отображение подробной информации о записи с зашифрованными данными. Поддерживает единственный обязательный флаг:
   - `-i`, `--id` - идентификатор запрашиваемой записи.


7. `record change` - изменение названия или примечания записи с зашифрованными данными. 
Для команды должен быть указан хотя бы один из необязательных флагов. 
Команда поддерживает следующие флаги:
    - `-i`, `--id` (обязательный) - идентификатор запрашиваемой записи;
    - `-n`, `--name` (необязательный) - новое название записи с приватными данными;
    - `-m`, `--metadata` (необязательный) - новое примечание к записи с приватными данными;


8. `record delete` - удаление записи с зашифрованными данными. Поддерживает единственный обязательный флаг:
    - `-i`, `--id` - идентификатор запрашиваемой записи.
---

9. `password add` - добавление записи с логином и паролем. Поддерживает следующие флаги:
   - `-l`, `--login` (обязательный) - логин для сохранения;
   - `-p`, `--password` (обязательный) - пароль для сохранения;
   - `-n`, `--name` (обязательный) - название для идентификации сохраняемых приватных данных;
   - `-m`, `--metadata` (необязательный) - примечание к сохраняемым приватным данным.


10. `password show` - отображение записи с логином и паролем. Поддерживает единственный обязательный флаг:
    - `-i`, `--id` - идентификатор запрашиваемой записи.


11. `password change` - изменение записи с логином и паролем. 
Для команды должен быть указан хотя бы один из необязательных флагов. 
Поддерживает следующие флаги:
    - `-i`, `--id` (обязательный) - идентификатор изменяемой записи;
    - `-l`, `--login` (необязательный) - логин для сохранения;
    - `-p`, `--password` (необязательный) - пароль для сохранения;
---

12. `card add` - добавление записи с данными банковской карты. Поддерживает следующие флаги:
    - `--number` (обязательный) - номер банковской карты;
    - `--holder` (обязательный) - имя владельца банковской карты;
    - `--year` (обязательный) - год истечения срока действия карты;
    - `--month` (обязательный) - месяц истечения срока действия карты;
    - `--cvc` (обязательный) - CVC/CVV карты;
    - `-n`, `--name` (обязательный) - название для идентификации сохраняемых приватных данных;
    - `-m`, `--metadata` (необязательный) - примечание к сохраняемым приватным данным.


13. `card show` - отображение записи с данными банковской карты. Поддерживает единственный обязательный флаг:
    - `-i`, `--id` - идентификатор запрашиваемой записи.


14. `card change` - изменение записи с данными банковской карты.
    Для команды должен быть указан хотя бы один из необязательных флагов.
    Поддерживает следующие флаги:
    - `-i`, `--id` (обязательный) - идентификатор изменяемой записи;
    - `--number` (необязательный) - номер банковской карты;
    - `--holder` (необязательный) - имя владельца банковской карты;
    - `--year` (необязательный) - год истечения срока действия карты;
    - `--month` (необязательный) - месяц истечения срока действия карты;
    - `--cvc` (необязательный) - CVC/CVV карты.
---

15. `text add` - добавление записи с текстовыми данными. 
Текстовые данные передаются вводом текста в консоль после запуска приложения. 
Или из файла указанием `<` в консоли между вызовом программы и именем файла.
Команда поддерживает следующие флаги:
    - `-n`, `--name` (обязательный) - название для идентификации сохраняемых приватных данных;
    - `-m`, `--metadata` (необязательный) - примечание к сохраняемым приватным данным.


16. `text show` - отображение записи с текстовыми данными. 
Текстовые данные выводятся в консоль или в файл. 
Для вывода в файл необходимо указать `>` в консоли между 
вызовом программы и именем файла. 
Команда поддерживает единственный обязательный флаг:
    - `-i`, `--id` - идентификатор запрашиваемой записи.
    
17. `text change` - изменение записи с текстовыми данными.
Текстовые данные передаются вводом текста в консоль после запуска приложения.
Или из файла указанием `<` в консоли между вызовом программы и именем файла.
Команда поддерживает единственный обязательный флаг:
    - `-i`, `--id` - идентификатор запрашиваемой записи.
---

18. `binary add` - добавление записи с бинарными данными. 
Бинарные данные передаются из файла указанием `<` в консоли 
между вызовом программы и именем файла.
Команда поддерживает следующие флаги:
    - `-n`, `--name` (обязательный) - название для идентификации сохраняемых приватных данных;
    - `-m`, `--metadata` (необязательный) - примечание к сохраняемым приватным данным. 


19. `binary show` - отображение записи с бинарными данными.
Бинарные данные выводятся в файл.
Необходимо указать `>` в консоли между вызовом программы и именем файла.
Команда поддерживает единственный обязательный флаг:
    - `-i`, `--id` - идентификатор запрашиваемой записи.


20. `binary change` - изменение записи с бинарными данными.
Бинарные данные передаются из файла указанием `<` в консоли 
между вызовом программы и именем файла.
Команда поддерживает единственный обязательный флаг:
    - `-i`, `--id` - идентификатор запрашиваемой записи.