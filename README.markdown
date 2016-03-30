# Daytime

Implementation of the Daytime protocol as outlined by [RFC 867](https://tools.ietf.org/html/rfc867). This is implemented as a package in Go and is relatively orthogonal to the net/http package.

# RFC 867

## Daytime Protocol

This RFC specifies a standard for the ARPA Internet community.  Hosts on the ARPA Internet that choose to implement a Daytime Protocol are expected to adopt and implement this standard.

A useful debugging and measurement tool is a daytime service. A daytime service simply sends a the current date and time as a character string without regard to the input.

### TCP Based Daytime Service

One daytime service is defined as a connection based application on TCP.  A server listens for TCP connections on TCP port 13.  Once a connection is established the current date and time is sent out the connection as a ascii character string (and any data received is thrown away).  The service closes the connection after sending the quote.

### UDP Based Daytime Service

Another daytime service service is defined as a datagram based application on UDP.  A server listens for UDP datagrams on UDP port 13.  When a datagram is received, an answering datagram is sent containing the current date and time as a ASCII character string (the data in the received datagram is ignored).

### Daytime Syntax

There is no specific syntax for the daytime.  It is recommended that it be limited to the ASCII printing characters, space, carriage return, and line feed.  The daytime should be just one line.

One popular syntax is: `Weekday, Month Day, Year Time-Zone`

Example: `Tuesday, February 22, 1982 17:37:43-PST`

# License

Copyright © 2015-2016 Christian R. Vozar of Rogue Ethic

Use of this source code is governed by a MIT license that can be found in the [LICENSE.markdown](LICENSE.markdown) file.
