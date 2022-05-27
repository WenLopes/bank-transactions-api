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

No diretório **docs**, adicionei a collection para importação no API client [Postman](https://github.com/WenLopes/transaction-api/blob/master/docs/Postman%20-%20Transaction%20API.json).

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

## Considerações
### Escolha da linguagem
Embora a linguagem de programação cujo a qual possuo mais familiaridade seja o PHP, cheguei na conclusão que o Go (que no momento estou estudando) seria a melhor opção para realização do teste. O motivo para tal, vem da necessidade de utilizar um estado global para gerenciar os dados de contas e de transações.
O PHP utiliza a arquitetura [share-nothing](https://tideways.com/profiler/blog/php-shared-nothing-architecture-the-benefits-and-downsides), onde cada request é independente de outro. Apesar dessa abordagem trazer benefícios, para o teste em questão não se torna a melhor opção, tendo em vista que optei por não utilizar nenhum mecanismo para persistir os dados (banco de dados, cache, etc...). Por outro lado, o Go não se utiliza dessa arquitetura e no momento do desenvolvimento do teste, está alinhado com meus estudos atuais.

Obs: Utilizando [Swoole](https://www.swoole.com/) ou [Roadrunner](https://roadrunner.dev/) (esse que por sua vez, foi escrito em Go), seria possível implementar o comportamento necessário para realizar o teste com PHP. Entretanto, optei por utilizar o Go pelo motivo de estar mais alinhado com meus estudos atuais e também por nativamente (sem necessidade de instalar nada a mais) oferecer o suporte necessário para realizar o teste.

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