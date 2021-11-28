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

## Examples

Run a script

```bash
arr script path/to/file.R
```

Run R code

```bash
arr run 'packer::bundle()'
```

Using built-in convenience for certain packages.

```bash
arr devtools document check install
arr usethis mit_license pipe
arr packer bundle
arr rsConnect writeManifest
```

Passing arguments.

```bash
arr usethis gpl_license -1=3 -2=TRUE
arr packer bundle --arg1s="prod"
```