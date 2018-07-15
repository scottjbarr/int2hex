# int2hex

A command line utility to convert integers to hex.


## Install

    go get github.com/scottjbarr/int2hex


## Usage

Int values via stdin

```
$ echo 8 83 63 107 | int2hex
8 53 3f 6b
```

Or int values as argument

```
$ int2hex 8 83 63 107
8 53 3f 6b
```

## License

The MIT License (MIT)

Copyright (c) 2018 Scott Barr

See [LICENSE.md](LICENSE.md)
