package client

import "github.com/dontpanic92/wxGo/wx"

var ui_username wx.TextCtrl
var ui_password wx.TextCtrl
var ui_login wx.Button
var ui_resume wx.Button
var ui_bar wx.Gauge
var ui_percent wx.StaticText

func setupFeeds(f *TheFrame) {
	f.sizer.Clear(true)
}

func setupError(f *TheFrame, text string) {
	row3 := wx.NewBoxSizer(wx.VERTICAL)
	msg := wx.NewStaticText(f.frame, wx.ID_ANY, text, wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(msg, 0, wx.ALL|wx.EXPAND, 5)
	msg2 := wx.NewStaticText(f.frame, wx.ID_ANY, "Please quit app and try again.", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(msg2, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row3, 0, wx.ALL|wx.EXPAND, 5)
	f.frame.Layout()
}

func setupLogin(f *TheFrame) {
	row := wx.NewBoxSizer(wx.HORIZONTAL)
	msg := wx.NewStaticText(f.frame, wx.ID_ANY, "username", wx.DefaultPosition, wx.DefaultSize, 0)
	row.Add(msg, 0, wx.ALL|wx.EXPAND, 5)
	ui_username = wx.NewTextCtrl(f.frame, wx.ID_ANY, "", wx.DefaultPosition, wx.NewSize(80, 25), 0)
	row.Add(ui_username, 0, wx.ALL|wx.EXPAND, 5)

	row2 := wx.NewBoxSizer(wx.HORIZONTAL)
	msg2 := wx.NewStaticText(f.frame, wx.ID_ANY, "password", wx.DefaultPosition, wx.DefaultSize, 0)
	row2.Add(msg2, 0, wx.ALL|wx.EXPAND, 5)
	ui_password = wx.NewTextCtrl(f.frame, wx.ID_ANY, "", wx.DefaultPosition, wx.NewSize(90, 25), wx.TE_PASSWORD)
	row2.Add(ui_password, 0, wx.ALL|wx.FIXED_MINSIZE, 5)

	row3 := wx.NewBoxSizer(wx.HORIZONTAL)
	ui_login = wx.NewButton(f.frame, wx.ID_ANY, "Login", wx.DefaultPosition, wx.DefaultSize, 0)
	row3.Add(ui_login, 0, wx.ALL|wx.FIXED_MINSIZE, 5)

	f.sizer.Add(row, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row2, 0, wx.ALL|wx.EXPAND, 5)
	f.sizer.Add(row3, 0, wx.ALL|wx.EXPAND, 5)

	wx.Bind(f.frame, wx.EVT_BUTTON, f.evtLogin, ui_login.GetId())
	f.frame.Layout()
}
