#quick start   
#官网链接  https://gqlgen.com/getting-started/

##创建对应的文件夹 初始化go mod
mkdir self-gqlgen
$ cd self-gqlgen
$ go mod init github.com/[username]/gqlgen-todos

##第一次 执行初始化操作
go run github.com/99designs/gqlgen init


###gqlgen的配置文件，控制生成代码的旋钮。
gqlgen.yml — The gqlgen config file, knobs for controlling the generated code.
###gqlgen生成的代码文件
generated.go — The GraphQL execution runtime, the bulk of the generated code.
###生成构建图所需的模型。通常，您将使用自己的模型覆盖它们。对于输入类型仍然非常有用。自定义的一些模型
models_gen.go — Generated models required to build the graph. Often you will override these with your own models. Still very useful for input types.
###这是应用程序代码所在的位置。生成的。go将调用这个函数来获取用户请求的数据。需要自己实现对应的resolver中没实现的方法
resolver.go — This is where your application code lives. generated.go will call into this to get the data the user has requested.
server/server.go — This is a minimal entry point that sets up an http.Handler to the generated GraphQL server.



##自定义调整配置目录后 重新生成时候执行 会读取gqlgen.yml中配置的相关文件去做对应的生成  -v表示查看详情
go run github.com/99designs/gqlgen -v