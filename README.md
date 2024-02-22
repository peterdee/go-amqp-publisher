## go-amqp-publisher

This is a server application that also acts as a RabbitMQ publisher service

Application uses https://financialmodelingprep.com API to load stock data, and then this data is published to the queue

Consumer service: https://github.com/peterdee/go-amqp-consumer

### Environment variables

The `.env` file is required, see [.env.example](.env.example)

### Launch

Run RabbitMQ in Docker (if necessary):

```shell script
docker compose up -d
```

Launch the server:

```shell script
go run ./
```

With [AIR](https://github.com/cosmtrek/air):

```shell script
air ./
```

### License

[MIT](./LICENSE.md)
