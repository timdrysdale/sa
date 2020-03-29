# SA
![alt text][logo]

![alt text][status]
![alt text][coverage]

Service Agent connects a local port with an external websocket relay server [timdrysdale/shellbar](https://github.com/timdrysdale/shellbar)

## Why

Remote laboratory experiments are often located behind firewalls without the prospect of installing a conventional jumpserver to aid administration. While there are a number of *nix tools that could relay connections on a one-off basis with just a bit of script-glue (various combinations of socat and websocat), there is no existing solution that

+ is well-suited to management of large numbers remote experiments
+ can handle fully interactive shells without some gremlins, such as fatal connection loss on Ctrl+C
+ supports multiple simultaneous independent client connections to the same service
+ allows all clients and servers to simultaneously connect to the relay on the same port yet keep traffic separate

A connection-aware local relay is required on the experiment, so that it can work with the external websocket relay to keep the streams separate, and to faciliate new connections to the service whenever a client joins, as is required to support server-speaks-first protocols like ```ssh```. On the client side, no special connection aware-ness is required.

## Features

+ reconnects to external relay if connection is lost (via [timdrysdale/reconws](https://github.com/timdrysdale/reconws))
+ makes a separate connection to the local port for each unique connection id for which messages are received
+ configure local and remote addresses with environment variables (container-friendly)
+ streams data (does not attempt to understand EOL etc)
+ acts as intermediary for both servers (```sa serve```) and clients (```sa connect```)

## Getting started

-- TODO --

## Usage

-- TODO --


[status]: https://img.shields.io/badge/alpha-in%20development-red "Alpha status; in development"
[coverage]: https://img.shields.io/badge/coverage-00%25-orange "Test coverage NA"
[logo]: ./img/logo.png "sa logo - shell with multiple connections"