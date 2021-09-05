# blt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/loozhengyuan/blt)](https://pkg.go.dev/github.com/loozhengyuan/blt)
[![Go Report Card](https://goreportcard.com/badge/github.com/loozhengyuan/blt)](https://goreportcard.com/report/github.com/loozhengyuan/blt)
[![test](https://github.com/loozhengyuan/blt/actions/workflows/test.yml/badge.svg)](https://github.com/loozhengyuan/blt/actions/workflows/test.yml)

Blocklist management tool.

## Usage

```console
$ blt -h
Blocklist management tool.

Usage:
  blt [command]

Available Commands:
  build       Builds blocklist according to a spec file
  help        Help about any command
  version     Prints the current version information

Flags:
  -h, --help   help for blt

Use "blt [command] --help" for more information about a command.
```

## Examples

### Export DNSBL - Simple Format

```yaml
export:
  destinations:
    - path: simple.txt
      customTemplate: |
        # DNS Blocklist - Simple Format
        # 
        # Generated using the blt tool.
        # https://github.com/loozhengyuan/blt

        {{ range .Items -}}
        {{ . }}
        {{ end -}}
```

### Export DNSBL - HOSTS Format

```yaml
export:
  destinations:
    - path: hosts.txt
      customTemplate: |
        # DNS Blocklist - HOSTS Format
        # 
        # Generated using the blt tool.
        # https://github.com/loozhengyuan/blt

        {{ range .Items -}}
        127.0.0.1   {{ . }}
        ::1         {{ . }}
        {{ end -}}
```

### Export DNSBL - DNSMASQ Format

```yaml
export:
  destinations:
    - path: dnsmasq.txt
      customTemplate: |
        # DNS Blocklist - DNSMASQ Format
        # 
        # Generated using the blt tool.
        # https://github.com/loozhengyuan/blt

        {{ range .Items -}}
        address=/{{ . }}/#
        {{ end -}}
```

### Export DNSBL - AdBlock Plus Format

```yaml
export:
  destinations:
    - path: adblockplus.txt
      customTemplate: |
        ! DNS Blocklist - AdBlock Plus Format
        ! 
        ! Generated using the blt tool.
        ! https://github.com/loozhengyuan/blt

        {{ range .Items -}}
        ||{{ . }}^
        {{ end -}}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
