package tgbotapi_test

import (
	"github.com/zhulik/telegram-bot-api"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	botToken := os.Getenv("TELEGRAM_API_TOKEN")

	if botToken == "" {
		log.Panic("You must provide a TELEGRAM_API_TOKEN env variable to test!")
	}

	os.Exit(m.Run())
}

func TestNewBotAPI_notoken(t *testing.T) {
	_, err := tgbotapi.NewBotAPI("")

	if err == nil {
		t.Fail()
	}
}

func TestNewBotAPI_token(t *testing.T) {
	_, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}
}

func TestGetUpdates(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	u := tgbotapi.NewUpdate(0)

	_, err = bot.GetUpdates(u)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithMessage(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewMessage(76918703, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithMessageReply(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewMessage(76918703, "A test message from the test library in telegram-bot-api")
	msg.ReplyToMessageID = 480
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithMessageForward(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewForward(76918703, 76918703, 480)
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewPhoto(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewPhotoUpload(76918703, "tests/image.jpg")
	msg.Caption = "Test"
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewPhotoReply(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewPhotoUpload(76918703, "tests/image.jpg")
	msg.ReplyToMessageID = 480

	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingPhoto(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewPhotoShare(76918703, "AgADAgADxKcxG4cBswqt13DnHOgbmBxDhCoABC0h01_AL4SKe20BAAEC")
	msg.Caption = "Test"
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewDocument(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewDocumentUpload(76918703, "tests/image.jpg")
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingDocument(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewDocumentShare(76918703, "BQADAgADBwADhwGzCjWgiUU4T8VNAg")
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewAudio(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewAudioUpload(76918703, "tests/audio.mp3")
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingAudio(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewAudioShare(76918703, "BQADAgADMwADhwGzCkYFlCTpxiP6Ag")
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"

	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewVoice(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewVoiceUpload(76918703, "tests/voice.ogg")
	msg.Duration = 10
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingVoice(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewVoiceShare(76918703, "AwADAgADIgADhwGzCigyMW_GUtWIAg")
	msg.Duration = 10
	_, err = bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithLocation(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	_, err = bot.Send(tgbotapi.NewLocation(76918703, 40, 40))

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewVideo(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewVideoUpload(76918703, "tests/video.mp4")
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err = bot.Send(msg)

	if err != nil {

		t.Fail()
	}
}

func TestSendWithExistingVideo(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewVideoShare(76918703, "BAADAgADRgADhwGzCopBeKJ7rv9SAg")
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err = bot.Send(msg)

	if err != nil {

		t.Fail()
	}
}

func TestSendWithNewSticker(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewStickerUpload(76918703, "tests/image.jpg")

	_, err = bot.Send(msg)

	if err != nil {

		t.Fail()
	}
}

func TestSendWithExistingSticker(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewStickerShare(76918703, "BQADAgADUwADhwGzCmwtOypNFlrRAg")

	_, err = bot.Send(msg)

	if err != nil {

		t.Fail()
	}
}

func TestSendWithNewStickerAndKeyboardHide(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewStickerUpload(76918703, "tests/image.jpg")
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardHide{true, false}
	_, err = bot.Send(msg)

	if err != nil {

		t.Fail()
	}
}

func TestSendWithExistingStickerAndKeyboardHide(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	msg := tgbotapi.NewStickerShare(76918703, "BQADAgADUwADhwGzCmwtOypNFlrRAg")
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardHide{true, false}

	_, err = bot.Send(msg)

	if err != nil {

		t.Fail()
	}
}

func TestGetFile(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	file := tgbotapi.FileConfig{"BQADAgADBwADhwGzCjWgiUU4T8VNAg"}

	_, err = bot.GetFile(file)

	if err != nil {
		t.Fail()
	}
}

func TestSendChatConfig(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	err = bot.SendChatAction(tgbotapi.NewChatAction(76918703, tgbotapi.ChatTyping))

	if err != nil {
		t.Fail()
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	_, err = bot.GetUserProfilePhotos(tgbotapi.NewUserProfilePhotos(76918703))
	if err != nil {
		t.Fail()
	}
}

func TestUpdatesChan(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))

	if err != nil {
		t.Fail()
	}

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	err = bot.UpdatesChan(ucfg)

	if err != nil {
		t.Fail()
	}
}

func ExampleNewBotAPI() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	err = bot.UpdatesChan(u)

	for update := range bot.Updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
