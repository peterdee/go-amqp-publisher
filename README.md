## go-amqp-publisher

### Environment variables

The `.env` file is required, see [.env.example](.env.example)

### Launch

Run RabbitMQ in Docker first (if necessary):

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
