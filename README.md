# Bruteforce #

Bruteforce it's a simple script/application for search servers with easy passwords.
Uses ssh protocol.

```
Example: 
login: test
password: test
```

## Install golang ##
```
wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
tar -xvf go1.6.linux-amd64.tar.gz 
mv go /usr/local
```
## Set $PATH-es ##
You have to add pathes into .bashrc
```
nano /home/%username%/.bashrc
```
## Install optional package ##
```
mkdir -p $GOPATH/src/golang.org/x/
Then clone crypto from github.

cd $GOPATH/src/golang.org/x/
git clone https://github.com/golang/crypto.git
```
```
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=/usr/local/go/
export PATH=$GOPATH/:$PATH
```

For run bruteforce make next: 

`go run bruteforce.go`

Application requires three dictionary and package ssh.
We have three dictionary ip(contains ip address of servers), logins(contains names users), passwords(contains password users) 


