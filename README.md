# WatchMUD Client

A (super simple) client for the [WatchMud](http://github.com/trasa/watchmud) Project.

This connects to the watchmud server using gRPC and translates
line input from the player into gRPC commands. It is quite
light on features - basically enough done to verify that the server
implementation is working.

## Building

To build, test, and create the executable:

    $ make
  
## Running

Command line arguments, connection information, other important info
like that is still very much TODO. :( 

    $ ./watchmud-client -player "Ed the Destroyer"
    
## Commands

help
    get list of commands -Note- Not Implemented yet...

help foo
    get help for 'foo' -Note- Not Implemented yet...

quit or /q
    exit the client

