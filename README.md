# Delete maya virus


Removes Vaccine and MayaMelUIConfigurationFile nodes from .ma files.

## How to use

You can run the script with no parameters and it will get all the .ma files in the current directory and the script will process them, or you can pass the path to the file or directory to the script to process.

### Flags
**-r** - recursive path traversal  
**-b** - make a backup copy of processed files

## How to build
[Install Go](https://golang.org/doc/install)  
In that project folder run:
```bash
go build

# then:
./delete-maya-virus .../...file.ma   # unix

# or:
delete-maya-virus .../...file.ma     # windows
```  

