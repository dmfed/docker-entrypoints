## System signals when using shell script as entry point

To try all the stuff yourself:
```
docker build -t test-app .			// to build
docker run -it --rm -p 8080:8080 test-app	// to run
docker exec -it <container> /bin/bash		// to get into the shell
```

1. **CMD ./entry.sh** in Dockerfile + **exec ./app** in entry.sh
This is a perfectly valid case. CMD runs shell under the hood 
which executes entry.sh script. It in turn replaces itself with 
app beause we're using **exec**, so we're getting 
app with PID=1.

The app is able to trap system signals and act accordingly. 
Running **docker stop** will have Docker send SIGTERM to PID 1.
If nothing happens Docker will force kill the app when grace period
expires. 

2. **CMD ./entry.sh** in Dockerfile + **./app** in entry.sh
This will result in CMD launching shell which will run app
as a separate process with PID != 1 (>1). 

The app will not be able to trap system signals. On calling 
**docker stop** SIGTERM will be send to container but there will be
no one around to react. The app will be force-killed as a 
result.
 
3. **CMD ["./entry.sh"]**  will fail, since **entry.sh** is not
a binary executable. 

3. **ENTRYPOINT ["/bin/sh", "./entry.sh"]** in Dockerfile 
results in **./entry.sh** being executed on launch regardless
of what command we pass to the container. This is useful when 
certain setup work needs to be done before we actually run anything.
Again because of using **exec** in shell script the actual command 
gets PID=1.


