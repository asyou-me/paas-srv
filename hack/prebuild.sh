#!/bin/bash

# 来源于网络，用于获取当前shell文件的路径
SOURCE="$0"
while [ -h "$SOURCE"  ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

protoc_path="$DIR/../../../../"

protoc3 "$DIR/../conf/proto/app.proto" \
--gofast_out="$DIR/../pkg/types" \
--js_out=import_style=commonjs,binary:"$DIR/../_out/js"  \
--proto_path=${protoc_path} \
--proto_path="$DIR/../conf/proto"

