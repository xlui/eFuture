# eFuture

[![Build Status](https://travis-ci.org/xlui/eFuture.svg?branch=master)](https://travis-ci.org/xlui/eFuture)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/xlui/eFuture)

Send an email to future!

It is a really amazing thing that some day you received a email from past, isn't it?

## Design

The key point of this project is to store future-emails and send them at a specified time point. 
So we need users to input or select:

1. The **Email address** to receive the future-email
1. The **subject** of future-email
1. The **date** this future-email been deliveried
1. The **content** of this future-email

After user fill in the four part above properly, we now need a reliably component to reserve these emails 
and to send them at a specified time point. For this project, I choose redis to guarantee future-email 
being consumed only at specified time and being backuped for unexpected server fault.

I choose to implement this project through `python`. At first I'd like to use `golang`, but using golang to
write web application is such a uncomfortable thing, so I finally choose to use `python`.

## RabbitMQ

We need rabbitmq to support delayed tasks. Using rabbitmq to provide delayed task we need 
a feature of queue: `x-dead-letter-exchange`. Declare a `x-dead-letter-exchange` property for a queue, 
the queue will automatically forward expired message to the specified exchange. What we need to do is 
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
