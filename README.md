# fresh [![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/fresh)](https://goreportcard.com/report/github.com/saromanov/fresh)

Check if exists new releases based on your go.mod file

## Usage

On the directory with `go.mod` execute

```
fresh check
```
And if the project contains outdated dependencies, it'll retrun something like

```
'semver'
current version: v1.5.0
new version v3.0.3
published at 2019-12-13 17:30:56 +0000 UTC
release description:
 ### Fixed

- #141: Fixed issue with <= comparison
```

## Docker

## License
MIT

