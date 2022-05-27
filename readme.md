# Bank transactions Api

API execução de transações bancárias

## Tecnologias utilizadas no projeto

- [Docker](https://www.docker.com/)
- [Golang (Versão 1.18)](https://www.php.net/)
- [Clean arch](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Instalação

Após baixar o projeto, execute

```bash
make up
```

## Rotas

No diretório **docs**, contém a collection para importação no API client [Postman](https://github.com/WenLopes/transaction-api/blob/master/docs/Postman%20-%20Transaction%20API.json).

### Reset
(POST) ```localhost:16091/reset```

```
curl --location --request POST 'localhost:16091/reset' \
--data-raw ''
```

### Balance
(GET) ```localhost:16091/balance?account_id=18```

```
curl --location --request GET 'localhost:16091/balance?account_id=18' \
--data-raw ''
```

### Event: Deposit
(POST) ```localhost:16091/event```

```
curl --location --request POST 'localhost:16091/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "deposit",
    "destination": "18",
    "amount": 1
}'
```

### Event: Withdraw
(POST) ```localhost:16091/event```

```
curl --location --request POST 'localhost:16091/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "withdraw",
    "origin": "18",
    "amount": 800
}'
```

### Event: Transfer
(POST) ```localhost:16091/event```

```
curl --location --request POST 'localhost:16091/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "transfer",
    "origin": "17",
    "amount": 50,
    "destination": "18"
}'
```

## Referências

* [Golang](https://laravel.com/docs/8.x/releases)
* [Docker](https://www.docker.com/)
* [Clean Code](https://github.com/jupeter/clean-code-php)
* [Clean Arch](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
* [Postman API CLient](https://www.postman.com/)
* [Golang with Clean Arch by Elton Minetto](https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f)
* [Golang with Clean Arch by Iman Tumorang](https://github.com/bxcodec/go-clean-arch)