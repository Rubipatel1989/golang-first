# glonag-first
## Setup
1. Download from officially website.
2. if you are using ubuntu then follow listed steps.
 a. sudo wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
 b. sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
 c. nano ~/.bashrc
 d. copy paste listed code in above file
       export PATH=$PATH:/usr/local/go/bin
       export GOPATH=$HOME/go 
       export PATH=$PATH:$GOPATH/bin
 e. Save and exit, then run below commond
       source ~/.bashrc
## check version 
       go version
## run file
       go run hello.go

## Initialize Directory
If your project is not in /home/go/src, then you can do any where and need to execute below commond in project.
       go mod init goLangFirst


# tutorial link
https://www.youtube.com/watch?v=oE_ZksdTtaU&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=5