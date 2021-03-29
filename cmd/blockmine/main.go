package main

import (
	"github.com/okokebot/blockmine/internal/server"
)

func main() {
	// ss := "[#64751](https://estoc.weseek.co.jp/redmine/issues/64751)) [SISng][frontend-design]UI SpecをもとにSI詳細画面の見た目を変更する その2 ([#65084](https://estoc.weseek.co.jp/redmine/issues/65084)) [nautilus]ネストされたhashをjbuilderでキャメライズする際、Jbuilder.deep_format_keys を true にして自動でキャメライズされるようにする ([#64458](https://estoc.weseek.co.jp/redmine/issues/64458)) [My.JPNAP][nautilus][ML] 別のCustomerで登録済みのメールアドレスを自Customerで登録できない ([#64748](https://estoc.weseek.co.jp/redmine/issues/64748)) [SISng] Autnum 詳細画面の Customer ASN Ownership に表示される Customer 名を to_s の値に統一する ([#64787](https://estoc.weseek.co.jp/redmine/issues/64787)) [nautilus] N+1警告の問題を改善する sprint-105 ([#65169](https://estoc.weseek.co.jp/redmine/issues/65169)) "
	// ids := importid.PulloutIssueIdFromReleaseNote(ss)
	// fmt.Println(ids)
	// c := redmine.NewClient()
	// for _, id := range ids {
	// 	p := c.GetIssue(id)
	// 	c.GetChildrenInfo(p)
	// 	s := p.CreateReleaseBlock(*c)
	// 	fmt.Println(s)
	// }
	server.Server()
}
