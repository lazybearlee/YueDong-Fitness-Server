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

## 3. 项目结构

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