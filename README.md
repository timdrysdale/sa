# sa
service agent: forwards local port to an external websocket server (e.g. relay), to facillitate admin on remote experiments behind restrictive firewalls 

Intended usage is via a relay that has topics, so that multiple experiments can access the same relay over :443.

Features
- reconnects to remote port if connection is broken (via [reconws](https://github.com/timdrysdale/reconws)
- makes a separate connection to the local port for each unique connection id for which messages are received
- local and remote port specified on the command line
- streams data (does not attempt to understand EOL etc)


