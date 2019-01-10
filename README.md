# generate proto files
```
docker run --rm -v $(pwd):$(pwd) \
-w $(pwd) znly/protoc:0.4.0 \
-I ./proto \
--go_out=plugins=grpc:./proto/ \
proto/hello.proto
```

```
docker run --rm -v $(pwd):$(pwd) \
-w $(pwd) znly/protoc:0.4.0 \
-I ./proto \
--go_out=plugins=grpc:./proto/ \
--govalidators_out=./proto/ \
proto/hello.proto
```
