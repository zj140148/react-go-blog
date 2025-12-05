package service

import (
	"backend/internal/about/models"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AboutService 关于页面服务接口
type AboutService interface {
	GetAboutContent() (*models.AboutInfo, error)
}

// aboutService 关于页面服务实现
type aboutService struct {
	contentPath string
}

// NewAboutService 创建关于页面服务实例
func NewAboutService() AboutService {
	return &aboutService{
		contentPath: filepath.Join("static", "about", "about1.md"),
	}
}

// GetAboutContent 获取关于页面内容
func (s *aboutService) GetAboutContent() (*models.AboutInfo, error) {
	// 检查文件是否存在
	if _, err := os.Stat(s.contentPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("关于页面文件不存在: %s", s.contentPath)
	}

	// 读取文件内容
	content, err := os.ReadFile(s.contentPath)
	if err != nil {
		return nil, fmt.Errorf("读取关于页面文件失败: %w", err)
	}

	// 清理和处理内容
	cleanContent := s.cleanContent(string(content))

	// 提取作者信息（从内容中解析）
	author := s.extractAuthor(cleanContent)

	return &models.AboutInfo{
		Content: cleanContent,
		Author:  author,
		Version: "1.0.0",
	}, nil
}

// cleanContent 清理内容
func (s *aboutService) cleanContent(content string) string {
	// 移除多余的空白行
	lines := strings.Split(content, "\n")
	var cleanLines []string
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			cleanLines = append(cleanLines, trimmed)
		}
	}
	
	return strings.Join(cleanLines, "\n")
}

// extractAuthor 从内容中提取作者信息
func (s *aboutService) extractAuthor(content string) string {
	// 简单的作者提取逻辑，寻找"我是"后面的名字
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.Contains(line, "我是") {
			// 提取 **Wh1stle** 这样的格式
			start := strings.Index(line, "**")
			if start != -1 {
				start += 2
				end := strings.Index(line[start:], "**")
				if end != -1 {
					return line[start : start+end]
				}
			}
		}
	}
	
	return "Wh1stle" // 默认作者
}