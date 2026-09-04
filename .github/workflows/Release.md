# Release 说明

本文说明 `.github/workflows/Release.yml` 的触发条件、人工操作步骤、工作流内部流程，以及发版时需要注意的事项。

## 1. 做什么

在 `main` 上推送符合规则的 tag 后，自动：

1. 确认该 tag 指向的提交位于 `main`
2. 交叉编译各平台二进制
3. 打包环境配置 `env_<tag>.zip`
4. 打包并生成校验和
5. 创建（或更新附件的）GitHub Release

不需要手动配置 `GITHUB_TOKEN`。GitHub 会为每次运行注入临时令牌；workflow 已声明 `contents: write`，用于创建 Release 和上传附件。

可调整常量在 `Release.yml` 顶部的 `env`：`APP_NAME`（产物与二进制名）。

## 2. 触发条件

```yaml
on:
  push:
    tags:
      - 'v*.*.*'
```

- 仅响应 **推送 tag**，不响应向 `main` 的普通 commit。
- tag 名必须匹配 `v*.*.*`，例如 `v1.0.3`、`v1.0.3-rc.1`。
- 下列名称 **不会触发**：`1.0.3`（缺 `v`）、`v1.0`（段数不够）、`release-1.0.3`。

GitHub 无法在 `on.push` 里把 `branches` 和 `tags` 组合成「只在 main 上打 tag」：两者是 **或** 关系，且分别匹配 `refs/heads/*` 与 `refs/tags/*`。因此本 workflow 用 `tags` 触发，再在 job 内校验提交是否在 `main` 上。

## 3. 发版前检查

- `Release.yml` 以及本次要发布的代码 **已经合入 `main`**。GitHub 使用的是 **被 tag 的那次提交** 上的 workflow 文件。
- 仓库已启用 Actions。
- **Settings → Actions → General → Workflow permissions** 允许 workflow 申请写权限（文件内已声明 `contents: write`）。
- 没有规则集禁止创建 `v*` tag。

## 4. 人工操作流程

以下命令均在已包含待发布提交的 `main` 上执行。

### 4.1 准备说明（建议）

按 tag 名新增说明文件，必须与 tag **完全一致**：

```text
notes/release/ReleaseNotes_<tag>.md
```

例如 tag `v1.0.3` 对应 `notes/release/ReleaseNotes_v1.0.3.md`。

可参考模板 `notes/release/ReleaseNotes.md`。该模板本身 **不会** 被 workflow 读取。

Release 正文规则：

- 若该文件存在：先写入文件全文（与 GitHub 上 v1.0.2 的手写说明相同），再追加 GitHub 自动生成的 notes。
- 若该文件存在且已含 `**Full Changelog**`：仍追加自动 notes，但去掉其中的 `**Full Changelog**` 行，避免与手写说明重复。
- 若该文件不存在：只使用 GitHub 自动生成的 notes（已合并 PR 列表、新贡献者、Full Changelog 链接）。

直接推到 `main`、未走 PR 的 commit 不会出现在「What's Changed」条目中。说明文件必须已经包含在被 tag 的那次提交里。

### 4.2 打 tag 并推送

```sh
git checkout main
git pull origin main
git tag v1.0.3
git push origin v1.0.3
```

不要在其它分支上打即将发版的 tag。若该提交不在 `main` 上，workflow 会在「Ensure tag is on main」失败退出。

### 4.3 确认结果

1. 仓库 **Actions** 中查看 `Release` 工作流是否成功。
2. 仓库 **Releases** 中确认标题、说明、附件。

## 5. 工作流执行步骤

| 步骤 | 行为 |
| --- | --- |
| checkout | 完整克隆，以便校验 tag 与 `main` 的祖先关系 |
| Ensure tag is on main | `git merge-base --is-ancestor $GITHUB_SHA origin/main`，不在 `main` 则失败 |
| setup-go | 使用 `go.mod` 中的 Go 版本 |
| Build release binaries | 交叉编译、打包环境配置、生成 `SHA256SUMS.txt` |
| Create GitHub Release | 组装说明（手写文件 + 自动 notes）、创建 Release；若已存在则覆盖附件并更新正文 |

`GITHUB_REF_NAME` 在本 workflow 中等于 **tag 短名**（如 `v1.0.3`），不是 `refs/tags/v1.0.3`。产物名、Release 标题、说明文件路径都使用该值。

## 6. 构建产物

平台：`linux` / `windows` / `darwin` × `amd64` / `arm64`（共 6 个）。

各平台包内包含：可执行文件、`LICENSE`、`README.md`、`README_EN.md`。

另附一份与平台无关的环境配置包 `env_<tag>.zip`。解压后根目录为 `evn_<tag>/`，内容来自仓库 `res/`：

- 目录：`db`、`lang`、`proxy`、`source`、`template`、`target`
- 文件：`excel.yaml`、`project.yaml`、`system.yaml`

`target` 只保留空目录，即使 `res/target` 下已有文件或子目录也不会打进包。不包含 `res/test`、日志或本地可执行文件。

| 平台 | 文件名示例 |
| --- | --- |
| Windows | `ExcelExporter_v1.0.3_windows_amd64.zip` |
| 其它 | `ExcelExporter_v1.0.3_linux_amd64.tar.gz` |
| 环境配置 | `env_v1.0.3.zip`（内含 `evn_v1.0.3/`） |
| 校验和 | `SHA256SUMS.txt` |

二进制 **未** 写入版本号；程序内没有独立的 `-version` 输出。

## 7. Pre-release

tag 名中含 `-` 时（如 `v1.0.3-rc.1`），创建的 GitHub Release 会标记为 **prerelease**。正式版使用不含 `-` 的 tag，如 `v1.0.3`。

## 8. 注意事项

1. **先合 `main`，再打 tag。** 只在功能分支上改 `Release.yml` 或说明文件，然后对该分支提交打 tag，会因不在 `main` 上而失败；即使校验放宽，用到的也是旧提交上的 workflow。
2. **不要把 `branches: [main]` 和 `tags` 写在一起指望变成「main 上的 tag」。** 那会变成：推 `main` **或** 推任意匹配 tag 都会跑，普通提交也会误触发发版。
3. **同一 tag 重跑** 会覆盖附件，并用当前提交里的说明文件重新写入 Release 正文。说明文件若在打 tag 之后才加入，重跑时 checkout 的仍是旧提交，读不到新文件。
4. **已存在的 tag 再 `git push` 不会再次触发。** 需要新版本时打新 tag。若必须复用同一 tag，需先处理远程 tag 与已有 Release，操作不可逆，应谨慎。
5. **没有人工审批。** tag 推送成功且校验通过后会直接发布，不会先做成 draft。
6. **依赖均为公开模块**（如 `infra-go`），构建不需要额外 private token。

## 9. 常见失败

| 现象 | 可能原因 |
| --- | --- |
| 推了 tag 但没有出现 Release 工作流 | tag 不符合 `v*.*.*`；或 Actions 未启用 |
| Ensure tag is on main 失败 | tag 打在非 `main` 提交上；远程 `main` 尚未包含该提交 |
| 创建 Release 权限错误 | 仓库/组织限制了 `GITHUB_TOKEN` 写权限 |
| Release 正文没有手写说明 | 缺少 `notes/release/ReleaseNotes_<tag>.md`，或文件名与 tag 不一致，或该文件不在被 tag 的提交中 |
| 重跑后仍看不到新写的说明 | 说明文件是在打 tag 之后才提交的，重跑不会读到后续 commit |

## 10. 相关路径

| 路径 | 用途 |
| --- | --- |
| `.github/workflows/Release.yml` | 发版工作流 |
| `notes/release/ReleaseNotes.md` | 手写说明模板（不会被自动读取） |
| `notes/release/ReleaseNotes_<tag>.md` | 对应 tag 的正式说明 |
