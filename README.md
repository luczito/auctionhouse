# AuctionHouse
## How to run 
Each manager and each client will need its own terminal window.
### run the managers
cd into auctionhouse/manager and run the command "go run ."
This will run each manager with the default number of managers (3) and default ports (5000, 5001 and 5002).

If you wish to set the port manually this can be done by typing "-port <int port>" after the "go run ."

If you wish to increase or decrease the number of managers manually this can be done by typing "-managers <int managers>".

### Run the clients
cd into auctionhouse/client and run the command "go run ." 
This will run each client with the default port (5000) and default primary manager (127.0.0.1:5002).

Use the commands as explained by the terminal.

If you wish to set the port manually this can be done by typing "-port <int port>" after the "go run ."

If you wish to set the primary manager manually this can be done by typing "-primary <string ip>" after the "go run ."

## Known issues and additional info
After crashing the primary manager, if you bid before the election have run, the client will crash.

There is currently no way to add an auction to the system.

There is currently no timer on the auction running. Therefore you will always be able to bid and the result will always return that the auction is over.

Since the bid is using int32 the max bid i equal to the max int32.
