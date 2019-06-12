#!/bin/bash
protoc -I ../api/proto/ ../api/proto/*.proto --go_out=plugins=grpc:../pkg/api/
