# goworker-mails
Sending mails using Goworker

## Usage

Start the redis server:

```
$ redis-server
```

Build the go file:

```
$ go build
```

Run your worker:

```
$ ./worker -queues=mail
```

And inside the ruby app make:
```
$ bundle install
$ ruby mailer.rb
```

Or also you can test with redis command line:

```
$ redis-cli -r 1 RPUSH resque:queue:mail '{"class":"Mailer","args":["sender@gmail.com","receiver@outlook.com"]}'
```