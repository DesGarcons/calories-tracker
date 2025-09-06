package bot

import (
	tele "gopkg.in/telebot.v4"
)

type Bot struct {
	*tele.Bot
	Buttons *Buttons
	Menus   *Menus
}

type Menus struct {
	AuthMenu *tele.ReplyMarkup
}

type Buttons struct {
	AuthorizationBtn *tele.Btn
}

func New(b *tele.Bot) *Bot {

	// Начальное меню (меню старта)
	authMenu := &tele.ReplyMarkup{}
	authorizationBtn := b.NewMarkup().Data("Начать", "startAuth", "")
	authMenu.InlineKeyboard = append(authMenu.InlineKeyboard, []tele.InlineButton{*authorizationBtn.Inline()})

	// Главное меню
	mainMenu := &tele.ReplyMarkup{ResizeKeyboard: true}
	mainMenu.Row(mainMenu.Text("Записать прием пищи"))

	menu := Menus{
		AuthMenu: authMenu,
	}

	buttons := Buttons{
		AuthorizationBtn: &authorizationBtn,
	}

	return &Bot{
		b,
		&buttons,
		&menu,
	}
}
