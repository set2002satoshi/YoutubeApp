package main

import (
	



	"github.com/set2002satoshi/YoutubeApp/backend/routes"
	"github.com/set2002satoshi/YoutubeApp/backend/db"
	
)
func main() {


	// マイグレート実行
	db.DbInit()
	// route
	routes.SetUpRoute()


	// テスト確認用
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	_, _ = fmt.Fprint(w, "Hello, 世界")
	// })

	// http.ListenAndServe(":8080", nil)

}


