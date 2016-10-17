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
