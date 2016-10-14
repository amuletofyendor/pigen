# pigen

Generates hexidecimal Pi (with a given starting index and span) as a series of bytes, using nstruthers/pihex.

See nstruthers/pihex for the implementation of the BBP formula.

## Sample usage

    pigen -i 0 -s 10000000 | hexdump 
    
Now sit back and wait nine hundred years or so as the secrets of the universe* spill by.

*Pi may or may not contain all the secrets of the universe
