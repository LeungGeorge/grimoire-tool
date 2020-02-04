# grimoire-tool

为简洁而生，基于 [urfave/cli](https://github.com/urfave/cli) 实现。

## Description

一键工具：

- `github` 代码库同步（sync）一键同步 github 仓库。
- `hexo` 博客预览（preview）、发布（publish）
- `space` 目录空间查看工具

## Usage

```html
NAME:
   grimoire-tool - grimoire-tool is a tool of grimoire, batch execute commands.

USAGE:
   grimoire-tool [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   space    space of current directory.
   update   upgrade to newest version of master.
   help, h  Shows a list of commands or help for one command

   git:
     git  git commands: sync.

   hexo:
     hexo  hexo blog commands: preview, publish.

GLOBAL OPTIONS:
   --message value, -m value  the commit message of git (default: "auto commit")
   --help, -h                 show help
   --version, -v              print the version
```

### git 仓库同步

执行：
```shell
grimoire-tool --m "auto commit " git sync
```

```html
2019/10/05 11:34:56 status:
2019/10/05 11:34:56 ## master...origin/master
 M command/git/status.go
 M command/git/sync.go

2019/10/05 11:34:56 has file
2019/10/05 11:34:56 Execute git status finished, some file will be committed
2019/10/05 11:34:56 Execute git add finished.
2019/10/05 11:34:56 Execute git commit finished.
2019/10/05 11:34:58 Execute git pull finished.
2019/10/05 11:35:03 Execute git push finished.
2019/10/05 11:35:03 Execute Command finished.
```

### grimoire-tool space

查看目录空间：

```html
➜  grimoire-tool git:(master) grimoire-tool space                      
16B     .gitignore
116B    after
117B    before
151B    CHANGELOG.md
396B    flag
814B    Gopkg.toml
1K      Gopkg.lock
2K      main.go
2K      README.md
7K      command
10K     lib
34K     LICENSE
1M      .git
```

## Install

To install, use `go get`:

```shell
go get -u github.com/LeungGeorge/grimoire-tool
```

## Contribution

1. Fork ([https://github.com/LeungGeorge/grimoire-tool/fork](https://github.com/LeungGeorge/grimoire-tool/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[LeungGeorge](https://github.com/LeungGeorge)
