# Go Command

go常用命令

格式：

- `command`

    > Comments

    - `<subcommand>`

        > Comments

    - `flags`

        > Comments

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

<!-- vim-markdown-toc -->

---

---

- `bug`：在'golang/go'仓库启动一个新的issue

    Usage: go bug

- `build`：编译软件包及其依赖项

    Usage: go build [-o output] [-i] [build flags] [packages]

    > 编译由导入路径(import paths)命名的包及其依赖项，并不会安装编译结果
    >
    > 如果要build的是单个目录中的'.go'文件列表，则build会将其视为单个程序包的源文件列表
    >
    > 编译软件包时，build会忽略以'_test.go'结尾的文件
    >
    > build执行先入为主：编译单个'main'程序包时，build将生成的可执行文件写入到以第一个源文件（'go build ed.go rx.go'写入'ed'或'ed.exe'）命名的输出文件或源码目录（'go build unix/sam'写入'sam'或'sam.exe'）
    >
    > 当编译多个软件包或者是单个的非'main'软件包时，build会正常编译，但会丢弃生成结果，仅用做检查编译对象是否可以构建软件包

    - `-o`：指定build执行结果的输出位置

        > 强制build指令将生成的可执行文件/对象写入到指定的输出文件/目录，而不是输出到默认文件/目录
        >
        > 如果指定的输出是一个已存在的目录，则生成的所有可执行文件都将写入该目录

    - `-i`：安装作为目标依赖项的软件包

    - '[build flags]'：**构建标志，由`build`、`clean`、`get`、`install`、`list`、`run`、`test`指令共享**

        - `-a`：强制重建软件包

        - `-n`：

- `env`：打印go的变量信息

    Usage: go env [-json] [-u] [-w] [var ...]

    > 默认以shell脚本的格式

    - `<ENV_NAME>`：变量名

        > 单独打印指定的变量信息

    - `-json`：以JSON格式打印
