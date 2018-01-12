# Gologistic

Este repositório contém um projeto desenvolvido em Golang :heart: para funcionar como Client para qualquer um que deseje integrar com a API de Logistica da Stone.

## Utilizando no seu projeto

### Setup

Basta baixar a biblioteca para o seu workspace utilizando o comando tradicional do go, ou qualquer gerenciador de dependencia de sua escolha.

```bash
go get -u github.com/stone-payments/gologistic
```

### Show me the code

As interfaces principais são as interfaces Client contidas em cada pacote que encapsula a versão da API. Essas interfaces agrupam serviços por contexto.

```go
import (
    logistic "github.com/stone-payments/gologistic/v1"
    "github.com/stone-payments/gologistic/v1/models"
    "github.com/stone-payments/gologistic/v1/enums"
)

client := logistic.New(url)
service := client.ServiceOrder()
serviceOrders := service.List(models.ServiceOrderFilters{
    Status: enums.StatusClosed,
})

```

É possível registrar interceptors que processem requests e responses. Basta implementar a interface Interceptor do pacote http.

```go
import (
    "github.com/stone-payments/gologistic/errors"
    "github.com/stone-payments/gologistic/http"
)

func init() {
    http.Interceptors = append(http.Interceptors,
        http.NewInterceptor(func(req http.Request) http.Request {
            println("on request event")
            return req
        }, func(resp http.Response) (http.Response, errors.Error) {
            println("on response event")
            return resp, nil
        }))
}

```

Ou substituir o http requester padrão
```go
import (
    "github.com/stone-payments/gologistic/http"
)

func init() {
    http.Requester = requester
}

```

## Feito com

[Go](https://golang.org/) 1.9

## Como contribuir?

Todo o processo necessário para você contribuir encontra-se em nosso [CONTRIBUTING](CONTRIBUTING.md).

## Versionamento

Por favor, consulte nosso [CHANGELOG](CHANGELOG.md).

## Autores

Build with :heart: by Dev RC Team!

## License

Copyright (C) 2018. Stone Co. All rights reserved.
