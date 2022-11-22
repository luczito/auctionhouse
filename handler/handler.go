package main

//handler file. Handler will be the middleman between clients and servers, and parse the requests from clients to all servers,
//and then wait for a response from a majority of the servers before sending ACK to the client.

//this is to ensure stability and that all the servers are synces.
//this file will create X servers.
