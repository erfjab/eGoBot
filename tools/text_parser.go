package tools

import (
	"html"
	"sort"
	"strconv"
	"strings"

	"github.com/erfjab/egobot/models"
)

type parseMode int

const (
    parseModeHTML parseMode = iota
    parseModeMarkdown
)

type entityRange struct {
    entity models.MessageEntity
    start  int
    end    int
    length int
}

// ParseMessageHTML extracts message text or caption as HTML.
func ParseMessageHTML(message *models.Message) string {
	if message == nil {
		return ""
	}
	if message.Text != "" {
		return ParseTextHTML(message.Text, message.Entities)
	}
	if message.Caption != "" {
		return ParseTextHTML(message.Caption, message.CaptionEntities)
	}
	return ""
}

// ParseMessageMarkdown extracts message text or caption as MarkdownV2.
func ParseMessageMarkdown(message *models.Message) string {
	if message == nil {
		return ""
	}
	if message.Text != "" {
		return ParseTextMarkdown(message.Text, message.Entities)
	}
	if message.Caption != "" {
		return ParseTextMarkdown(message.Caption, message.CaptionEntities)
	}
	return ""
}

// ParseTextHTML parses a text with entities to HTML.
func ParseTextHTML(text string, entities []models.MessageEntity) string {
	return parseTextWithEntities(text, entities, parseModeHTML)
}

// ParseTextMarkdown parses a text with entities to MarkdownV2.
func ParseTextMarkdown(text string, entities []models.MessageEntity) string {
	return parseTextWithEntities(text, entities, parseModeMarkdown)
}

func parseTextWithEntities(text string, entities []models.MessageEntity, mode parseMode) string {
	if text == "" {
		return ""
	}

	runes := []rune(text)
	positions := buildUTF16Positions(runes)
	ranges := buildEntityRanges(entities, positions, len(runes))

	opens := make(map[int][]entityRange)
	closes := make(map[int][]entityRange)
	for _, r := range ranges {
		opens[r.start] = append(opens[r.start], r)
		closes[r.end] = append(closes[r.end], r)
	}

	var builder strings.Builder
	inCode := 0
	inPre := 0

	for i := 0; i <= len(runes); i++ {
		if items := closes[i]; len(items) > 0 {
			sort.Slice(items, func(a, b int) bool { return items[a].length < items[b].length })
			for _, item := range items {
				builder.WriteString(closeTag(item.entity, mode))
				if mode == parseModeMarkdown {
					switch item.entity.Type {
					case "code":
						if inCode > 0 {
							inCode--
						}
					case "pre":
						if inPre > 0 {
							inPre--
						}
					}
				}
			}
		}

		if items := opens[i]; len(items) > 0 {
			sort.Slice(items, func(a, b int) bool { return items[a].length > items[b].length })
			for _, item := range items {
				builder.WriteString(openTag(item.entity, mode))
				if mode == parseModeMarkdown {
					switch item.entity.Type {
					case "code":
						inCode++
					case "pre":
						inPre++
					}
				}
			}
		}

		if i == len(runes) {
			break
		}

		char := string(runes[i])
		if mode == parseModeHTML {
			builder.WriteString(html.EscapeString(char))
			continue
		}

		if inCode > 0 || inPre > 0 {
			builder.WriteString(escapeMarkdownCode(char))
		} else {
			builder.WriteString(escapeMarkdownV2(char))
		}
	}

	return builder.String()
}

func buildEntityRanges(entities []models.MessageEntity, positions []int, runeLen int) []entityRange {
	ranges := make([]entityRange, 0, len(entities))
	for _, entity := range entities {
		start := runeIndexFromUTF16Offset(positions, entity.Offset)
		end := runeIndexFromUTF16Offset(positions, entity.Offset+entity.Length)
		if start < 0 {
			start = 0
		}
		if end < start {
			end = start
		}
		if start > runeLen {
			start = runeLen
		}
		if end > runeLen {
			end = runeLen
		}
		ranges = append(ranges, entityRange{
			entity: entity,
			start:  start,
			end:    end,
			length: end - start,
		})
	}
	return ranges
}

func buildUTF16Positions(runes []rune) []int {
	positions := make([]int, len(runes)+1)
	count := 0
	for i, r := range runes {
		positions[i] = count
		if r > 0xFFFF {
			count += 2
		} else {
			count += 1
		}
	}
	positions[len(runes)] = count
	return positions
}

func runeIndexFromUTF16Offset(positions []int, offset int) int {
	if offset <= 0 {
		return 0
	}
	idx := sort.Search(len(positions), func(i int) bool { return positions[i] >= offset })
	if idx < 0 {
		return 0
	}
	if idx > len(positions)-1 {
		return len(positions) - 1
	}
	return idx
}

func openTag(entity models.MessageEntity, mode parseMode) string {
	switch mode {
	case parseModeHTML:
		switch entity.Type {
		case "bold":
			return "<b>"
		case "italic":
			return "<i>"
		case "underline":
			return "<u>"
		case "strikethrough":
			return "<s>"
		case "spoiler":
			return "<span class=\"tg-spoiler\">"
		case "code":
			return "<code>"
		case "pre":
			if entity.Language != "" {
				return "<pre><code class=\"language-" + html.EscapeString(entity.Language) + "\">"
			}
			return "<pre><code>"
		case "text_link":
			return "<a href=\"" + html.EscapeString(entity.URL) + "\">"
		case "text_mention":
			if entity.User != nil {
				return "<a href=\"tg://user?id=" + strconv.FormatInt(entity.User.ID, 10) + "\">"
			}
			return ""
		case "custom_emoji":
			if entity.CustomEmojiID != "" {
				return "<tg-emoji emoji-id=\"" + html.EscapeString(entity.CustomEmojiID) + "\">"
			}
			return ""
		default:
			return ""
		}
	case parseModeMarkdown:
		switch entity.Type {
		case "bold":
			return "**"
		case "italic":
			return "_"
		case "underline":
			return "__"
		case "strikethrough":
			return "~~"
		case "spoiler":
			return "||"
		case "code":
			return "`"
		case "pre":
			if entity.Language != "" {
				return "```" + entity.Language + "\n"
			}
			return "```\n"
		case "text_link":
			return "["
		case "text_mention":
			return "["
		default:
			return ""
		}
	default:
		return ""
	}
}

func closeTag(entity models.MessageEntity, mode parseMode) string {
	switch mode {
	case parseModeHTML:
		switch entity.Type {
		case "bold":
			return "</b>"
		case "italic":
			return "</i>"
		case "underline":
			return "</u>"
		case "strikethrough":
			return "</s>"
		case "spoiler":
			return "</span>"
		case "code":
			return "</code>"
		case "pre":
			return "</code></pre>"
		case "text_link":
			return "</a>"
		case "text_mention":
			return "</a>"
		case "custom_emoji":
			return "</tg-emoji>"
		default:
			return ""
		}
	case parseModeMarkdown:
		switch entity.Type {
		case "bold":
			return "**"
		case "italic":
			return "_"
		case "underline":
			return "__"
		case "strikethrough":
			return "~~"
		case "spoiler":
			return "||"
		case "code":
			return "`"
		case "pre":
			return "\n```"
		case "text_link":
			return "](" + escapeMarkdownV2(entity.URL) + ")"
		case "text_mention":
			if entity.User != nil {
				return "](tg://user?id=" + strconv.FormatInt(entity.User.ID, 10) + ")"
			}
			return "]()"
		default:
			return ""
		}
	default:
		return ""
	}
}

func escapeMarkdownV2(text string) string {
	var builder strings.Builder
	for _, r := range text {
		switch r {
		case '_', '*', '[', ']', '(', ')', '~', '`', '>', '#', '+', '-', '=', '|', '{', '}', '.', '!':
			builder.WriteRune('\\')
		}
		builder.WriteRune(r)
	}
	return builder.String()
}

func escapeMarkdownCode(text string) string {
	var builder strings.Builder
	for _, r := range text {
		if r == '\\' || r == '`' {
			builder.WriteRune('\\')
		}
		builder.WriteRune(r)
	}
	return builder.String()
}