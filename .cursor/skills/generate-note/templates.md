# ReleaseNote 模板

生成或补充 `notes/release/ReleaseNotes_{tag}.md` 前阅读本文件。节名与 `notes/release/ReleaseNotes.md` 一致；文风与现有 `notes/release/ReleaseNotes_v*.md` 一致。

结构参考 [quic-go v0.62.0](https://github.com/quic-go/quic-go/releases/tag/v0.62.0)：开头亮点总述，再分 Notable / Breaking / Fixes，最后是完整 Changelog、New Contributors 与 Full Changelog 对照。

## 新文件结构

```markdown
## Release Notes

+ 一两句总述（自上一版本到本 tag 的要点，对应发布页开头亮点）。

### Known Issues

### Notable Changes

### Improvements

### Breaking Changes

### API Changes

### Changes

### Notable Fixes

### Fixes

### Changelog

### New Contributors

**Full Changelog**: {RANGE_START}...{tag}

## Library Changes

### library Updated

#### Updated

#### No Longer Available

#### Added
```

## 填写规则

- 不要写 `# ExcelExporter` 大标题，文档从 `## Release Notes` 起头
- 标题用模板英文；条目用中文，`+ ` 开头
- 模板中部分条目没有内容时，生成的文档中可以不包含该条目：无条目的 `###` / `####` 整节不要写出；`## Library Changes` 下没有任何库变更时整节不要写出。不要留空标题
- 总述对应发布页开头：写本版本最重要的一两件事，不要把下面条目再抄一遍
- **Notable Changes**：最值得关注的能力或行为（亮点列表，不必覆盖全部提交）
- **Improvements**：其余新能力、脚本、CI、文档完善
- **Breaking Changes**：不兼容变更（升级语言版本、改签名、删除 API 等）
- **API Changes**：公开 API 增删改（若已写入 Breaking Changes 且无更多条目，可只保留一处）
- **Changes**：对外行为、配置、默认策略等非修复变更
- **Notable Fixes**：重要缺陷修复（不必列出全部）
- **Fixes**：其余修复；与 Notable Fixes 重复的不要再写一遍
- **Changelog**：范围内全部提交的完整列表（主题 + 短哈希或 PR 号），对应发布页 Changelog
- **New Contributors**：该范围内首次出现的提交者；无则省略
- **Full Changelog**：上一 tag 与本 tag 的对照。有 GitHub 远程时写成 `https://github.com/{owner}/{repo}/compare/{RANGE_START}...{tag}`，否则写成 `{RANGE_START}...{tag}`。无上一 tag 则省略该行
- `Library Changes`：只写 `go.mod` 里实际变化的模块与版本（`旧 → 新`）
- `### library Updated` 中的 `library` 换成真实模块名，或不用该占位、直接在 `## Library Changes` 下写 `+ 模块：旧 → 新`
- **Known Issues**：明确未解决的限制；无则省略
