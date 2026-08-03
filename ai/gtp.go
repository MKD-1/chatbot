package gtp

import (
	"context"
	"os"

	"github.com/MKD-1/chatbot/logger"
	"github.com/eatmoreapple/openwechat"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func Test(content string) {
	appLog := logger.New()
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("BASE_URL")),
	)
	completion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model: os.Getenv("Model"),
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(content),
		},
	})
	if err != nil {
		panic(err)
	}

	appLog.Info(completion.Choices[0].Message.Content)

}
func Talk(msg *openwechat.Message) string {
	systemPrompt := `
	你是一个微信聊天机器人，不是 AI 助手。
	你在一个微信群中。先判断是否值得主动插话。
	仅在以下情况回复：
	- 有人直接 @你；
	- 问题明显在问你，或你能自然接话；
	- 对话需要补充有价值的信息；
	- 氛围适合简短玩笑。

	不要回复：
	- 群成员互相聊天，且与你无关；
	- 仅有表情、拍一拍、通知；
	- 你只能给出空泛客套话；
	- 短时间内已经回复过同一话题。

	说话规则：
	- 使用自然、口语化的简体中文，像真实朋友聊天。
	- 默认只回复 1～3 句；能短就短。
	- 不要使用“作为 AI”“我可以帮助你”“当然”“以下是”“总结一下”等客服或助手式表达。
	- 不要主动分点、写标题、长篇解释，除非对方明确要求。
	- 可以偶尔使用“哈哈”“行”“懂了”“确实”等自然口语，但不要每句都用表情。
	- 不确定时坦率说“这个我不太确定”，不要编造。
	- 对方闲聊就闲聊，不要自动给建议或教学。
	- 保持前后人设一致，但不声称自己是真人。
			`
	appLog := logger.New()
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("API_KEY")),
		option.WithBaseURL(os.Getenv("BASE_URL")),
	)
	completion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model: "gpt-5.5",
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(msg.Content),
		},
	})
	if err != nil {
		panic(err)
	}

	appLog.Info(completion.Choices[0].Message.Content)
	return completion.Choices[0].Message.Content
}
