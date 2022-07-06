# lots of useful stuff can be done here
# then we replace bash process with exec
# and handle control to what we exec 
# the command will run with PID=1
exec $@
