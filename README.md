#shadowdemo

### 介绍
> 
> shadowdemo 是一个基于gin+grom一个快速开发得单体应用代码，里面集成了
> 
> 基于casbin的权限认证(middleware.Authorizer("APPID"))
> 
> 用户登录校验(securitymiddleware.UsernamePasswordLoginFilter(loginPath))
>
> 基于redis管理的jwt认证(middleware.JWTParse(tnb9Y0du$2a$10$KmatydruRTKlaUwErUOtNOXiPHVPunb9Y0dup9newm, 360000))
>
> ip白名单的过滤(middleware.VisitSettingFilter(server.ServerConfig.ServerName, service.NewVisitSettingService().FrontVisitFilter))
>
> 同ip的限流(concurrentlimit.ConcurrentLimit(100, 5*time.Minute))
>
> 谷歌验证码(middleware.GoogleTokenValidator(loginPath, getProfileID))
>
### 目录讲解
> 
> ACCServer 具体服务名称，可以有多个，具体的端口配置信息在 server.json中
> 
> config 配置信息，里面包含有国际化配置，权限配置，验证码的配置
> 
> env 线上发布所引用的配置文件
> 
> middleware 项目中所要用到的中间件
> 
> model 实体类层，其中包含了entity，以及dao 
> 
> service 具体业务实现类
> 
> shadow-framework 框架（里面包含了数据库的实例化、oss的实例化、全局主键生成、日志、登录的校验、casbin的校验，server.jSON 文件解析，还有一些工具类）
> 
> sql 发布文件sql 初始化文件 
>

### 开发步骤
> 
> 1、配置server.json需要的项目配置信息，例如端口（Port） 数据库配置信息等
> 
> 2、如果多加了一个（"AdminServer": {
        "Host": "http://localhost:8031",
        "Port": ":8031",
        "ContextPath": "/acc"
    },）  则需要在shadow-framework server中配置对应的结构体，在项目中可引用（server.ServerConfigInstance().AdminServer.Host）
> 
> 实体类定义在model/do 中进行定义，每个文件定义一个结构体,然后在外面执行 go run .\model\daogen\main\main.go ，会生成对应的dao方法
> 
> 3、router定义在对应服务ACCServer中，例如增加admin_router.go ,然后在main_router.go中增加对应方法 adminRouter(router)
> 
> 4、启动 go run ACCServer/main.go


### 新增的APiserver讲解
> 
> 由于项目中经常要被第三方来进行接入，故接入了一套api的方法
> 
> 实现了三方信息的管理，保存在merchant中，生成信息的时候会给三方生成对应的key
> 
> 实现了三方请求日志的管理，做的比较简单可以根据自己的一些规则来进行扩展


> 
> 具体接口 http://127.0.0.1:8030/api/v1/action

> 
> 例如一个名字叫test的商户进行接入（aes_key:g3im_lF46_h6eRmU1FmyS4AKVBZ8QWwo ,md5:tbKAdNFwroxe5acEee98T0ZZXenoXC18sdtr4lwIILY=）

> 
>接入查询银行信息接口 UserLevel=1 用 aes_key 进行加密后，转hex 得到一个data（331E1E03450D2CFBFB35FB05C9B4D12E32BEF313DBE28C70）


> 
>对公共参数进行取MD5,其中sign 不加入
> 
>	vo := vo.APIPublicRequest{
> 
>		MerchantNo: "test",
> 
>		Action:     "query_banklist",
> 
>		Data:       "331E1E03450D2CFBFB35FB05C9B4D12E32BEF313DBE28C70",
> 
>		Sign:       "",
> 
>	}
> 
>
> 
>下面这一串取md5（按照首字母大小顺序）
> 
>Action=query_banklist&Data=331E1E03450D2CFBFB35FB05C9B4D12E32BEF313DBE28C7047F95661AC963263&MerchantNo=liujian&Key=tbKAdNFwroxe5acEee98T0ZZXenoXC18sdtr4lwIILY=
> 
>
> 
>得到一个sign 4D42A083D7794B8116B1D5DFCDCF6CD8
>
> 
>然后把
>
>       MerchantNo: "test",
> 
>		Action:     "query_banklist",
> 
>		Data:       "331E1E03450D2CFBFB35FB05C9B4D12E32BEF313DBE28C70",
> 
>		Sign:       "4D42A083D7794B8116B1D5DFCDCF6CD8",
> 
>        
> 
>用表单格式请求过来






