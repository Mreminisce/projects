# beego-monteblog

### Introduction:

A blog demo adapted from Monte. -- [Gitee](https://gitee.com/perma/goBlog) 

```bash
$ mysql -u root -p
$ create database monteblog;
# Import monteblog.sql into the database
# Modify conf/app.conf if necessary
$ go mod init monteblog
$ go get
$ bee run
# Default admin account:username:admin, password:123456
```

### Requirements:

- Beego
- Bee
- Mysql

### Todo:

- Fix bugs: register & login, search, paginator...