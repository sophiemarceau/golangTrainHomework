按照自己的构想，写一个项目满足基本的目录结构和工程，
代码需要包含对数据层、业务层、API 注册，
以及 main 函数对于服务的注册和启动，信号处理，
使用 Wire 构建依赖。可以使用自己熟悉的框架。

1 一个项目架构tree目录及结构
|- api                  # api目录是对外保留的proto文件及生成的pb.go文件
|- cmd		        # 项目主干，main所在
|   |-- myapp
|      |--- main.go
|- configs		 # configs 为配置文件目录
| --db.toml					
|- internal              # 项目内部包
|   |--dao               # dao层，用户数据库、cache、MQ等资源访问
|   |--di	         # 依赖注入层 采用wire静态分析依赖
|      |--- wire.go      # wire 声明
|      |--- wire_gen.go  # go generate 生成的代码
|   |--model		 # model 层，用于声明业务结构体
|   |--server            # server层，用于初始化grpc和http serverå
|   |--service           # service层，用于业务逻辑处理
|- pkg                   # 放置可以被外部程序安全导入的包  pkg 目录下的包一般会按照功能进行区分，例如 /pkg/cache 、 /pkg/conf 等
如果你的目录结构比较简单，内容也比较少，其实也可以不使用 pkg 目录，直接把上面的这些包放在最上层即可
|- test                  # 测试资源层






项目简介
Author：sophiemarceau_qu



在 api 目录当中我们会按照 product name/app name/版本号/app.proto 的方式进行组织
具体怎么组织可能每个公司都不太一样，但是总的来说就是应用的 唯一名称+版本号 来进行一个区分

API 命名规范
包名
产品名	product
应用名	app
版本号	v1
包名	product.app.v1
目录结构	api/product/app/v1/xx.proto

API 定义
命名规则：方法 + 资源
标准方法：参考 Google API 设计指南
标准方法	HTTP 映射
List	GET
Get	GET
Update	PUT 或者 PATCH
Create	POST
Delete	DELETE
除了标准的也有一些非标准的，例如同步数据可能会用 Sync 等，不过大部分的 api 应该都是标准的

// api/product/app/v1/blog.proto

syntax = "proto3";

package product.app.v1;

import "google/api/annotations.proto";

// blog service is a blog demo
service BlogService {

	rpc GetArticles(GetArticlesReq) returns (GetArticlesResp) {
		option (google.api.http) = {
			get: "/v1/articles"
			additional_bindings {
				get: "/v1/author/{author_id}/articles"
			}
		};
	}
}


/cmd
我们一般采用 /cmd/[appname]/main.go 的形式进行组织

首先 cmd 目录下一般是项目的主干目录
这个目录下的文件不应该有太多的代码，不应该包含业务逻辑
main.go 当中主要做的事情就是负责程序的生命周期，服务所需资源的依赖注入等，其中依赖注入一般而言我们会使用一个依赖注入框架，这个主要看复杂程度，后续会有一篇文章单独介绍这个

/internal
internal 目录下的包，不允许被其他项目中进行导入，这是在 Go 1.4 当中引入的 feature，会在编译时执行

所以我们一般会把项目文件夹放置到 internal 当中，例如 /internal/app
如果是可以被其他项目导入的包我们一般会放到 pkg 目录下
如果是我们项目内部进行共享的包，而不期望外部共享，我们可以放到 /internal/pkg 当中
注意 internal 目录的限制并不局限于顶级目录，在任何目录当中都是生效的

/pkg
一般而言，我们在 pkg 目录下放置可以被外部程序安全导入的包，对于不应该被外部程序依赖的包我们应该放置到 internal 目录下， internal 目录会有编译器进行强制验证

pkg 目录下的包一般会按照功能进行区分，例如 /pkg/cache 、 /pkg/conf 等
如果你的目录结构比较简单，内容也比较少，其实也可以不使用 pkg 目录，直接把上面的这些包放在最上层即可
一般而言我们应用程序 app 在最外层会包含很多文件，例如 .gitlab-ci.yml makefile .gitignore 等等，这种时候顶层目录会很多并且会有点杂乱，建议还是放到 /pkg 目录比较好
