For POS devices we may need to have a TCP-based NTP server , 
So just build it and use it as a simple NTP TCP Go project (socket) ;)

to test befor build:
```$go run TCPNTP.go <TCP PORT>```
  
 Hint :
  After build finished, You may make it as a systemd service which could be run on startup and act an NTP Server!
  
  Have fun!!!
