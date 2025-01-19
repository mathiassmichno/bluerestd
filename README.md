# bluerestd

A cross-platform Bluetooth daemon with a OpenAPI-based REST API interface to control Bluetooth Classic functions.

Currently works on Linux and Windows, with FreeBSD and MacOS support coming soon.

## Funding

This project is funded through [NGI Zero Core](https://nlnet.nl/core), a fund established by [NLnet](https://nlnet.nl) with financial support from the European Commission's [Next Generation Internet](https://ngi.eu) program. Learn more at the [NLnet project page](https://nlnet.nl/project/bluetuith).

[<img src="https://nlnet.nl/logo/banner.png" alt="NLnet foundation logo" width="20%" />](https://nlnet.nl)
[<img src="https://nlnet.nl/image/logos/NGI0_tag.svg" alt="NGI Zero Logo" width="20%" />](https://nlnet.nl/core)

# Features
See the [feature matrix](https://github.com/bluetuith-org/bluetooth-classic?tab=readme-ov-file#feature-matrix) for a complete
list of currently supported features.

# Installation
- First, install the necessary [dependencies](https://github.com/bluetuith-org/bluetooth-classic?tab=readme-ov-file#dependencies) for your OS.
- Download the binary from the Releases page, and run it.

_Note for Windows_:
These builds are currently not signed, which means while launching this application,
Microsoft SmartScreen warnings may pop up. Press "Run anyway" to run the application.
Also, Windows Security (i.e. Antimalware Service Executable) may try to scan the application while it is being launched,
which will delay and increase the startup time.

# Usage
`bluerestd` provides a set of endpoints to control various Bluetooth classic functions.
Type `bluerestd -h` for a documentation on available commands.

## OpenAPI specifation
Type `bluerestd openapi -h` for a documentation on available OpenAPI commands.

To see the OpenAPI specification of these endpoints, type:
```
bluerestd openapi
```
and optionally pipe the output via [jq](https://jqlang.org/) to see the complete list of endpoints and
properties supported by the instance.

## Launch
Type `bluerestd launch -h` for a documentation on available launch commands.

There are two ways of launching the daemon:
- Using a TCP address, or
- Using a UNIX socket

### TCP
To launch it using a TCP address, type:
```
bluerestd launch -a <tcp-address>:<port>
```

Or, to use the application's default TCP address:port setting to launch the daemon, type:
```
bluerestd launch -t
```

For example, using a TCP address of "127.0.0.1:8000", the command would be:
```
bluerestd launch -a "127.0.0.1:8000"
```

### UNIX Socket 
To launch it using a UNIX socket, type:
```
bluerestd launch -s "<path/to/socket.sock>"
```

Or, to use the application's default UNIX socket setting to launch the daemon, type:
```
bluerestd launch -u
```

For example, using a UNIX socket of name "/tmp/bd.sock", the command would be:
```
bluerestd launch -a "/tmp/bd.sock"
```

## Accessing endpoints
If the TCP address is used and being listened on, an interactive API viewer
is present at the `/docs` endpoint. Access it in a web browser with the address
`http://127.0.0.1:8000/docs` if bluerestd is listening on "127.0.0.1:8000".

If the UNIX socket path is being listened on, the 'http+unix' protocol is used, and clients can connect using this protocol.
For example, to connect to the socket via 'curl', use:
`curl --unix-socket /tmp/bd.sock http://localhost/<endpoint>`

# Note
Bluerestd isn't really useful for Linux users since Bluez already exists, and usually should be preferred.
However, this project can act as documentation on how to interact with the Bluez daemon. 

# Credits
- [Huma](https://huma.rocks/) - For the OpenAPI conforming REST API framework
- [Scalar Docs](https://github.com/scalar/scalar) - For the interactive API viewer
- [PTerm](https://github.com/pterm/pterm) - For console handling and pretty printing
- [urfave/cli/](github.com/urfave/cli) - Command line parser framework