# Golang concurrency example

此庫為 [Go 的並發：Goroutine 與 Channel 介紹](https://peterhpchen.github.io/2020/03/08/goroutine-and-channel.html) 一文的範例程式。

## How to use

* 使用客製的 Docker image：

```bash
docker run peter3598768/golang-concurrency-example single-thread.go
```

* 使用官方 Docker image：

```bash
docker run -v $(PWD)/src:/go golang go run single-thread.go
```

* 本機：

```bash
go run src/single-thread.go
```

> 可以將 `single-thread.go` 替換成任何想執行的檔案名。
