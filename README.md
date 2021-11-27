# arr

A light wrapper around R.

## Install

```bash
go get github.com/devOpifex/arr
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