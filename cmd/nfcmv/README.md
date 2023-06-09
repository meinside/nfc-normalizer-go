# nfcmv

Renames file names to NFC recursively.

Inspired by [convmv](https://formulae.brew.sh/formula/convmv).

## Install

```go
$ go install github.com/meinside/nfc-normalizer-go/cmd/nfcmv@latest
```

## Usage

```
$ nfcmv -h
Convert files/directories' names to NFC-normalized ones.

Usage:

        $ nfcmv [parameters ...] [FILES_OR_DIRS ...]

* parameters:
        -h, --help: show this help message.
        -v, --version: show the version string of this application.
        -r, --recursive: convert all files and directories recursively.
        -d, --dryrun: dry run, do not actually convert anything.
        -i, --interactive: convert interactively, asking [y/N] on each file/directory.
        -f, --forcereplace: force replace if there are files/directories with the same name.
        -V, --verbose: print verbose messages.
```

For renaming a file or directory:

```bash
$ nfcmv /some/file-or-directory
```

For renaming all the things **recursively** in a directory:

```bash
$ nfcmv -r /some/directory
```

For forcing it to **overwrite** all the existing files/directories:

```bash
$ nfcmv -r -f /some/directory
```

For **confirming** everything interactively:

```bash
$ nfcmv -r -i /some/directory
```

For **simulation only**, touching nothing:

```bash
$ nfcmv -r -d /some/directory
```

## License

MIT

