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

## Completion

See documentation.

```bash
arr completion -h
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
# arr package function
arr devtools document check install
arr usethis use_mit_license use_pipe
arr packer bundle
arr rsConnect writeManifest
```

Passing arguments.

```bash
# bare arguments
arr usethis use_gpl_license -1=3 -2=TRUE

# as quoted string
arr packer bundle --arg1s="prod"
```
