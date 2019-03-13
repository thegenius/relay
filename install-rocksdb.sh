git clone https://github.com/facebook/zstd.git
cd zstd
make
sudo make install
cd -
sudo yum install -y snappy snappy-devel
sudo yum install -y zlib zlib-devel
sudo yum install -y bzip2 bzip2-devel
sudo yum install -y lz4-devel

git clone https://github.com/facebook/rocksdb.git rocksdbc++
cd rocksdbc++
make static_lib
cd -

export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH 
sudo ldconfig

CGO_CFLAGS="-I${PWD}/rocksdbc++/include" \
CGO_LDFLAGS="-L${PWD}/rocksdbc++ -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
  go get github.com/tecbot/gorocksdb
