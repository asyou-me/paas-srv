#!/bin/bash

# 来源于网络，用于获取当前shell文件的路径
SOURCE="$0"
while [ -h "$SOURCE"  ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

protoc_path=/Users/maojian/Work/programfiles/protobuf/src/
gogoproto_path="$DIR/../../../"

protoc3 "$DIR/../conf/api.proto" --gofast_out="$DIR/../pkg/types" --js_out="$DIR/../gen/js"  --proto_path=${gogoproto_path} --proto_path="$DIR/../conf/"