# chatbot

基于[openwechat](https://github.com/eatmoreapple/openwechat)，参考[djun](https://github.com/djun/wechatbot)

## 使用配置文件`config`

将`config\config.template.json`复制到同目录下，命名为`config.json`

## 新增提示词

在`config\config.json`的字段`promptPaths`里新增字段

```json
"promptPaths": {
	"PromptPath": "relative_path",
	"NewPromptPath": "new_relative_path",
	"注意一致": "another_path"
}
```

和在`prompts\internal.go`新增常量

```go
const (
	Prompt = "PromptPath"
	NewPrompt = "NewPromptPath"
	BPath = "注意一致"
)
```

两个<u>注意一致</u>应**保持一致**

## logger

logger作为参数不传整个logger而是传细分的logger是依赖最小化原则