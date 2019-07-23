package main

import (
	"log"
	"regexp"
	"strings"
)

var Regexp = []string{
	"^(なし|無し|無)[。|.]?[\n|\t|\r]?$",
	"^ありません[。|.]?[\n|\t|\r]?$",
	"^(特に|とくに)(なし|無し)[。|.]?[\n|\t|\r]?$",
}

var Strs = []string{
	"なし",
	"無し",
	"無",
	"無。",
	"無。\n",
	"なし\n",
	"なし。",
	"なし。\n",
	"特になし",
	"特になし\n",
	"特になし。",
	"特になし。\n",
	"とくになし",
	"とくになし\n",
	"とくになし。",
	"とくになし。\n",
	"特に無し",
	"特に無し\n",
	"特に無し。",
	"特に無し。\n",
	"とくに無し",
	"とくに無し\n",
	"とくに無し。",
	"とくに無し。\n",
	"ありません",
	"ありません\n",
	"ありません。",
	"ありません。\n",
}

var Dangers = []string{
	"[ア|ｱ|あ|ァ|ｧ][レ|ﾚ|れ][ル|ﾙ|る](ギ|ｷﾞ|ぎ)(ー|-|ｰ)?",
	"(食べ|たべ|くい|食い)(切|き|ら)れ(ない|ません)",
	"(苦手|にがて|嫌い|きらい)",
}

var CheckStrs = []string{
	"カニがアレルギーです",
	"カニがアレルギです",
	"カニがアﾚルｷﾞです",
	"カニがｱﾚﾙｷﾞｰです",
	"カニがあれるぎーです",
	"カニがあれるぎです",
	"カニが食べきれない",
	"カニが食べられない",
	"カニが食べ切れない",
	"カニがたべきれません",
	"カニがたべられません",
	"カニがたべ切れません",
	"カニがくいきれません",
	"カニが食いきれません",
	"カニが食い切れません",
	"カニが嫌いなので",
	"カニが苦手なので",
	"カニがきらいなので",
	"カニがにがてなので",
}

func main() {
	var mcnt int
	for _, s := range Strs {
		for _, r := range Regexp {
			reg := regexp.MustCompile(r)
			m := reg.FindAllStringSubmatch(s, -1)
			if len(m) > 0 {
				mcnt++
				log.Printf("%s\n", strings.Trim(s, "\n"))
			}
		}
	}
	log.Printf("Match Count: %d / %d\n\n", mcnt, len(Strs))

	var mcnt2 int
	for _, s := range CheckStrs {
		for _, r := range Dangers {
			reg := regexp.MustCompile(r)
			m := reg.FindAllStringSubmatch(s, -1)
			if len(m) > 0 {
				mcnt2++
				log.Printf("%s\n", m)
				log.Printf("%s\n", strings.Trim(s, "\n"))
			}
		}
	}
	log.Printf("Match Count: %d / %d\n", mcnt2, len(CheckStrs))
}
