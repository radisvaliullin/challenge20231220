# Design

## Requirements
* Simple cli application "Scanner" to scan Port on HOST and detect MySQL server.
* Scanner should get as much info form connection handshake as can without authentication.
* It should be simple tool just doing one thing but as good as can.\
For given a port and a host return maximum info about MySQL server.
* provide built, test details.

## Approach
* Golang cli applicatio.
* Input data. Default host is localhost and default port is default MySQL port (hardcoded in code).\n
Accept host and port variable via cli arguments.
* Output. Print MySQL server information to Stdout.
* Implement package for handle MySQL connection (handshake).
* Implement scan package for scanning host/port using MySQL conn package.
* Using conn and scan package implement main package.
* Add tests.
* Add Makefile (build).
