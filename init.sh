#############################################################
#   File Name: init.sh
#     Autohor: Hui Chen (c) 2018
# Create Time: 2018/04/24-14:41:49
#############################################################
#!/bin/sh 
cd /home/mengchun/go/src/httpproxy4blockchain
nohup go run server/server.go > logpath/out.log &
