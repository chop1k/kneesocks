# Kneesocks

Kneesocks is dockerized socks server, with a lot of restriction features like whitelist, blacklist, rate limiting and
more.

Our official docker image located ad [docker hub](https://hub.docker.com/r/chop1k/kneesocks).

## Features

| Protocol Functionality  | V4  | V4a | V5  |
|:-----------------------:|:---:|:---:|:---:|
|         Connect         |  +  |  +  |  +  |
|          Bind           |  *  |  *  |  *  |
|     Udp Association     |  -  |  -  |  +  |
|    No authentication    |  +  |  +  |  +  |
| Password Authentication |  -  |  -  |  +  |
|    IPv4 Address type    |  +  |  -  |  +  |
|   Domain Address type   |  -  |  +  |  +  |
|    IPv6 Address type    |  -  |  -  |  +  |
|    Udp fragmentation    |  -  |  -  |  -  |

**NOTE**: Bind implementation have random behavior and is not ready for production, although you can enable it in
config.

|                Restrictions Functionality                | V4  | V4a | V5  |
|:--------------------------------------------------------:|:---:|:---:|:---:|
|               Whitelist for specific user                |  +  |  +  |  +  |
|               Blacklist for specific user                |  +  |  +  |  +  |
|                Prohibit specific command                 |  +  |  +  |  +  |
|              Prohibit specific address type              |  -  |  -  |  +  |
|             Rate limiting for specific user              |  *  |  *  |  +  |
| Maximum simultaneous connections limit for specific user |  *  |  *  |  +  |
|     Timeout in seconds for specific protocol command     |  +  |  +  |  +  |

**NOTE**: Socks v4 and v4a protocol not support authentication and users, but you can restrict it anyway.

## Specifications implemented

- SOCKS V4 [Protocol](http://ftp.icm.edu.pl/packages/socks/socks4/SOCKS4.protocol)
- SOCKS V4a [Protocol](http://www.openssh.com/txt/socks4a.protocol)
- SOCKS V5 [Protocol](https://datatracker.ietf.org/doc/html/rfc1928)
- Password authentication for socks v5 [protocol](https://datatracker.ietf.org/doc/html/rfc1929)

## Getting started

For installation guide read the [wiki](https://github.com/chop1k/kneesocks/wiki).

## Built with

- [rs/zerolog](https://github.com/rs/zerolog) - For logging to file and console.
- [sarulabs/di](https://github.com/sarulabs/di) - For dependency injection and management.
- [go-playground/validator](https://github.com/go-playground/validator) - For validating config.
- [stretchr/testify](https://github.com/stretchr/testify) - For testing assertions.
- [emicpasic/gods](https://github.com/emirpasic/gods) - For data structures, like sets and maps.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details about contributing.
See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for our code of conduct.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details