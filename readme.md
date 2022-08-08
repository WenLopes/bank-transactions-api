# Bank transactions Api

API execução de transações bancárias

## Tecnologias utilizadas no projeto

- [Docker](https://www.docker.com/)
- [Golang (Versão 1.18)](https://www.php.net/)
- [Clean arch](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Instalação

Após baixar o projeto, crie o arquivo .env com base no .env.example. 

Obs: Se desejar alterar a porta padrão da API, certifique-se de usar a mesma porta nas collections do postman (se caso desejar utiliza-la).

```bash
cp .env.example .env
```

Suba a aplicação, executando o comando

```bash
make up
```

## Rotas

No diretório **docs**, adicionei a collection para importação no API client [Postman](https://github.com/WenLopes/transaction-api/blob/master/docs/Postman%20-%20Transaction%20API.json).

### Reset
(POST) ```localhost:{API_PORT}/reset```

```
curl --location --request POST 'localhost:{API_PORT}/reset' \
--data-raw ''
```

### Balance
(GET) ```localhost:{API_PORT}/balance?account_id=18```

```
curl --location --request GET 'localhost:{API_PORT}/balance?account_id=18' \
--data-raw ''
```

### Event: Deposit
(POST) ```localhost:{API_PORT}/event```

```
curl --location --request POST 'localhost:{API_PORT}/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "deposit",
    "destination": "18",
    "amount": 1
}'
```

### Event: Withdraw
(POST) ```localhost:{API_PORT}/event```

```
curl --location --request POST 'localhost:{API_PORT}/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "withdraw",
    "origin": "18",
    "amount": 800
}'
```

### Event: Transfer
(POST) ```localhost:{API_PORT}/event```

```
curl --location --request POST 'localhost:{API_PORT}/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "transfer",
    "origin": "17",
    "amount": 50,
    "destination": "18"
}'
```

### Próximos passos
Considerando que o projeto evoluisse para notas etapas, observabilidade e documentação certamente seriam pontos primordiais para o bom funcionamento da aplicação. Sendo assim, seguem abaixo sugestões de ferramentas que em meu ponto de vista, atenderiam essa necessidade.

- Adicionar ferramenta para APM - Application Performance Monitoring ([New relic](https://newrelic.com/), [Dynatrace](https://www.dynatrace.com/platform/application-performance-monitoring/))

- Adicionar ferramenta para Tracing ([Jaeger](https://www.jaegertracing.io/), [New relic](https://newrelic.com/), [Dynatrace](https://www.dynatrace.com/platform/application-performance-monitoring/))

- Adicionar ferramenta para Dashboards ([Grafana](https://grafana.com/), [New relic](https://newrelic.com/))

- Adicionar ferramenta para Documentação ([Swagger](https://swagger.io/))

## Referências

* [Golang](https://laravel.com/docs/8.x/releases)
* [Docker](https://www.docker.com/)
* [Clean Code](https://github.com/jupeter/clean-code-php)
* [Clean Arch](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
* [Postman API CLient](https://www.postman.com/)
* [Golang with Clean Arch by Elton Minetto](https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f)
* [Golang with Clean Arch by Iman Tumorang](https://github.com/bxcodec/go-clean-arch)
* [Swoole](https://www.swoole.com/)
* [PHP - Roadrunner](https://roadrunner.dev/)
