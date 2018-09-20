# eFuture

[![Build Status](https://travis-ci.org/xlui/eFuture.svg?branch=master)](https://travis-ci.org/xlui/eFuture)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/xlui/eFuture)

Send an email to future!

It is a really amazing thing that some day you received a email from past, isn't it?

## Design

The key point of this project is to store future-emails and send them at a specified time
point. So at user interface we need users to input or select:

1. The **Email address** to receive the future-email
1. The **subject** of future-email
1. The **date** this future-email been deliveried
1. The **content** of this future-email

After user fill in the four part above properly, we now need a reliably component to reserve these emails and to send them at a specified time point. For this project, I choose rabbitmq to guarantee future-email being consumed only at specified time and being backuped for unexpected server fault.

I choose to implement this project through `golang`.
