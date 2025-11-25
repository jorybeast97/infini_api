# Infini API

轻量分层的 Go 后端，支持本地使用 Gin 运行与在 Vercel 以 Serverless Functions 部署。为 `infini_web` 前端提供博客、作者、应用与地图聚合等接口。

## 技术栈
- Go 1.19+（本地）/ Vercel @vercel/go（Serverless）
- Gin（HTTP 路由与中间件）
- 分层架构：API / Service / DAO / Domain / Server / Middleware

## 目录结构
```
infini_api/
├─ main.go                  # 本地开发入口（Gin）
├─ go.mod / go.sum          # 模块与依赖
├─ vercel.json              # Vercel 部署与路由配置
├─ api/
│  └─ server.go             # Vercel Serverless 入口（export Handler）
└─ src/
   ├─ api/                  # API 层（Gin Handler，每个端点一个文件）
   │  ├─ auth/login.go
   │  ├─ authors/list.go
   │  ├─ posts/list.go
   │  └─ hello/hello.go
   ├─ service/              # Service 层（业务逻辑与校验）
   │  ├─ auth/
   │  ├─ authors/
   │  ├─ posts/
   │  └─ photos/
   ├─ dao/                  # DAO 层（仓储接口与实现）
   │  ├─ repo.go            # 仓储接口定义
   │  └─ memory/            # 并发安全内存实现（可替换成 DB）
   ├─ domain/               # 领域模型与通用类型
   ├─ server/               # 依赖注入与路由构建
   │  ├─ di.go
   │  └─ router.go
   ├─ middleware/           # 中间件（CORS 等）
   └─ response/             # 统一 JSON 响应（扩展位）
```

## 快速开始（本地开发）
1. 安装依赖并启动：
   - `cd infini_api`
   - `go mod tidy`
   - `go run .`
2. 验证端点：
   - `GET /api/hello`
   - `POST /api/auth/login`（Body：`{"username":"admin","password":"password"}`）
   - `GET /api/authors`
   - `GET /api/posts?sort=date:desc&page=1&limit=2`

## Vercel 部署
- 入口：`api/server.go`（包名 `handler`，导出 `Handler(w http.ResponseWriter, r *http.Request)`）
- 配置：`vercel.json`
```
{
  "version": 2,
  "builds": [{
    "src": "api/*.go",
    "use": "@vercel/go",
    "config": { "goVersion": "1.23" }
  }],
  "routes": [{ "src": "/(.*)", "dest": "api/server.go" }],
  "env": { "APP_ENV": "production" },
  "buildCommand": "go mod tidy"
}
```
- 部署方式：
  - 控制台：关联 Git 仓库 → 新建项目（Framework: Other）→ Deploy
  - CLI：`npm i -g vercel` → `vercel login` → `vercel` / `vercel --prod`
- 本地模拟：
  - `npm i -D @vercel/go`
  - `vercel dev`（默认 3000）

## 配置与环境变量
- `PORT`：本地监听端口（默认 `8080`）
- `INFINI_SECRET`：JWT 签名秘钥（必配于生产环境）
- `UPLOAD_DIR`：静态上传目录（本地为 `uploads`；在 Vercel 建议改为 `public/uploads` 或使用外部对象存储）
- `APP_ENV`：运行环境（如 `production`），可在 `vercel.json` 或 Vercel 控制台配置

## 分层说明
- API 层（`src/api/*`）：HTTP 参数解析、调用 Service、返回 JSON；一个端点一个文件，分包管理
- Service 层（`src/service/*`）：业务规则、校验（长度、URL、坐标范围、状态枚举、伙伴存在性等）、组合查询
- DAO 层（`src/dao/*`）：数据访问接口与实现（当前为内存实现，后续可替换为 `sqlite/ postgres` 等）
- Domain（`src/domain/*`）：领域模型与分页类型
- Server（`src/server/*`）：依赖注入与 Gin 路由构建
- Middleware（`src/middleware/*`）：跨域、日志、鉴权等（已提供 CORS）

## 已提供端点
- `GET /api/hello`：联通性测试，返回 `{"message":"hello world"}`
- `POST /api/auth/login`：登录，返回 `{ token, user }`
- `GET /api/authors`：作者列表，支持 `q,page,limit`
- `GET /api/posts`：文章列表，支持 `q,status,hasLocation,sort,page,limit`

## 与前端协作
- 字段与结构与 `infini_web` 的 `types.ts` 保持一致（`BlogPost`、`Photo`、`Author`、`AppProject` 等）
- 分页响应统一为 `{ data, meta }`
- 错误响应建议统一 `{ error: { code, message, details? } }`

## 生产建议
- 设置可信代理（Gin `TrustedProxies`）与 Release 模式
- 在 Vercel 上静态资源建议移至 `public/` 或采用外部对象存储（例如 Vercel Blob / S3），上传接口返回签名 URL 或元数据
- 将敏感配置（秘钥、私有包访问）置于项目 Settings → Environment Variables；如需安装私有包，配置 `GIT_CREDENTIALS`

## 开发脚本
- 启动（本地）：`go run .`
- 编译检查：`go build ./...`

---
如需继续完善其余端点（apps、photos、preview、search、map、uploads、stats）或接入数据库实现（`sqlite/postgres`），可以在现有分层上扩展并通过配置切换实现。 
