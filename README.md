**Errors** - обеспечивает простую обработку ошибок, а так же включает в себя возможность помечать их конретными кодами
для передачи их в ответах на запросы к серверу.

**Базовые функции:**

```go
// New(text string, errCode ...uint32) error
err := errors.New("foo")
// or, econst.InvalidPasswordValue = 3226689784
err := errors.New("foo", econst.InvalidPasswordValue)
```

```go
// Errorf(format string, args ...interface{}) error
err := errors.Errorf("foo %s", "bar")
```

```go
// Wrap(err error, text string) error
err := errors.Wrap(err, "bar")
```

```go
// Wrapf(err error, format string, args ...interface{}) error
err := errors.Wrapf(err, "bar %s", "foo")
```

```go
// Cause(err error) error
err := errors.Cause(err)
```

**Дополнения:**

> SetSeparator заменяет разделитель по умолчанию ": " на пользовательский.

```go
// SetSeparator(sep string)
err := errors.SetSeparator("-> ")
```

**Особенности**

> Mark - помечает ошибку кодом, который можно вывести в ответах на запросы
```go
// Mark(err error, code uint32) error
err := errors.Mark(err, econst.InvalidEmailValue /* 366256089 */)
```

> Markf - помечает ошибку кодом, который можно вывести в ответах на запросы, а так же добавляет контекст к ошибке
```go
// Markf(err error, code uint32) error
err := errors.Markf(err, econst.InvalidEmailValue /* 366256089 */, "foo %s", "bar")
```


> GetMark - выводит последний установленный код из ошибки
```go
// GetMark(err interface{}) (uint32, bool)
code, ok := errors.Mark(err)
```

> Response - формирует json ответ следующего формата:

* *если есть ошибка*
```json
{
    "data": null,
    "error": {
        "text": "foo",
        "code": 10000000
    }
}
```

* *если ошибка отсутствует*
```json
{
    "data": {...},
    "error": null
}
```

```go
// Response(data interface{}, err error, httpStatuses map[uint32]map[string]interface{}) (int, string, []byte)
httpStatus, contentType, body := errors.Response(data, err, econst.HttpResponse)
```

> Json - функция, которую можно использовать в http handlers
```go
// Json(w http.ResponseWriter, data interface{}, err error, httpStatuses map[uint32]map[string]interface{})
defer func() { errors.Json(w, data, err, econst.HttpResponse) }()
```

**Пример:**

```go
type user struct {
  Username string `json:"username"`
}

func handleGetUserNameByID(c echo.Context) error {
  var (
    err error
    id int
    data *user
  )

  /*
    если есть ошибка, ответит на запрос с отличным от 200 статусом и следующим содержимым:
    {
        "data": null,
        "error": {
            "text": "foo",
            "code": 10000000
        }
    }

    если ошибки нет, ответит 200 статусом и следующим содержимым:
    {
      "data": {
        "username": "Foobar"
      },
      "error": null
    }
  */
  defer func() { c.Blob(errors.Response(data, err, econst.HttpResponse)) }()

  if id, err = getUserID(c); err != nil {
    return err
  }

  if data, err = getUserData(id); err != nil {
    return err
  }

  return nil
}

func getUserID(c echo.Context) (int, error) {
  id, err := strconv.Atoi(c.FormValue("id"))
  if err != nil {
    return 0, errors.Mark(err, econst.InvalidUserIDValue)
  }

  return id, nil
}

func getUserData(id int) (*user, error) {
    if id != 1 {
        return nil, errors.New("user not found", econst.UserNotFound)
    }

    return &user{"Foobar"}, nil
}
```

**Генератор ошибок:**
**egen** - генерирует **econst.go** и файл **errors.response.json** для **frontend** части на основе файла **errors.json** (пример в **errors.example.json**)

##### Установка:
```bash
$ go install github.com/marval-devsystems/errors/cmd/egen@latest
```
##### Использование:
```bash
// '.' -> путь до файла errors.json, там же и будут созданы файлы econst.go и errors.response.json
$ egen .
```