package service

import (
	"fmt"
	"log"
	"strings"
	"tgbot/dao/mysql"

	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// global b function interface
var b *gotgbot.Bot

type tgbotService struct {
}

var TgbotService = new(tgbotService)

// This bot demonstrates some example interactions with telegram callback queries.
// It has a basic start command which contains a button. When pressed, this button is edited via a callback query.
func (t *tgbotService) TgCallbackInit() {
	// Get token from the environment variable
	// token := os.Getenv("token")
	token := viper.GetString("tgbot.token")
	if token == "" {
		panic("TOKEN is empty")
	}
	var err error
	// Create bot from environment value.
	b, err = gotgbot.NewBot(token, nil)
	if err != nil {
		zap.L().Error("failed to create new bot:", zap.Error(err))
		// panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("help", t.callbakhelp))
	dispatcher.AddHandler(handlers.NewCommand(testA_Entrance, t.dev))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("一个大大的测试"), t.testMesaage))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal(testA_background), t.triggerBuild))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal(testA_h_wcapise), t.triggerBuild))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal(testB_IMG_Before), t.triggerBuild))

	// Set Menu to SetMyCommands
	t.setBotCommands()

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}

// bot menu function
func (t *tgbotService) setBotCommands() {
	commands := []gotgbot.BotCommand{
		// command : interface name,description: mark
		{Command: testA_Entrance, Description: fmt.Sprintf("%s the bot 测试", testA_Entrance)},
		{Command: "help", Description: "Get help"},
	}
	_, err := b.SetMyCommands(commands, nil)
	if err != nil {
		log.Fatalf("failed to set bot commands: %v", err)
	}
	log.Println("Bot commands set successfully.")
}

// /dev env type arys1 arys1
// Cut Input Data text, Get slice
func (t *tgbotService) stringSplit(text string) (firsttext, twotext, manytext string, total int) {
	parts := strings.Split(text, " ")
	if len(parts) == 1 {
		return parts[0], "", "", len(parts)
	}
	if len(parts) == 2 {
		return parts[1], "", "", len(parts)
	}
	if len(parts) == 3 {
		return parts[1], parts[2], "", len(parts)
	}
	return parts[1], parts[2], strings.Join(parts[3:], " "), len(parts)
}

// This is the function that handles help,used to prompt usage
// text fromat see https://sendpulse.com/knowledge-base/chatbot/telegram/format-text
func (t *tgbotService) callbakhelp(b *gotgbot.Bot, ctx *ext.Context) error {
	texts := `Hello, I'm @%s. 
Please make sure the input format is correct. By default, the latest code will be pulled, and no version selection is provided.

<b>DT update example is as follows:</b>
Test environment: <code>/%s Grayscale vue 101 102</code>
Production environment: <code>/%s Online vue 101 102</code>

<b>Code_Type:</b> %s
<b>Environment:</b> %s
<ins>Placeholder 0 is Handler, 1 is platform, 2 is type, and 3 is webid. Separated by space</ins>

Update without parameters: <code> /%s </code>
`

	err := t.SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf(texts, b.User.Username, DT_Entrance, DT_Entrance, DT_CODE_TYPE, Platform, DT_Entrance))
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// Receive /dev
// Receive /dev input parameters and process them
// handlers.NewCommand <-
func (t *tgbotService) dev(b *gotgbot.Bot, ctx *ext.Context) error {

	zap.L().Info(fmt.Sprintf("bot:%s, 原始输入:%s", b.User.Username, ctx.Message.Text))

	// cut input message text
	firsttext, twotext, manytext, total := t.stringSplit(ctx.Message.Text)
	fmt.Println("", firsttext, twotext, manytext)

	// first text
	// if cut text total is eqel one，open button
	// 判断输入的参数，如果只是接口名称，则调用button按钮的逻辑
	if total == 1 {
		_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, I'm @%s. 其他选项通过/help查看帮助.", b.User.Username), &gotgbot.SendMessageOpts{
			ParseMode: "html",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: testA_background, CallbackData: testA_background},
				}, {
					{Text: testB_IMG_Before, CallbackData: testB_IMG_Before},
				}, {
					{Text: testA_h_wcapise, CallbackData: testA_h_wcapise},
				},
				},
			},
		})
		if err != nil {
			return fmt.Errorf("failed to send start message: %w", err)
		}
		return err
	}

	// 如果小于4个参数位置则会弹出帮助信息
	if total < 4 {
		return t.callbakhelp(b, ctx)
	}
	// check if the user exists in db list
	// 检查固定的项目中的多个参数构建，项目对应jenkins中的项目名称
	// 用户发起接口调用后传递的参数将被进行切割，将在Test的db表中查询发起的用户是否存在
	// res, err := checkUser(Test, ctx.EffectiveUser.FirstName, Dev, ctx)
	res, err := t.checkUser(Test, ctx.EffectiveUser.FirstName, firsttext, ctx)
	if err != nil || !res {
		return err
	}
	// 如果满足,检查参数是否为内置的参数
	if !t.PlatformCheck(Platform, firsttext) {
		err = t.SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintln(ErrorOne))
		if err != nil {
			zap.L().Fatal(fmt.Sprintf("%s:,args: %s,err: %s\n", ErrorOne, ctx.Message.Text, err))
			return err
		}
		return err
	}
	// Check if the passed parameter exists in the existing list
	// 如果满足,检查参数是否为内置的参数
	if !t.CodeTypeCheck(testA_CODE_TYPE, twotext) {
		err = t.SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintln(ErrorOne))
		if err != nil {
			zap.L().Fatal(fmt.Sprintf("%s:,args: %s,err: %s\n", ErrorOne, ctx.Message.Text, err))
			return err
		}
		return err
	}

	// sed message to Groups
	// 发起一个构建前通知
	err = t.SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf("Hello, I'm @%s. 我为%s环境工作.\n即将发起对%s的更新,web id包括: %s", b.User.Username, firsttext, twotext, manytext))
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	// connent jenkins build
	// 向jenkins发起构建
	// result, buildnum, buildtime, err := maintrigger(Test, firsttext, twotext, manytext)
	buildtime, result, buildnum := 1.00, "不调用测试运行", 15
	if err != nil {
		err := t.SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf("更新失败:%s\n确保你的格式正确!,通过/help查看帮助.", twotext))
		if err != nil {
			return fmt.Errorf("failed to send start message: %w", err)
		}
		return fmt.Errorf("->: %s", err)
	}
	// sed build result info to Groups
	// 返回构建的结果
	// err = SendMessage(ctx, viper.GetInt64("tggroup.chatid1"), fmt.Sprintf("%s环境的%s,%s.更新完成\n状态:%s, id:%d, %s", firsttext, twotext, manytext, result, buildnum, buildtime))
	err = t.SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf("更新完成,%f,状态: %s, id: %d", buildtime, result, buildnum))
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// 发送消息到tg
// chatid是群或者聊天对话的id，通过以下方式获取
// get user id @userinfobot /start
// get group id  https://api.telegram.org/bot[token,不包括括号]/getUpdates
func (t *tgbotService) SendMessage(ctx *ext.Context, chatId int64, text string) error {
	_, err := ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
func (t *tgbotService) SendMessageAlone(ctx *ext.Context, chatId int64, text string) error {
	_, err := b.SendMessage(chatId, text, &gotgbot.SendMessageOpts{})
	if err != nil {
		zap.L().Fatal("failed to send start message: %s\n", zap.Error(err))
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

// and test
func (t *tgbotService) testMesaage(b *gotgbot.Bot, ctx *ext.Context) error {
	fmt.Println("发起机器人:", b.User.Username)
	fmt.Println("发起回调用户:", ctx.CallbackQuery.From.FirstName)
	fmt.Println("data:", ctx.CallbackQuery.Data)
	cb := ctx.Update.CallbackQuery

	if ctx.CallbackQuery.From.FirstName == "Mark" {
		fmt.Println("test......")

		_, err := cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text: "You pressed a button!",
		})
		if err != nil {
			return fmt.Errorf("failed to answer start callback query: %w", err)
		}
		_, _, err = cb.Message.EditText(b, fmt.Sprintf("[%s]触发了[ %s ]的更新.", ctx.CallbackQuery.From.FirstName, ctx.CallbackQuery.Data), nil)
		if err != nil {
			return fmt.Errorf("failed to edit start message text: %w", err)
		}
	} else {
		_, _, err := cb.Message.EditText(b, fmt.Sprintf("[%s]没有权限使用[ %s ]更新.", ctx.CallbackQuery.From.FirstName, ctx.CallbackQuery.Data), nil)
		if err != nil {
			return fmt.Errorf("failed to edit start message text: %w", err)
		}
	}
	return nil
}

func (t *tgbotService) checkUser(project, runUser, runEnv string, ctx *ext.Context) (bool, error) {
	err, result := mysql.Calldb.SelectProject(project, runUser, runEnv)
	if err != nil {
		zap.L().Fatal("mysql.SelectProject: %s\n", zap.Error(err))
		return false, err
	}
	if !result {
		err = t.SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintln(ErrorUser))
		if err != nil {
			zap.L().Fatal("failed to send start message:: %s\n", zap.Error(err))
			return result, err
		}
		return result, err
	}
	return result, nil
}

func (t *tgbotService) CodeTypeCheck(codetype, args string) bool {
	types := strings.Split(codetype, "|")
	for _, t := range types {
		if t == args {
			return true
		}
	}
	return false
}
func (t *tgbotService) PlatformCheck(platform, args string) bool {
	types := strings.Split(platform, "|")
	for _, t := range types {
		if t == args {
			return true
		}
	}
	return false
}

// 处理点击Buttion的事件.
func (t *tgbotService) triggerBuild(b *gotgbot.Bot, ctx *ext.Context) error {
	fmt.Println("data:", ctx.CallbackQuery.Data, b.FirstName, ctx.EffectiveUser.FirstName)
	// check button callback user auth
	// 检查点击从按钮中传递的参数进行db查询所在项目中的发起用户是否存在
	res, err := t.checkUser(ctx.CallbackQuery.Data, ctx.EffectiveUser.FirstName, Online, ctx)
	if err != nil || !res {
		return err
	}

	// sed message to Groups
	err = t.SendMessageAlone(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf("Hello, I'm @%s.\n即将发起对%s的更新", b.User.Username, ctx.CallbackQuery.Data))
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	// connent jenkins build
	// 发起jenkins构建
	// result, buildnum, buildtime, err := JenkinsService.maintrigger(ctx.CallbackQuery.Data, "", "", "")
	result, buildnum, buildtime := "逻辑测试ok", 10, "test"
	if err != nil {
		err := t.SendMessageAlone(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf("更新失败:%s\n确保你的格式正确,通过/help查看帮助.", ctx.CallbackQuery.Data))
		if err != nil {
			return fmt.Errorf("failed to send start message: %w", err)
		}
		return fmt.Errorf("->: %s", err)
	}
	// sed build result info to Groups
	// err = SendMessage(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf("%s环境的%s,%s.更新完成\n状态:%s, id:%d, %s", firsttext, twotext, manytext, result, buildnum, buildtime))
	err = t.SendMessageAlone(ctx, viper.GetInt64(TG_GROUP_CHATID), fmt.Sprintf("更新完成,%s,状态: %s, id: %d", buildtime, result, buildnum))
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
