# mqstomp
send messages over stomp tooling

Example: go run send/send.go -message="this is a message 101." -header="CustomHeader1=value1" -header="CustomHeader2=value2" -header="JMSXGroupID=group1" -header="persistent=true"

Usage:
  -header value
    	Headers for message. -header="CusomHeader=Value"
  -host string
    	host header (default "/")
  -login string
    	login credentials (default "admin=admin")
  -message string
    	Message to to sent
  -queue string
    	Destination queue (default "/queue/test")
  -server string
    	STOMP server (default "localhost:61613")
