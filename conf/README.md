bucket.conf.example is an runtime config file example. 
When you run your own application, please add a new conf file named bucket.conf in the folder.

### Modify Hosts File
To run the example, you should modify your hosts first:
```sh
sudo vim /etc/hosts
```
Add the following line:
```sh
contianerops.me 0.0.0.0
```

If it doesn't work immediately, restart your system.

### Compile Wharf

```sh
cd /your/path/to/wharf
export GOPATH=/your/path/to/wharf
mkdir src
mkdir src/github.com
mkdir src/github.com/containerops.me
cd src/github.com/containerops.me
git clone -b 0.3.1-fix https://github.com/containerops/wharf.git 

cd /your/path/to/wharf
go get -u github.com/astaxie/beego
go get -u github.com/codegangsta/cli
go get -u github.com/siddontang/ledisdb/ledis
go get -u github.com/garyburd/redigo/redis
go get -u github.com/shurcooL/go/github_flavored_markdown
go get -u github.com/satori/go.uuid
go get -u github.com/nfnt/resize
go get -u github.com/tools/godep
go build
```

### Run Application
```sh
sudo ./wharf web --address 0.0.0.0
```

That's all.

