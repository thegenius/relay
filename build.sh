export LD_LIBRARY_PATH=/usr/local/mysql/lib:$LD_LIBRARY_PATH 
CGO_CFLAGS="-I${PWD}/rocksdbc++/include" \
CGO_LDFLAGS="-L${PWD}/rocksdbc++ -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
  go build
