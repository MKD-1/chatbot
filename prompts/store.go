package prompts

import (
	"fmt"
	"os"
	"strings"
)

type Store struct {
	prompts map[string]string
}

func Load(promptPaths map[string]string) (*Store, error) {
	prompts := make(map[string]string, len(promptPaths))
	for name, path := range promptPaths {
		name = strings.TrimSpace(name)
		path = strings.TrimSpace(path)

		if name == "" {
			return nil, fmt.Errorf("提示词名称为空")
		}
		if path == "" {
			return nil, fmt.Errorf("提示词 %q 的路径为空", name)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("读取提示词 %q 失败 (%s) : %w", name, path, err)
		}
		prompt := strings.TrimSpace(string(data))
		if prompt == "" {
			return nil, fmt.Errorf("提示词 %q 内容为空 (%s) ", name, path)
		}
		prompts[name] = prompt
	}
	return &Store{prompts: prompts}, nil
}

func (s *Store) Get(name string) (string, error) {
	prompt, exist := s.prompts[name]
	if exist == false {
		return "", fmt.Errorf("未配置提示词 %q", name)
	}
	return prompt, nil
}
