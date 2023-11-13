package action

import (
	"github.com/go-lark/lark"
	"github.com/leekchan/accounting"
)

const (
	appOrderID     = "cli_a59c2696a738902f"
	appOrderSecret = "xXhhgKuLJylw8oHJhgUiRg2cnMos26id"
	groupSaleId    = "oc_49c4a333affb4f6cd9b5197b02f25329"
)

func SendMessage(msg string) {

	bot := lark.NewChatBot(appOrderID, appOrderSecret)
	bot.GetTenantAccessTokenInternal(true)

	ms := lark.NewMsgBuffer(lark.MsgPost)
	postContent := lark.NewPostBuilder().
		Title("post title").
		TextTag("hello, world", 1, true).
		LinkTag("Google", "https://google.com/").
		AtTag("www", "oc_xxxxxx").
		ImageTag("d9f7d37e-c47c-411b-8ec6-9861132e6986", 300, 300).
		Render()
	om := ms.BindOpenChatID(groupSaleId).Post(postContent).Build()

	bot.PostMessage(om)

}

func NotifySuccessOrder(user_id, customer string, totalMoney int64) {

	bot := lark.NewChatBot(appOrderID, appOrderSecret)
	bot.GetTenantAccessTokenInternal(true)

	ac := accounting.Accounting{Symbol: "", Precision: 0}

	ms := lark.NewMsgBuffer(lark.MsgText)
	tb := lark.NewTextBuilder()

	if totalMoney >= 100000 && totalMoney <= 999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(0.01*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 1000000 && totalMoney <= 1999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(0.5*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 2000000 && totalMoney <= 2999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(0.65*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 3000000 && totalMoney <= 3999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(0.75*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 4000000 && totalMoney <= 4999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(1*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 5000000 && totalMoney <= 5999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(1.25*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 6000000 && totalMoney <= 6999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(1.5*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 7000000 && totalMoney <= 7999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(1.75*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 8000000 && totalMoney <= 9999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(1.9*float64(totalMoney)/100)).
			Textln("")
	} else if totalMoney >= 10000000 && totalMoney <= 99999999 {
		tb = tb.
			Mention(user_id).
			Textf("<b> đã kiếm được %vVNĐ tiền hoa hồng 🎉🎉</b>", ac.FormatMoney(2*float64(totalMoney)/100)).
			Textln("")
	}

	tb = tb.Text("Chúc mừng ").
		Mention(user_id).
		Textf(" đã bán thành công đơn hàng trị giá <b>%vVNĐ</b> 🥳🥳\n", ac.FormatMoney(totalMoney)).
		Textln("").
		Textf("Hãy cố gắng giữ phong độ ").
		Mention(user_id).
		Textf(" nhé 💪💪 và đừng quên theo đuổi và chăm sóc khách VIP <b>%v</b> nha!\n", customer)

	om := ms.BindChatID(groupSaleId).Text(tb.Render()).Build()

	bot.PostMessage(om)

}
