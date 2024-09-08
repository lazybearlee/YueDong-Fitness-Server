# YueDong-Fitness-Server

## 1. 项目简介

本项目是一个运动健康类的APP项目，主要功能包括用户注册登录、个人信息管理、运动数据记录、运动数据展示、个人健康状况查看等功能。本项目是后端部分，采用Gin+Gorm框架，数据库采用MySQL。

## 2. 项目结构

```
YueDong-Fitness-Server
├─annex
├─api
│  └─v1
│      ├─app
│      └─system
├─config
├─core
│  └─initialize
│      ├─app
│      └─system
├─deploy
├─docs
├─global
├─log
│  ├─2024-08-25
│  ├─2024-08-26
│  ├─2024-08-27
│  └─2024-08-28
├─middleware
├─model
│  ├─app
│  │  ├─request
│  │  └─response
│  ├─common
│  │  ├─gc
│  │  ├─request
│  │  └─response
│  └─system
│      ├─request
│      └─response
├─oss
│  └─uploads
├─router
│  ├─app
│  └─system
├─service
│  ├─app
│  ├─oss
│  └─system
├─tasks
└─utils
    └─timer
```

## 3. 项目运行说明

### 3.1 基础环境配置

- Go 1.18+ 版本，本项目采用 Go 1.22 版本。
- MySQL 8.0+ 版本，本项目采用 MySQL 8.0.21 版本。

### 3.2 项目配置

- 在项目根目录下复制 `config_backup.yaml` 文件，并重命名为 `config.yaml`。
- 修改 `config.yaml` 文件中的数据库配置。
  - `mysql`：数据库连接配置。修改path、username、password等字段。
- 确保 `system` 中的 `mysql-init-data` 字段配置正确。
  - 如果初次运行项目，需要将 `mysql-init-data` 字段设置为 `true`，项目会自动初始化数据库。
- 如果需要使用云服务器部署，请修改 `system` 的 `addr` 和 `port` 字段。
  - `addr`：服务器地址。(一般为服务器内网网址)

### 3.3 项目运行

- 确保 MySQL 服务已启动，但不需要手动创建数据库和表，项目会自动初始化。
- 在项目根目录下执行 `go run main.go -c config.yaml` 命令运行项目。
- 也可以使用 `go build` 命令编译项目，然后执行编译后的文件。
- 也可以使用 `go build -o <filename>` 命令编译项目，生成指定文件名的可执行文件。
- 如果是 `Windows` 系统，可以在命令行输入 `.\server_windows64.exe -c config.yaml` 运行项目。
- 如果是 `Linux` 系统，可以使用 `nohup ./server_linux64 -c config.yaml &` 命令后台运行项目。(需要先给文件执行权限)
- 如果不能直接运行，请尝试在本机安装 `Go` 环境，然后运行项目。

## 4. git 使用规范

### 4.1 commit message 规范

```git
<type>: <subject>
```

#### 1. type

用于说明 commit 的类别，只允许使用下面7个标识。

- `feat`：新功能（feature）
- `fix`：修补bug
- `docs`：文档（documentation）
- `style`： 格式（不影响代码运行的变动）
- `refactor`：重构（即不是新增功能，也不是修改bug的代码变动）
- `test`：增加测试
- `chore`：构建过程或辅助工具的变动

#### 2. subject

是 commit 目的的简短描述，不超过50个字符。

### 4.2 分支管理规范

- `main` 分支：主分支，只能用来发布新版本，不能在上面干活。
- `develop` 分支：开发分支，用于存放临时的开发版本。

采用 Feature Branching 模型，采用多分支进行管理。

#### 4.2.1 项目管理

- 首先从 main 分支拉取最新代码：`git pull origin main`
- 新建分支：`git checkout -b develop/<branch-name>`
- 修改、新增或删除代码
- 提交代码：`git add .`、`git commit -m "<type>(<scope>): <subject>"`
- 推送分支：`git push origin develop/<branch-name>`
- 提交 PR

### 4.3 PR 规范

- PR 标题：简洁明了，包含了本次 PR 的目的。
- PR 描述：详细描述本次 PR 的内容。
- Review：指派给相关人员进行 Review。
- Assignees：指派给相关人员进行处理。