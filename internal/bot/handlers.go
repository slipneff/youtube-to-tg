package bot

import (
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/mymmrac/telego"
)

func (b *Bot) HandleUpdates(updates <-chan telego.Update) {
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				b.handleStart(update.Message.Chat.ID)
			default:
				b.handleYoutubeURL(youtubeURL{
					URL:    update.Message.Text,
					ChatID: update.Message.Chat.ID,
				})
			}
		}
	}
}
func (b *Bot) handleStart(chatID int64) {
	b.SendMessage(&telego.SendMessageParams{
		ChatID: telego.ChatID{
			ID: chatID,
		},
		Text: "Hello, world!",
	})
}

type youtubeURL struct {
	URL    string
	ChatID int64
}

func (b *Bot) handleYoutubeURL(input youtubeURL) {
	if !b.ValidateYoutubeURL(input.URL) {
		b.SendMessage(&telego.SendMessageParams{
			ChatID: telego.ChatID{
				ID: input.ChatID,
			},
			Text: "Invalid URL",
		})
		return
	}
	m, _ := b.SendMessage(&telego.SendMessageParams{
		ChatID: telego.ChatID{
			ID: input.ChatID,
		},
		Text: "Downloading...",
	})
	path, err := b.script.Run(input.URL)
	if err != nil {
		b.EditMessageText(&telego.EditMessageTextParams{
			MessageID: m.MessageID,
			ChatID: telego.ChatID{
				ID: input.ChatID,
			},
			Text: "Error: " + err.Error(),
		})
		return
	}
	b.EditMessageText(&telego.EditMessageTextParams{
		MessageID: m.MessageID,
		ChatID: telego.ChatID{
			ID: input.ChatID,
		},
		Text: "Uploading...",
	})
	file, err := os.Open(path)
	if err != nil {
		b.EditMessageText(&telego.EditMessageTextParams{
			MessageID: m.MessageID,
			ChatID: telego.ChatID{
				ID: input.ChatID,
			},
			Text: "Error:" + err.Error(),
		})
		return
	}
	defer file.Close()
	defer os.Remove(path)
	b.EditMessageText(&telego.EditMessageTextParams{
		MessageID: m.MessageID,
		ChatID: telego.ChatID{
			ID: input.ChatID,
		},
		Text: "Done!",
	})
	b.SendAudio(&telego.SendAudioParams{
		ChatID: telego.ChatID{
			ID: input.ChatID,
		},
		Audio: telego.InputFile{
			File: file,
		},
	})

}

func (b *Bot) ValidateYoutubeURL(urlStr string) bool {
	var YouTubeVideoURLPattern = regexp.MustCompile(`^(https?\:\/\/)?(www\.)?(youtube\.com|youtu\.?be)\/.+$`)

	if YouTubeVideoURLPattern.MatchString(urlStr) {
		parsedURL, err := url.Parse(urlStr)
		if err == nil && (strings.Contains(parsedURL.Host, "youtube.com") || strings.Contains(parsedURL.Host, "youtu.be")) {
			return true
		}
	}

	return false
}
