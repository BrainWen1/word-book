word-book/
├── cmd/
│   └── server/
│       └── main.go                # 程序入口，加载配置，启动服务器
├── internal/
│   ├── config/
│   │   └── config.go              # 读取.env 配置到config.AppConfig
│   ├── model/                     # 数据模型定义（GORM 模型）
│   │   ├── user.go
│   │   └── word.go
│   ├── repo/                      # 数据访问层
│   │   ├── user_repo.go
│   │   └── word_repo.go
│   ├── service/                   # 业务逻辑层
│   │   ├── user_service.go        # 注册、登录
│   │   ├── word_service.go        # 生词本
│   │   └── dict_service.go        # 查词（缓存）
│   ├── handler/                   # HTTP 处理器
│   │   ├── user_handler.go
│   │   ├── word_handler.go
│   │   ├── dict_handler.go
│   │   └── middleware/
│   │       ├── auth.go            # 认证中间件
│   │       └── cors.go            # 跨域中间件
│   ├── router/
│   │   └── router.go              # 初始化各层组件，路由注册
│   ├── infra/                     # 基础设施层（数据库连接、缓存实现、外部 API 调用）
│   │   ├── database/
│   │   │   └── mysql.go           # 连接数据库、自动迁移函数
│   │   ├── cache/
│   │   │   ├── cache.go           # Cache 接口定义
│   │   │   └── redis_cache.go     # Redis 实现
│   │   └── external/
│   │       └── search_word.go     # 调用外部词典 API
│   ├── util/                      # 工具函数（JWT 生成、统一响应封装）   
│   │   ├── jwt/
│   │   │   └── jwt.go             # JWT 生成和解析验证函数
│   │   └── response/
│   │       └── response.go        # 统一响应结构体和函数
|   └── webapp/                    # 前端资源
|       ├── index.html             # 前端页面
|       └── webapp.go              # 嵌入前端资源，提供静态文件服务
├── docs/                          # API 文档（Swagger 注释生成）
│   └── docs.go
|   └── swagger.json
|   └── swagger.yaml
├── go.mod
├── go.sum
├── README.md                      # 项目说明文档
├── .env                           # 环境变量配置
├── .gitignore
└── project_struc.md               # 项目结构说明