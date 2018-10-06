# eFuture

[![Build Status](https://travis-ci.org/xlui/eFuture.svg?branch=master)](https://travis-ci.org/xlui/eFuture)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/xlui/eFuture)

Send an Email to future!

It is a really amazing thing that some day you received a Email from past, isn't it?

## Design

The key point of this project is to store future-emails and send them at a specified time point. 
So we need users to input or select:

1. The **Email address** to receive the future-email
1. The **subject** of future-email
1. The **date** this future-email been deliveried
1. The **content** of this future-email

After user fill in the four part above properly, we now need a reliably component to reserve these Emails 
and to send them at a specified time point. For this project, I choose redis to guarantee future-email 
being consumed only at specified time and being backuped for unexpected server fault.

I choose to implement this project through `python`. At first I'd like to use `golang`, but using `golang` to
write web application is such an uncomfortable thing, so I finally choose to use `python`.

## RabbitMQ

We need rabbitmq to support delayed tasks. Using rabbitmq to provide delayed task we need 
a feature of queue: `x-dead-letter-exchange`. Declare a `x-dead-letter-exchange` property for a queue, 
the queue will automatically forward expired messages to the specified exchange. What we need to do is 
to receive messages for a queue which is bind to the exchange too.

## Redis

After trying to use RabbitMQ to implement `delayed tasks` I found that rabbitmq is not a good idea. 
The implement has a fatal flaw that in the dead queue, if the top task is not expired, 
although the following one or more tasks is expired, they won't be forwarded the specified exchange. 
This causes the shorter task waiting for the longer task and the shorter task cannot be consumed at time.

So I choose redis to implement the `delayed task`. `ZSet` is a PriorityQueue-Like data structure which 
can automatically sort tasks by a `score`. We can attach a timestamp to task when putting them into redis. 
By doing so, the top task in zset is the one bind to the smallest timestamp and also is the first task 
that should be done. And in our code, we check the first task in zset every second, compare its `score` with
`time.Now`'s timestamp. If the score is bigger than `time.Now`'s timestamp, this means that it is time to 
consume the top task.

## How to run?

In order to run this project, you need to edit your own config files first, here is a template:

```bash
{
  "username": "future@example.com",
  "password": "password",
  "smtp": "smtp.example.com",
  "smtp_port": "587",
  "redis_host": "127.0.0.1",
  "redis_port": "6379",
  "redis_password": "",
  "redis_db": 0,
  "log_path": "/tmp/eFuture.log",
  "log_level": "DEBUG"
}
```

`username`, `password`, `smtp` and `smtp_port` is required, because of eFuture need them to send Emails.
If you use docker to run this project, make sure to change the value of `redis_host` to `redis` if you use
link to link this container to a exist redis container.

In this project I have declared two ways to read configurations.

1. Environment variable `EFUTURE_CONFIG`
1. Config file `/data/eFuture/config.json`

The code will first try to read `${EFUTURE_CONFIG}` and if fails it will try to read `/data/eFuture/config.json`.
So make sure you have the above two ways at least one available.

For example, if you'd like to declare a environment variable and the file path is `/home/xxx/eFuture/config.json`.
You need to declare that:

```bash
export EFUTURE_CONFIG=/home/xxx/eFuture/config.json
```

Also, this project use Redis to provide delayed task queue, so redis is required.

After configuring config file and redis properly, run the project:

```bash
python main.py
```

## Docker

A docker version project is provided also.

```bash
docker pull xlui/efuture
docker run --name redis -p 6379:6379 -d redis
docker run --name efuture --link redis
-p 8080:5000
-v /path/to/config.json:/data/eFuture/config.json
-d xlui/efuture
```

## LICENSE

MIT
