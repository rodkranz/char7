#/bin/bash/

go-bindata -o "./modules/chatdata/chatdata.go" -pkg "chatdata" .charset
