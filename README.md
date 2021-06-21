# FTPLyzer 

FTPLyzer is tool to brute-force valid usernames on FTP server.

With certain setting administrator can validate allowed usernames against list and after them user will be prompt for credentials. This will give availability to enumerate through list of usernames. This tool will make concurrent connection to server for user validation. This tool was made as challenge project to understand golang and FTP security.

## Usage

```bash
$ ./FTPLyzer -h

        ___ _____ ___ _                   
        | __|_   _| _ \ |  _  _ ______ _ _ 
        | _|  | | |  _/ |_| || |_ / -_) '_|
        |_|   |_| |_| |____\_, /__\___|_|  
                            |__/            
   
   
        FTP User Enumeration tool.

flag needs an argument: -h
Usage of ./FTPLyzer:
  -c int
        Concurrency default (default 10)
  -continue
        Will Continue search for username upon first finding 
  -h string
        Enter hostname (default "127.0.0.1")
  -p string
        Enter port (default 21) (default "21")
  -w string
        Wordlist of usernames
```

-c flag will iterate through whole list. If this flag is not set. Program exit upon first finding.

## Example

```bash
$ ./FTPLyzer -c 25 -w usernames.txt -h 10.10.10.197

        ___ _____ ___ _                   
        | __|_   _| _ \ |  _  _ ______ _ _ 
        | _|  | | |  _/ |_| || |_ / -_) '_|
        |_|   |_| |_| |____\_, /__\___|_|  
                            |__/            
   
   
        FTP User Enumeration tool.

Attacking 10.10.10.197 on port 21 
Testing username: jenkinssortor

FOUND : developer

```

