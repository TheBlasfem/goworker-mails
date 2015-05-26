# goworker-mails
Sending mails using Goworker

## Usage

Start the redis server:

```
$ redis-server
```

Build the go file:

```
$ go-build
```

Run your worker:

```
$ ./worker -queues=mail
```

You can test with redis command line:

```
$ redis-cli -r 10 RPUSH resque:queue:mail '{"class":"Mailer","args":[]}'
```