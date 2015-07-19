Description
===========
This package contains example golang code that shows how to serve HTTPS in three different modes:

- verified-none: Neither the client nor the server validate the hostname or the certificate chain. Any certificate presented by the server is accepted by the client. MITM is trivially possible in this mode.
- verified-server: Client validates the certificate presented by the server.
- verified-mutual: Client and Server mutually validate each other's certificates.

A script sets up a dummy CA and some self signed certificates to run the example code.
