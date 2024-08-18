# Task Management Tool (tmt)

`tmt` is a simple command line tool for managing tasks.  
In particular, manage only repetitive tasks

```bash
$ tmt add -t "Task name" -r "* * 1-5"
$ tmt list
ID      TITLE           RECURRENCE      DESCRIPTION     NEXT
#1      Task name       * * 1-5                         2024-08-19
```

## Installation

Install from GitHub Releases.

https://github.com/imishinist/tmt/releases

### example

Install to `/tmp/tmt` directory.

```bash
mkdir -p /tmp/tmt
version=0.1.0 curl -sSL -o- https://github.com/imishinist/tmt/releases/download/v${version}/tmt_Linux_x86_64.tar.gz | tar xzvf - -C /tmp/tmt
```

or install from go install.

```bash
go install github.com/imishinist/tmt@latest
```

## License

Apache License 2.0

## Author

- Taisuke Miyazaki [@imishinist](https://github.com/imishinist)

