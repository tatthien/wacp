# wacp

This cli tool watches a provided source file changes and copies to the provided destination.

`wacp` = Watch & copy. Is it simple?

## Installation

```bash
go get github.com/tatthien/wacp
```

For non-Go users

```bash
curl -sf https://gobinaries.com/tatthien/wacp | sh
```

## Usage

```bash
Usage: wacp -src <source path> -dest <destination path>
  -src  extension entrypoint source (optional). Default: ./dist/index.js
  -onChange command to be executed after file changed
```

## Examples

```bash
# Copy index.js to a new path after its content has been modified.
$ wacp -src index.js -onChange "cp index.js /new/path/index.js"
```

## License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022-present, Thien Nguyen
