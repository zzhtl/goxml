package gooxml_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zzhtl/goxml/document"
)

// DocParagraph -
type DocParagraph struct {
	Title    string          `json:"title"`
	Level    string          `json:"level"`
	Children []*DocParagraph `json:"children"`
	Index    int64           `json:"index"`
}

func Test_readParagraphs(t *testing.T) {
	docx, err := document.Open("testdata/office/template.docx")
	if err != nil {
		panic(err)
	}
	var docxParagraphs []*DocParagraph
	var index int64
	for _, p := range docx.Paragraphs() {
		level := getHeadingLevel(p.Style())
		if level != "" {
			index++
			title := getParagraphText(p)
			newParagraph := &DocParagraph{Title: title, Level: level, Index: index}
			fmt.Println("level:", level, "title:", title)
			if len(docxParagraphs) > 0 {
				// 查找合适的父级
				parent := findParent(docxParagraphs, level)
				if parent != nil {
					parent.Children = append(parent.Children, newParagraph)
				} else {
					docxParagraphs = append(docxParagraphs, newParagraph)
				}
			} else {
				docxParagraphs = append(docxParagraphs, newParagraph)
			}
		}
	}
}

func getHeadingLevel(style string) string {
	switch style {
	case "1", "2", "3", "4", "5", "6", "7", "8":
		return style
	case "a0", "afe", "aff0":
		return style
	default:
		return ""
	}
}

func getParagraphText(p document.Paragraph) string {
	var title string
	for _, run := range p.Runs() {
		if isNotBlank(run.Text()) || title != "" {
			title += run.Text()
		}
	}
	return title
}

func isNotBlank(s string) bool {
	trimmed := strings.TrimSpace(s)
	return len(trimmed) > 0
}

func findParent(paragraphs []*DocParagraph, level string) *DocParagraph {
	// 从最后一个标题开始，向上查找合适的父级标题
	for i := len(paragraphs) - 1; i >= 0; i-- {
		if paragraphs[i].Level < level {
			// 找到合适的父级标题
			if len(paragraphs[i].Children) > 0 {
				// 如果父级有子标题，递归查找
				parent := findParent(paragraphs[i].Children, level)
				if parent != nil {
					return parent
				}
			}
			// 如果没有找到递归的父级，当前标题就是父级
			return paragraphs[i]
		}
	}
	return nil
}
