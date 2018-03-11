### go 实现的缩小版linux计算器

### 目录结构

``` shell
├── bin
│   ├── calc
│   └── readme.md
├── pkg
│   └── linux_amd64
│       ├── simple.a
│       └── simplemath.a
└── src
    ├── calc        # main func
    │   ├── calc.go
    │   └── readme.md
    └── simplemath  # add, sqrt 方法包
        ├── add.go
        ├── add_test.go
        ├── sqrt.go
        └── sqrt_test.go
```
### 如何运行

1. 编译代码

``` shell
$ cd bin/
$ go build calc
```

3. 编译成功后bin目录下生成**calc**文件, 运行命令

- help 

``` shell
$ ./calc # 输出帮助信息
############ output ###############
➜  bin git:(master) ✗ ./calc
usage: calc command [arguments] ...

The commands are:
        add     Addtion of two values.
        sqrt    Square root of a non-negative value.
```

- add

``` shell
# add 
$ ./calc add 3 2 # result: 5
```

- sqrt

``` shell
# sqrt
$ ./calc sqrt 9 # result: 3
```
4. 运行测试代码

``` shell
$ go test simplemath

############ output ###############
➜  go git:(master) ✗ go test simplemath
ok      simplemath      (cached)
```