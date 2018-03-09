# netsend
cli network sharing
(no error checking in the golang app, sorry)
## netshare
Core application
### Usage

Wait for a file and output to "myfile.txt"
```
./netsend -l 5000 > myfile.txt
```

Send to a listening computer
```
cat myfile.txt | ./netsend -s mycomputer:5000
```

Serve file over http (for curl)
```
cat cat myfile.txt | ./netsend -w 5000
```

## ns.sh
Wrapper for constant sending between machines (change REMOTE_HOST)

### Usage
Wait for a file and output to "myfile.txt"
```
./ns.sh get > myfile.txt
```

Send to a listening computer
```
cat myfile.txt | ./ns.sh put
```

Serve file over http
```
cat cat myfile.txt | ./ns.sh web
```
Get file from web
```
./ns.sh wget > myfile.txt
```
