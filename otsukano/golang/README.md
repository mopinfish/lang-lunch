## otsukano golang project

### install gvm

get installer

```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

add below to .bashrc

```
[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"
source $HOME/.bashrc
```

install golang by gvm

```
gvm listall
gvm install go1.4 -B
gvm use go1.4
gvm install go1.7.5
```

### install golang by apt

```
sudo add-apt-repository ppa:evarlast/golang1.4
sudo apt-get update
sudo apt-get install golang
```

confirm version

```
go version
```
