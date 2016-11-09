Char7  [![Build Status](https://travis-ci.org/rodkranz/char7.svg?branch=master)](https://travis-ci.org/rodkranz/char7) [![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](https://gobuild.io/rodkranz/char7/master)
==========================================
Replace Utf8 encoding to Ascii code

-----------------------
### Source ###

* Char7 Source
* Version: 1.0.0
* License: ISC


-----------------------

## How to Compile it

 You can see the full available binary list [here](https://gobuild.io/rodkranz/char7) 
 or compile those files from different platforms in your owner computer.

### Download 
    
   Download for [Mac OSx](http://tmpcode.com/char7/char7-darwin-amd64.zip)
   
   Download for [Linux 386](http://tmpcode.com/char7/char7-linux-386.zip)
   
   Download for [Linux amd64](http://tmpcode.com/char7/char7-linux-amd64.zip)
   
   Download for [Linux arm](http://tmpcode.com/char7/char7-linux-arm.zip)
   
   Download for [Windows 386](http://tmpcode.com/char7/char7-windows-386.zip)
   
   Download for [Windows amd64](http://tmpcode.com/char7/char7-windows-amd64.zip)
    
### Requirements

* [GO Language](https://golang.org/doc/install)

#### Compiling to *Linux*

	$ env GOOS=linux GOARCH=arm GOARM=7 go build -o c7 main.go


#### Compiling to *MAcOSX*

	$ env GOOS=darwin GOARCH=386 go build -o c7 main.go


#### Compiling to *Windows*

	$ env GOOS=windows GOARCH=386 go build -o c7.exe main.go


-----------------------

## Install

You can copy the application for: 
 
 Linux: `sudo mv c7 /usr/bin`
 
 MacOSx: `sudo mv c7 /usr/bin`

 Window: `copy c7.exe /window/c7.exe`

-----------------------

## Parameters

##### Main Helper: 
    
Execution: `c7 -h`

    NAME:
       CharSet - CharSet convert
    
    USAGE:
       c7 [global options] command [command options] [arguments...]
       
    VERSION:
       v1.0.0
       
    COMMANDS:
         charset   Change charset to html cod
         recovery  recovery backup
         help, h   Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version

##### Charset Helper:

Execution: `c7 charset -h`

    NAME:
       c7 charset - Change charset to html cod
    
    USAGE:
       c7 charset [command options] [arguments...]
    
    DESCRIPTION:
       change key in map to value
    
    OPTIONS:
       --ext value, -e value           Define the extension type to filter files (default: ".php,.html,.htm,.js,.asp,.tpl,.txt")
       --file value, -f value          Define the file name that needs to change.
       --dir value, -d value           Define the folder that will looking for files. (default: "./")
       --backup, -b                    Disable backup file when change charset
       --backupName value, --bn value  Define the file name that needs to change. (default: ".c7")
       
##### Charset Helper:

Execution: `c7 recovery -h`
    
    NAME:
       c7 recovery - recovery backup
    
    USAGE:
       c7 recovery [command options] [arguments...]
    
    DESCRIPTION:
       restore backup files and overwride the file wripped.
    
    OPTIONS:
       --dir value, -d value           Define the folder that will looking for backup's file. (default: "./")
       --backupName value, --bn value  Search files name with extension backup c7. (default: ".c7")
       --remove, -r                    Delete the backup file when finish.
       

-----------------------

## Configuration / Customization

You can customizer the list of maps that will be replaced, you can find the list in your 
home folder e.g.: `~/.c7map`, everything has in `key` at map will replace for `value`.
 
Config example: 

    [
        {
            "key": "»",
            "value": "&#187;"
        },
        {
            "key": "¼",
            "value": "&#188;"
        },
        {
            "key": "½",
            "value": "&#189;"
        }
    ]
    
Tips: You can use anything to change, example `if(` for `if(!`.

