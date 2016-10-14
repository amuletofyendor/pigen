# pigen

Generates hexidecimal Pi (with a given starting index and span) as a series of bytes, using nstruthers/pihex.

See nstruthers/pihex for the implementation of the BBP formula.

## Sample usage

    pigen -i 0 -s 10000000 | hexdump

Now sit back and wait nine hundred years or so as the secrets of the universe* spill by.

Or dump to a file while using `pv` to monitor your progress (remember the byte size of the output will always be half the chosen span of digits):

    pigen -i 0 -s 250000 | pv -pterb -s 125000 > pi-0-249999.hex

*Pi may or may not contain all the secrets of the universe.
