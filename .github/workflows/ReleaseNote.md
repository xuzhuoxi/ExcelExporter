# ReleaseNote 说明

本文说明 `.github/workflows/ReleaseNote.yml` 的用途、操作步骤和注意事项。

只更新 **已有** GitHub Release 的正文，不重新构建、不改附件。发版请使用 `Release.yml`（见同目录 `Release.md`）。

不需要手动配置 `GITHUB_TOKEN`。GitHub 会为每次运行注入临时令牌；workflow 已声明 `contents: write`。

## 1. 做什么

从当前所选分支读取 `notes/release/ReleaseNotes_<tag>.md`，用文件全文覆盖对应 GitHub Release 的说明。

适合：打 tag 发版之后才补写或修改了说明文件，需要把正文同步到已有 Release。

## 2. 能更新与不能更新

工作流只改 GitHub Release 的 **说明正文**。成功时用 `ReleaseNotes_<tag>.md` 全文覆盖，不会自动发布 Draft，也不会改其它字段。

### 能更新

| 对象 | 说明 |
| --- | --- |
| 已发布的 Release | 远程已有该 git tag，且已有正式 Release（`Release.yml` 打出来的属于这类） |
| Draft Release | 远程 **已经有** 同名 git tag，Release 处于草稿状态；更新后 **仍是 Draft** |
| 打 tag 之后才写入/修改的说明文件 | 文件从 Run workflow 所选分支读取，不要求文件存在于 tag 指向的旧提交 |

### 不能更新

| 对象 | 说明 |
| --- | --- |
| 尚不存在的 Release | 不会新建 Release，请先用 `Release.yml` 发版 |
| 没有对应 git tag 的 Draft | 网页上存成草稿、tag 还未推到远程时，会在「找不到 tag」处失败 |
| 不存在的 tag / 填错的 tag | 会失败，日志中列出远程已有的 `v*` tag |
| 所选分支上没有 `ReleaseNotes_<tag>.md` | 会失败 |
| 附件、标题、prerelease 标记 | 保持原样，本工作流不改 |
| Draft → 正式发布 | 不会把 Draft 发布成正式版 |
| 二进制重新构建 | 不构建、不上传附件；改产物请重跑 `Release.yml` |

仓库若开启了 **immutable releases**，已发布 Release 的附件不能改；本工作流只改正文，一般仍可更新说明。若 GitHub 对不可变 Release 连正文也锁定，则会权限/API 失败。

## 3. 触发条件

仅 `workflow_dispatch`，只能在 GitHub 网页上手动运行，推送代码或打 tag **不会**触发。

Run workflow 弹窗里有 **两处** 选择，含义不同：

| 控件 | 选的是什么 | 本工作流应怎么选 |
| --- | --- | --- |
| **Use workflow from** | 用哪一次提交上的 workflow 文件，以及从哪次提交读取 `ReleaseNotes_*.md` | 选 **分支**（一般为 `main`），不要选 tag |
| **tag 输入框** | 要更新哪一个 GitHub Release | **手填** tag 名，例如 `v1.1.0` |

**Use workflow from** 点开后，和仓库里切换分支一样，顶部可以切 **Branches / Tags**。切到 Tags 就能选 tag，但那只表示「在这个 tag 的代码上跑工作流」。该 tag 提交里往往还没有后来补的说明文件，本工作流会读不到 `ReleaseNotes_*.md`。

**tag 输入框** 不能变成仓库 tag 的动态下拉列表。`workflow_dispatch` 的 `choice` 只能写死选项，没有「列出全部远程 tag」这种类型，所以必须手填。

## 4. 运行前检查

- `ReleaseNote.yml` 已合入要运行的分支（一般为 `main`），否则 Actions 列表里看不到。
- 目标 tag 已推送到远程，且已有对应的 GitHub Release（正式版或已关联该 tag 的 Draft 均可）。
- `notes/release/ReleaseNotes_<tag>.md` 已提交到即将选择的分支。文件名必须与 tag **完全一致**，例如 tag `v1.1.0` 对应 `notes/release/ReleaseNotes_v1.1.0.md`。

## 5. 操作步骤

1. 将说明文件提交并推到目标分支（一般为 `main`）。
2. 打开 **Actions → ReleaseNote → Run workflow**。
3. **Use workflow from** 选包含该说明文件的 **分支**（一般为 `main`）。点开后若出现 Tags，不要用它来指定要更新的 Release。
4. 在 **tag** 输入框 **填写** 已有 tag（例如 `v1.1.0`）。也可写成 `refs/tags/v1.1.0`，工作流会去掉前缀。
5. 运行成功后，到 **Releases** 页确认正文已换成该文件内容。

写入的是文件全文，**不会**再追加 GitHub 自动生成的 notes。附件、标题、prerelease 标记保持不变。

## 6. 工作流执行步骤

| 步骤 | 行为 |
| --- | --- |
| checkout | 检出 Run workflow 时选择的分支（浅克隆），并拉取 tag |
| Update GitHub Release notes | 校验 tag、Release、说明文件均存在后，`gh release edit --notes-file` |

说明文件来自 **所选分支的当前提交**，不是 tag 指向的旧提交。因此打 tag 之后再改 `ReleaseNotes_<tag>.md` 并合入该分支，即可被读到。

## 7. 注意事项

1. **没有对应 Release 时不会创建**，工作流会失败。请先用 `Release.yml` 发版。
2. **填错 tag** 会失败，日志中会列出远程已有的 `v*` tag。
3. 当前分支没有对应说明文件会失败。确认 **Use workflow from** 选对了分支。
4. 发布者仍会显示为 `github-actions[bot]`。

## 8. 常见失败

| 现象 | 可能原因 |
| --- | --- |
| Actions 里没有 ReleaseNote | 工作流文件尚未合入默认分支 |
| 找不到 tag | tag 名填错，或不存在于远程 |
| 找不到 GitHub Release | 该 tag 还没有发过版 |
| 找不到 tag（Draft） | 网页上的草稿尚未关联已推送的 git tag |
| 找不到说明文件 | 文件名与 tag 不一致，或文件不在所选分支上 |
| 权限错误 | 仓库/组织限制了 `GITHUB_TOKEN` 写权限 |

## 9. 相关路径

| 路径 | 用途 |
| --- | --- |
| `.github/workflows/ReleaseNote.yml` | 手动更新已有 Release 正文 |
| `notes/release/ReleaseNotes.md` | 手写说明模板（不会被自动读取） |
| `notes/release/ReleaseNotes_<tag>.md` | 对应 tag 的正式说明 |
| `.github/workflows/Release.yml` | 发版工作流 |
