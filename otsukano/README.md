## Scala install
### jdk install
```
$ sudo apt-get -y install default-jdk
$ sudo apt-get -y install scala
```

### sbt install
```
$ wget https://dl.bintray.com/sbt/debian/sbt-0.13.7.deb
$ sudo apt-get update
$ sudo dpkg -i sbt-0.13.7.deb
$ sbt
```

### compile by sbt
```
$ sbt
> run
```

## Learning Play Framework

### install activator
```
wget http://downloads.typesafe.com/typesafe-activator/1.2.12/typesafe-activator-1.2.12-minimal.zip
unzip typesafe-activator-1.2.12-minimal.zip
mv activator-1.2.12-minimal /usr/local/lib
ln -s /usr/local/lib/activator-1.2.12-minimal/activator /usr/bin/activator
```

### initialize play project
```
activator new play-hello-world play-scala
```

### run play application
```
activator run -Dhttp.port=3000
```

### start h2
```
cd slick-codegen/h2/
sh start.sh
```

### generate models
```
./sbt.sh gen-tables
```

### execute api commands
#### read
#### create
```
curl -H "Content-type: application/json" -XPOST -d '{"name":"TestUser", "companyId":1}' http://localhost:3000/json/create
```
#### update
```
curl -H "Content-type: application/json" -XPOST -d '{"id":1, "name":"TestUser", "companyId":1}' http://localhost:3000/json/update
```
#### raise error intentionally
```
curl -H "Content-type: application/json" -XPOST -d '{"userName":"TestUser"}' http://localhost:3000/json/create
```
#### delete
```
curl -XPOST http://localhost:3000/json/remove/1
```
