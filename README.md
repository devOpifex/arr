# arr

A light wrapper around R.

## Install

```bash
go get github.com/devOpifex/arr
```

or

```bash
go install github.com/devOpifex/arr@latest
```

## Help

```bash
arr -h
```

```bash
Call common R functions from your terminal

Usage:
  arr [flags]
  arr [command]

Aliases:
  arr, r

Available Commands:
  completion  generate the autocompletion script for the specified shell
  devtools    interact with the {devtools} package
  help        Help about any command
  packer      interact with the {packer} package
  rsconnect   interact with the {rsconnect} package
  run         run code
  script      run scripts
  usethis     interact with the {usethis} package

Flags:
  -h, --help          help for arr
  -p, --path string   path to R installation
  -v, --verbose       verbose output
      --version       version for arr

Use "arr [command] --help" for more information about a command.
```

## Examples

With built-in convenience for certain packages.

```bash
arr devtools document check install
arr usethis mit_license pipe
arr packer bundle
arr rsConnect writeManifest
```

Run a script

```bash
arr script path/to/file.R
```

Run R code

```bash
arr run 'packer::bundle()'
```