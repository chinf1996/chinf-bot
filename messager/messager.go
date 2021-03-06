package messager

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func init() {
	log.Println("Messager is camming...")
}

//PushMessage 詢問是否要加入會員
func PushMessage(UserID string, bot *linebot.Client) {
	text := fmt.Sprintf("[%v][yes]", UserID)
	yesBtn := linebot.NewPostbackAction("我願意", text, "", "")
	butTemplate := linebot.NewButtonsTemplate(
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRxbYZs9-LkRScXKWthdxw8gwUDUBkG34q0DgZnkI1pOkfybDx-",
		"來註冊成為我們的會員吧！",
		"有跟我一起買Spotify的夥伴就來成為會員吧",
		yesBtn)
	message := linebot.NewTemplateMessage("Sorry :(, please update your app.", butTemplate)
	_, err := bot.PushMessage(UserID, message).Do()
	if err != nil {
		print(err)
	}
}

//PushMessageSay 可以藉由此講出你想講的話
func PushMessageSay(UserID string, bot *linebot.Client, say string) {
	message := linebot.NewTextMessage(say)
	_, err := bot.PushMessage(UserID, message).Do()
	if err != nil {
		print(err)
	}
}

//JoinMember 處理已經有加好友，但還沒加會員的使用者
func JoinMember(db *sql.DB, bot *linebot.Client) {
	query := `select line_id from spotify_user where name is null;`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var userID string
		if err := rows.Scan(&userID); err != nil {
			log.Println(err)
		}
		PushMessage(userID, bot)
	}
}
