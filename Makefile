run_test:
	#参数covermode用来控制是否输出代码覆盖率，covermode的定义: count: 统计代码访问次数；set: 统计代码是否被访问; atomic: 一般在并发工程中使用
	#coverpkg用于指定统计哪些包下的覆盖率，若不指定该参数则只统计有测试代码的源文件。
	go test -v ./... -gcflags=all=-l -covermode=count -coverpkg=`go list ./... | tr '\n' ','` -coverprofile=./report/coverage.out
	go tool cover -html=./report/coverage.out -o ./report/coverage.html