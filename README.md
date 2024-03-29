# revshellz

## Features

- Interact with multiple TCP connections concurrently.
- Terminal based UI.
- Simple stdout message handlers.
- TODO: Informational REST API.
- TODO: Web Socket based real time frontend.

![Terminal UI](ui.jpg)

## Architecture

- Interact with TCP connections via Go channel's for each connection.
- Thread Safe access to each connection's channels.
- REST API built on go-chi.
- Terminal UI using tview.

## Run

```bash
    cd src
    go run main.go
```

- Opens a listener on port 18000.
- Test it with `nc 127.0.0.1 18000`.

## TODO

- Better Logging.
- REST API: connection information.
- WebSockets: real time shell interaction.
