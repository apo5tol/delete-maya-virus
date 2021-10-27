# Delete maya virus
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/apo5tol/delete-maya-virus)
![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/apo5tol/delete-maya-virus/Go/main)

Removes Vaccine and MayaMelUIConfigurationFile nodes from .ma files.

## Download
[Windows x64 version](https://github.com/apo5tol/delete-maya-virus/releases/download/v1.2/delete-maya-virus_windows_x64.exe)  
[Linux x64 version](https://github.com/apo5tol/delete-maya-virus/releases/download/v1.2/delete-maya-virus_linux_x64)

## How to build

[Install Go](https://golang.org/doc/install)  
In that project folder run:

```bash
go build .
```

## How to use

You can run the script with no parameters and it will get all the .ma files in the current directory and the script will process them, or you can pass the path to the file or directory to the script to process.

```bash
./delete-maya-virus .../...file.ma   # unix

delete-maya-virus .../...file.ma     # windows
```

### Flags

**-r** - recursive path traversal  
**-b** - make a backup copy of processed files
