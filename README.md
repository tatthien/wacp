This cli tool watches a provided source file changes and copies to the provided destination.

## Installation

```
go get github.com/tatthien/directus-extension-copy
```

For non-Go users

```
curl -sf https://gobinaries.com/tatthien/directus-extension-copy | sh
```

## Usage

```
Usage: directus-extension-copy -src <source path> -dest <destination path>
  -src  extension entrypoint source (optional). Default: ./dist/index.js
  -dest extension entrypoint destination (required)
```
