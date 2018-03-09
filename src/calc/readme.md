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
    ├── calc
    │   ├── calc.go
    │   └── readme.md
    └── simplemath
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

``` shell
$ ./calc # 输出帮助信息
$ ./calc add 3 2 # result: 5
$ ./calc sqrt 9 # result: 3
```