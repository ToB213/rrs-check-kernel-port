```
 ├── Dockerfile
 ├── main.go
 └── maps（変換したいマップを格納したディレクトリ）
     ├── austria_austria18_01
     │   └── map
     │       ├── map.gml
     │       └── scenario.xml
     ├── canada_canada41_01
     │   └── map
```

```
 docker build ./ -t go:dev
 docker container run -it --rm --name check-klernel-port --mount type=bind,src=$PWD/maps,dst=/go/src/maps go:dev
```

```
 root@xxxxxx:go/src# go run main.go
 Do you want to update the port number? (1: Yes, 2: No)

 Processing maps/austria_austria18_01/config/common.cfg
 Current Port: 7000
 Replace Port: 27931
 Updated maps/austria_austria18_01/config/common.cfg
 ============================================================
 Processing maps/canada_canada41_01/config/common.cfg
 Current Port: 7000
 Replace Port: 27931
 Updated maps/canada_canada41_01/config/common.cfg
 ============================================================
 ...
```
