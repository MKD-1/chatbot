package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	Ai         AiConfig         `json:"ai"`
	PromptPath PromptPathConfig `json:"promptPath"`
}

type AiConfig struct {
	ApiKey         string `json:"apiKey"`
	BaseUrl        string `json:"baseUrl"`
	Model          string `json:"model"`
	TimeoutSeconds int    `json:"timeoutSeconds"`
}
type PromptPathConfig struct {
	SystemPrompt string `json:"systemPrompt"`
	TgPrompt     string `json:"tgPrompt"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("读取配置文件失败: %w", err)
	}
	var cfg Config

	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("解析 JSON 配置失败: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg Config) Validate() error {
	if err := cfg.aiValidate(); err != nil {
		return err
	}
	if err := cfg.promptPathValidate(); err != nil {
		return err
	}
	return nil
}
func (cfg Config) aiValidate() error {
	if strings.TrimSpace(cfg.Ai.ApiKey) == "" {
		return errors.New("ai.apiKey 不能为空")
	}
	baseUrl, err := url.ParseRequestURI(cfg.Ai.BaseUrl)
	if err != nil || baseUrl.Scheme == "" || baseUrl.Host == "" {
		return fmt.Errorf("ai.baseUrl 无效: %q", cfg.Ai.BaseUrl)
	}
	if strings.TrimSpace(cfg.Ai.Model) == "" {
		return errors.New("ai.model 不能为空")
	}
	if cfg.Ai.TimeoutSeconds <= 0 {
		return errors.New("ai.timeoutSeconds 必须大于0")
	}
	return nil
}
func (cfg Config) promptPathValidate() error {
	return nil
}
