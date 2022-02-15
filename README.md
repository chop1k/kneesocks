# Kneesocks

Kneesocks is dockerized socks proxy server with restriction functional.

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

|                Restrictions Functionality                | V4  | V4a | V5  |
|:--------------------------------------------------------:|:---:|:---:|:---:|
|               Whitelist for specific user                |  +  |  +  |  +  |
|               Blacklist for specific user                |  +  |  +  |  +  |
|                Prohibit specific command                 |  +  |  +  |  +  |
|              Prohibit specific address type              |  -  |  -  |  +  |
|             Rate limiting for specific user              |  *  |  *  |  +  |
| Maximum simultaneous connections limit for specific user |  *  |  *  |  +  |
|     Timeout in seconds for specific protocol command     |  +  |  +  |  +  |

## Getting started

For installation guide read the wiki.

### Use prepared configurations

If you don't want to waste your time for making your own, you can use prepared configuration. Just copy config that you
like and paste it. For more info see the wiki.

### Make your own config

If you want more understanding of how server will work, or if our prepared configurations don't fit your needs, you can
make your own config. For more information see the wiki.

## How to set up your environment and run tests

Read appropriate section of the wiki.

## Build with

- rs/zerolog - For logging to file and console.
- sarulabs/di - For dependency injection and management.
- go-playground/validator - For validating config.
- stretchr/testify - For testing assertions.
- emicpasic/gods - For data structures, like sets and maps.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull
requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details