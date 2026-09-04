## Release Notes

+ 本版本将构建要求提升到 Go 1.24.0，适配 Go module 交叉编译，并补齐本地构建脚本与 GitHub CI / 发版工作流（含环境配置包）。

### Known Issues

+ yaml, toml, hcl, env, properties 数据导出时，key 会转为**小写**，本意要求**大小写相关**。
+ project.yaml 中 encoding 与 buff 相关的配置未实现。
+ C++ 表头模板未完善，C++ 常量模板未完善。

### Notable Changes

+ 最低 Go 版本升级为 1.24.0，并按 Go module 模式编译。
+ 新增 `build/build.sh` / `build/build.bat`：单元测试、多平台交叉编译可执行文件、打包源码并嵌入版本信息。
+ 新增 GitHub Workflows：`CI.yml`（构建与测试）、`Release.yml`（打 tag 发版，并打包 `env_<tag>.zip`）、`ReleaseNote.yml`（更新已有 Release 正文）。

### Improvements

+ 添加 Cursor Skill `generate-note`，用于根据提交范围生成 Release 说明。
+ 更新 Release 说明模板 `notes/release/ReleaseNotes.md`。
+ `.gitignore` 增加 `go.work`、`go.work.sum`。
+ 部分依赖本机路径的测试改为条件测试，CI 可用 `-tags=skiptest` 跳过。
+ 构造脚本增加打包源代码的功能。

### Breaking Changes

+ 构建所需 Go 版本由 1.16 提升到 1.24.0。
+ `go.mod` 移除 `replace` 指令，`infra-go` 改为直接使用模块版本。

### Changes

+ Release 发版时额外上传环境配置包 `env_<tag>.zip`，解压后根目录为 `evn_<tag>/`；其中 `target` 始终为空目录。

### Fixes

+ 修复构造脚本中多出的 `-ldflags`。
+ 修正构造脚本的打印信息。

### Changelog

+ 升级 Go 到 1.24，适配 Go module 编译并更新构造脚本（`5e562d5`）
+ 优化构造脚本，嵌入版本信息（`111da5c`）
+ 修复构造脚本多出一个 `-ldflags`（`2c55ae5`）
+ 构造脚本添加打包源代码的功能（`9caa3f6`）
+ 修正构造脚本的打印信息（`90241b5`）
+ 更新 go.mod 中 `infra-go` 依赖（`606688f`）
+ 部分测试改为条件测试，并将 `infra-go` 更新为 v1.4.1（`dbb05c5`）
+ 添加 generate-note Skill 与 GitHub Workflows：CI、Release、ReleaseNote（`a547523`）
+ Release 流程增加环境配置 zip 打包与上传（`51d5408`）
+ Release 环境配置包中的 `target` 保持为空目录（`f19c185`）

**Full Changelog**: https://github.com/xuzhuoxi/ExcelExporter/compare/v2.2...v2.2.1

## Library Changes

+ `github.com/xuzhuoxi/infra-go`：v1.0.4 → v1.4.1
+ `golang.org/x/crypto`：v0.0.0-20220411220226-7b82a4e95df4 → v0.48.0
+ `golang.org/x/net`：v0.0.0-20220412020605-290c469a71a5 → v0.50.0
+ `golang.org/x/text`：v0.3.7 → v0.34.0
