## GO BUILD HELPER

Let me introduce the Golang build helper.
It is true that trivial Goland demo examples have  valid files for describing extern packages. In such cases 'go build' command we see lines with advisies to fetch these packages by 'go get package_name' command.
This helper avoids such uncomfortable casies. It just parses output of 'go build' command, and fetches needed packages in automatic mode. Next it runs 'go build' again.
So, building Golang examples from many handy works becomes easier. 

# Compile it:
go build

# Installation and using it

1)Just copy executable binary gg and g.cmd or g.sh to some directory wich is declared in Path variable of operating system.

For Linux it may be /usr/local/bin
For Windows it may be C:/Windows or C:/Windows/System32 or something else as you prefere.

2)Run g.cmd or g.sh and compile golang code with automatic fetching all needed packagies 

# Content of Linux g.sh script for call it:

( go build | gg ) || go build

# Content of Windows g.cmd script for call it:

go build | gg.exe /n
go build


