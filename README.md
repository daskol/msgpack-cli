# msgpack-cli

*command line msgpack encoder & decoder*

## Overview


### Building

```bash
    $ go install github.com/daskol/msgpack-cli
```

### Usage

```bash
    $ msgpack-cli encode <<< "Hello, World!" > hello.msgpack
    $ msgpack-cli decode < hello.msgpack
    "Hello, World!"
```
