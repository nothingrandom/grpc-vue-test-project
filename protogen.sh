#!/usr/bin/env bash

set -e

INPUT_DIR='proto/examplecom/library'
OUTPUT_DIR='_proto'

mkdir -p ${OUTPUT_DIR}
rm -r ${OUTPUT_DIR} 2> /dev/null
mkdir ${OUTPUT_DIR}

ls ${INPUT_DIR}/*.proto

protoc \
--proto_path=${INPUT_DIR} \
--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
--js_out=import_style=commonjs,binary:${OUTPUT_DIR} \
--ts_out=service=true:${OUTPUT_DIR} \
${INPUT_DIR}/*.proto

ls ${OUTPUT_DIR}

ls ${INPUT_DIR}/*.proto

protoc \
--proto_path=${INPUT_DIR} \
--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
--js_out=import_style=commonjs,binary:${OUTPUT_DIR} \
--ts_out=service=true:${OUTPUT_DIR} \
${INPUT_DIR}/*.proto

ls ${OUTPUT_DIR}
