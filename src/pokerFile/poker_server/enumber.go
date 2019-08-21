package poker_server

const (
	royal    = iota //皇家同花顺
	flush           //同花顺
	four            //4条
	threeTwo        //3+2
	suit            //同花
	sequence        //顺子
	three           //3张
	couple2         //两对
	couple          //一对
	alone           //散
)

const (
	spades = iota
	hearts
	diamonds
	clubs
)
