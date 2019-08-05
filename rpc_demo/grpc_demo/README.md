# Grpc使用案例

## protoc生成代码

```sh
protoc --go_out=plugins=grpc:. hello.proto
```

## 生成TSL证书

```sh
openssl req -passout pass:password -new -x509 -keyout ca_p.pem -out ca.pem -subj "/CN=CHC/OU=RPC/O=GRPC/L=NY/ST=NY/C=CN"

openssl req -newkey rsa:2048 -nodes -out server.csr -keyout server.key -subj '/CN=chc.com/OU=RPC/O=GRPC/L=NY/ST=NY/C=CN'
openssl x509 -passin pass:password -days 3650 -sha256 -req -in server.csr -signkey server.key -CA ca.pem -CAkey ca_p.pem -CAcreateserial -out server.crt


openssl req -newkey rsa:2048 -nodes -out client.csr -keyout client.key -subj '/CN=chc.com/OU=RPC/O=GRPC/L=NY/ST=NY/C=CN'

openssl x509 -passin pass:password -sha256 -days 3650 -req -in client.csr -signkey client.key -CA ca.pem -CAkey ca_p.pem -CAcreateserial -nameopt RFC2253 -out client.crt

```

* 客户端证书:`client.crt,client.key`

* 服务端证书:`server.crt,server.key`,CA证书:`ca.pem`

  

