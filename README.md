# wacp

This cli tool watches a provided source file changes and copies to the provided destination.

`wacp` = Watch & copy. Is it simple?

## Installation

```
go get github.com/tatthien/wacp
```

For non-Go users

```
curl -sf https://gobinaries.com/tatthien/wacp | sh
```

## Usage

```
Usage: wacp -src <source path> -dest <destination path>
  -src  extension entrypoint source (optional). Default: ./dist/index.js
  -dest extension entrypoint destination (required)
```
