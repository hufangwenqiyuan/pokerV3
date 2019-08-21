package poker_server

import "strings"

func createFullTableWithoutGhost() map[string]seqCards {
	ret := make(map[string]seqCards)
	smallTab := createSequenceTableWithoutGhost()

	for k, v := range smallTab {
		max := table[k[0]]
		if k == "A5432" {
			max = 5
		}
		createCardsTable(5, ret, divideStr(k), "", []int{max}, v)
	}
	return ret
}

func divideStr(s string) (ret []string) {
	for i := 0; i < len(s); i++ {
		ret = append(ret, string(s[i]))
	}
	return
}
func createCardsTable(num int, tab map[string]seqCards, all []string, part string, max []int, mode int) {
	if len(part) == num {
		sc := seqCards{mode, max}
		tab[part] = sc
		return
	}
	ss := divideStr(part)
	for i := 0; i < len(all); i++ {
		temp := part
		if len(ss) == 0 || !stringContainedInSlice(all[i], ss) {
			temp += all[i]
			createCardsTable(num, tab, all, temp, max, mode)
		}
	}
}

func stringContainedInSlice(str string, sli []string) bool {
	for _, s := range sli {
		if str == s {
			return true
		}
	}
	return false
}

func createSequenceTableWithoutGhost() map[string]int {
	table := make(map[string]int)
	str := []string{"K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	table["AKQJT"] = royal
	table["A5432"] = flush
	for i := 0; i < 8; i++ {
		key := str[i] + str[i+1] + str[i+2] + str[i+3] + str[i+4]
		table[key] = flush
	}
	//fmt.Println(table)
	return table
}

func create32Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(2, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]), string(item[0]), string(item[0]), string(item[1]), string(item[1])}
		kind := seqCards{threeTwo, []int{table[item[0]], table[item[1]]}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p] = kind
		}
	}
}
func isSwap(str []string, begin, end int) bool {
	for i := begin; i < end; i++ {
		if str[i] == str[end] {
			return false
		}
	}
	return true
}

func permutate(sample []string, begin, end int, kinds *[]string) {
	if begin == end {
		*kinds = append(*kinds, strings.Join(sample, ""))
		return
	}

	for i := begin; i < end; i++ {
		if isSwap(sample, begin, i) {
			sample[begin], sample[i] = sample[i], sample[begin]
			permutate(sample, begin+1, end, kinds)
			sample[begin], sample[i] = sample[i], sample[begin]
		}
	}
}

func selectCards(num int, part string, ret *[]string) {
	if len(part) == num {
		*ret = append(*ret, part)
		return
	}
	all := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	ss := divideStr(part)
	for i := 0; i < len(all); i++ {
		temp := part
		if len(ss) == 0 || !stringContainedInSlice(all[i], ss) {
			temp += all[i]
			//fmt.Println(temp)
			selectCards(num, temp, ret)
		}
	}
}

func NewCardBuf() *cardBuf {
	cb := &cardBuf{cardState: alone, same: true}
	cb.sOrder = make([]int, 7)
	cb.dOrder = make([]int, 3)
	cb.tOrder = make([]int, 2)
	cb.color = make([]cardColor, 4)
	cb.color[spades].cards = make([]int, 7)
	cb.color[clubs].cards = make([]int, 7)
	cb.color[hearts].cards = make([]int, 7)
	cb.color[diamonds].cards = make([]int, 7)

	cb.tabGhost = createSequenceTableWithGhost()
	cb.tabNoGhost = createSequenceTableWithoutGhost()
	cb.tabFullNoGhost = createFullTableWithoutGhost()

	return cb
}

func createSequenceTableWithGhost() map[string]int {
	table := make(map[string]int)
	str := []string{"Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	table["AKQJ"] = royal
	table["AQJT"] = royal
	table["AKQT"] = royal
	table["AKJT"] = royal
	table["KQJT"] = royal

	table["KJT9"] = flush
	table["KQT9"] = flush
	table["KQJ9"] = flush
	table["QJT9"] = flush

	table["A543"] = flush
	table["A432"] = flush
	table["A532"] = flush
	table["A542"] = flush

	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			key := ""
			for k := 0; k < 5; k++ {
				if k != j {
					key += str[i+k]
				}
			}
			table[key] = flush
		}
	}
	return table
}

func CompareResult(mode1 int, v1 []int, mode2 int, v2 []int) int {
	if mode1 < mode2 {
		return over
	} else if mode1 > mode2 {
		return less
	} else {
		switch mode1 {
		case royal:
			return equal
		default:
			for i := 0; i < len(v1); i++ {
				if v1[i] > v2[i] {
					return over
				} else if v1[i] < v2[i] {
					return less
				}
			}
		}
	}
	return equal
}

func NewSimpleCards() *SimpleCards {
	ret := &SimpleCards{}
	ret.buf = make([]int, 5)
	ret.table5 = createFullTableWithoutGhost()
	create32Table(ret.table5)
	create4Table(ret.table5)
	create3Table(ret.table5)
	createCouple2Table(ret.table5)
	createCoupleTable(ret.table5)
	return ret
}

func createCoupleTable(result map[string]seqCards) {
	raw := []string{}
	selectCards(4, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]), string(item[0]), string(item[1]), string(item[2]), string(item[3])}
		seq := []int{table[item[1]], table[item[2]], table[item[3]]}
		sortCard(seq, 3)
		sorted := []int{table[item[0]]}
		kind := seqCards{couple, append(sorted, seq...)}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p] = kind
		}
	}
}

func createCouple2Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(3, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]), string(item[0]), string(item[1]), string(item[1]), string(item[2])}
		max, min := table[item[0]], table[item[1]]
		if max < min {
			max, min = min, max
		}
		kind := seqCards{couple2, []int{max, min, table[item[2]]}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p] = kind
		}
	}
}

func create3Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(3, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]), string(item[0]), string(item[0]), string(item[1]), string(item[2])}
		max, min := table[item[1]], table[item[2]]
		if max < min {
			max, min = min, max
		}
		kind := seqCards{three, []int{table[item[0]], max, min}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p] = kind
		}
	}
}

func create4Table(result map[string]seqCards) {
	raw := []string{}
	selectCards(2, "", &raw)
	//fmt.Println(raw)
	for _, item := range raw {
		material := []string{string(item[0]), string(item[0]), string(item[0]), string(item[0]), string(item[1])}
		kind := seqCards{four, []int{table[item[0]], table[item[1]]}}

		possible := []string{}
		permutate(material, 0, 5, &possible)
		for _, p := range possible {
			result[p] = kind
		}
	}
}

func Preprocess(raw string) (string, []int) {
	str := ""
	color := []int{0, 0, 0, 0}
	for i := 0; i < len(raw); i += 2 {
		str += string(raw[i])
		se := table[raw[i+1]]
		color[se]++
	}
	return str, color
}

var table [128]int

func init() {
	table['2'] = 2
	table['3'] = 3
	table['4'] = 4
	table['5'] = 5
	table['6'] = 6
	table['7'] = 7
	table['8'] = 8
	table['9'] = 9
	table['T'] = 10
	table['J'] = 11
	table['Q'] = 12
	table['K'] = 13
	table['A'] = 14

	table['S'] = spades
	table['H'] = hearts
	table['D'] = diamonds
	table['C'] = clubs
	table['s'] = spades
	table['h'] = hearts
	table['d'] = diamonds
	table['c'] = clubs
}

var invTable = map[int]string{
	2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "T", 11: "J", 12: "Q", 13: "K", 14: "A",
}

type cardColor struct {
	count int
	cards []int
}

type seqCards struct {
	mode int
	max  []int
}

type cardBuf struct {
	sOrder     []int
	single     int
	sCcnt      int
	dOrder     []int
	double     int
	dCnt       int
	tOrder     []int
	tripple    int
	tCnt       int
	card4      int
	cardState  int
	color      []cardColor
	same       bool
	tabNoGhost map[string]int
	tabGhost   map[string]int

	tabFullNoGhost map[string]seqCards
}

func (cb *cardBuf) clear() {
	cb.sCcnt = 0
	cb.single = 0
	cb.dCnt = 0
	cb.double = 0
	cb.tCnt = 0
	cb.tripple = 0
	cb.card4 = 0
	cb.color[diamonds].count, cb.color[clubs].count, cb.color[hearts].count, cb.color[spades].count = 0, 0, 0, 0
}

func sortCard(good []int, length int) {
	for j := 0; j < length-1; j++ {
		max := good[j]
		pivot := j
		for i := j + 1; i < length; i++ {
			if good[i] > max {
				max = good[i]
				pivot = i
			}
		}
		good[pivot] = good[j]
		good[j] = max
	}
}

func (cb *cardBuf) addCard(cards string, length int) bool {
	ghost := false
	for i := 0; i < length; i += 2 {
		if string(cards[i]) == "X" {
			ghost = true
			continue
		}
		card := table[cards[i]]
		se := table[cards[i+1]]
		cb.color[se].cards[cb.color[se].count] = card
		cb.color[se].count++

		off := 1 << uint(card)
		if (cb.single & off) == 0 {
			cb.single |= off
			cb.sOrder[cb.sCcnt] = card
			cb.sCcnt++
		} else if (cb.double & off) == 0 {
			cb.double |= off
			cb.dOrder[cb.dCnt] = card
			cb.dCnt++
		} else if (cb.tripple & off) == 0 {
			cb.tripple |= off
			cb.tOrder[cb.tCnt] = card
			cb.tCnt++
		} else {
			cb.card4 = card
		}
	}
	return ghost
}

func (cb *cardBuf) check4ColorSame() (bool, int, []int) {
	if cb.color[spades].count > 3 {
		return true, cb.color[spades].count, cb.color[spades].cards
	}
	if cb.color[hearts].count > 3 {
		return true, cb.color[hearts].count, cb.color[hearts].cards
	}
	if cb.color[clubs].count > 3 {
		return true, cb.color[clubs].count, cb.color[clubs].cards
	}
	if cb.color[diamonds].count > 3 {
		return true, cb.color[diamonds].count, cb.color[diamonds].cards
	}
	return false, 0, nil
}

func (cb *cardBuf) check5ColorSame() (bool, int, []int) {
	if cb.color[spades].count > 4 {
		return true, cb.color[spades].count, cb.color[spades].cards
	}
	if cb.color[hearts].count > 4 {
		return true, cb.color[hearts].count, cb.color[hearts].cards
	}
	if cb.color[clubs].count > 4 {
		return true, cb.color[clubs].count, cb.color[clubs].cards
	}
	if cb.color[diamonds].count > 4 {
		return true, cb.color[diamonds].count, cb.color[diamonds].cards
	}
	return false, 0, nil
}

func combineKey(dat ...int) string {
	length := len(dat)
	str := ""
	for i := 0; i < length; i++ {
		str += invTable[dat[i]]
	}
	return str
}

func (cb *cardBuf) checkSingleCards(ghost bool) (mode int, ret int) {
	dat := cb.sOrder
	length := cb.sCcnt
	if length < 4 || (length == 4 && !ghost) {
		// 有赖子得满4张 || 无赖子得满5张
		return alone, dat[0]
	}

	if ghost {
		for i := 0; i < length-3; i++ {
			v, ok := cb.tabGhost[combineKey(dat[i], dat[i+1], dat[i+2], dat[i+3])]
			if ok {
				if dat[i] == 14 && v == flush {
					// A5432
					return sequence, 5
				}
				if v == royal {
					return sequence, 14
				}
				if dat[i] < 14 && (dat[i+3]-dat[i] == 3) {
					// 赖子插入位置， 头部
					return sequence, dat[i] + 1
				}
				return sequence, dat[i]
			}
		}
	} else {
		for i := 0; i < length-4; i++ {
			v, ok := cb.tabNoGhost[combineKey(dat[i], dat[i+1], dat[i+2], dat[i+3], dat[i+4])]
			if ok {
				if dat[i] == 14 && v == flush {
					// A5432
					return sequence, 5
				}
				return sequence, dat[i]
			}
		}
	}

	return alone, dat[0]
}

func (cb *cardBuf) checkBomb(ghost bool) (bool, int) {
	if cb.card4 > 0 {
		return true, cb.card4
	}
	if ghost {
		if cb.tCnt > 0 {
			return true, cb.tOrder[0]
		}
	}
	return false, 0
}

func (cb *cardBuf) checkThreeTwo(ghost bool) (bool, int, int) {
	if cb.tCnt > 0 {
		z3 := cb.tOrder[0]
		max := 0
		for i := 0; i < cb.dCnt; i++ {
			if z3 != cb.dOrder[i] {
				max = cb.dOrder[i]
				break
			}
		}
		if ghost {
			//3+1+赖子
			for i := 0; i < cb.sCcnt; i++ {
				if cb.sOrder[i] != z3 {
					if max < cb.sOrder[i] {
						max = cb.sOrder[0]
						break
					}
				}
			}
		}
		if max > 0 {
			return true, z3, max
		}
	} else {
		if ghost {
			// 2+2+赖子
			if cb.dCnt > 1 {
				return true, cb.dOrder[0], cb.dOrder[1]
			}
		}
	}

	return false, 0, 0
}

func (cb *cardBuf) checkThree(ghost bool) (bool, int) {
	max := 0
	if cb.tCnt > 0 {
		max = cb.tOrder[0]
	}
	if ghost {
		for i := 0; i < cb.dCnt; i++ {
			if max < cb.dOrder[i] {
				max = cb.dOrder[i]
			}
		}
	}
	if max > 0 {
		return true, max
	}
	return false, 0
}

func (cb *cardBuf) check2Couple(ghost bool) (bool, int, int) {
	first := 0
	second := 0
	if cb.dCnt > 1 {
		first, second = cb.dOrder[0], cb.dOrder[1]
	} else if cb.dCnt > 0 {
		first = cb.dOrder[0]
	}
	if first > 0 && second > 0 {
		return true, first, second
	}
	return false, 0, 0
}

func (cb *cardBuf) checkCouple(ghost bool) (bool, int) {
	max := 0
	if cb.dCnt > 0 {
		max = cb.dOrder[0]
	}
	if ghost {
		if cb.sOrder[0] > max {
			return true, cb.sOrder[0]
		}
	}
	if max > 0 {
		return true, max
	}
	return false, 0
}

func (cb *cardBuf) checkType(ghost bool) (int, []int) {
	m2, v2 := cb.checkBomb(ghost)
	if m2 {
		ret := []int{v2, 0}
		for i := 0; i < len(cb.sOrder); i++ {
			if cb.sOrder[i] != v2 {
				ret[1] = cb.sOrder[i]
				return four, ret
			}
		}
	}

	m3, v31, v32 := cb.checkThreeTwo(ghost)
	if m3 {
		return threeTwo, []int{v31, v32}
	}

	m1, v1 := cb.checkSingleCards(ghost)
	if m1 < four {
		return m1, []int{v1}
	}

	if m1 == suit {
		ret := []int{v1, 0, 0, 0, 0}
		j := 1
		for i := 0; i < len(cb.sOrder); i++ {
			if v1 != cb.sOrder[i] {
				ret[j] = cb.sOrder[i]
				j++
				if j == 5 {
					return m1, ret
				}
			}
		}
	}
	if m1 == sequence {
		return m1, []int{v1}
	}

	m4, v4 := cb.checkThree(ghost)
	if m4 {
		ret := []int{v4, 0, 0}
		j := 1
		for i := 0; i < len(cb.sOrder); i++ {
			if cb.sOrder[i] != v4 {
				ret[j] = cb.sOrder[i]
				j++
				if j == 3 {
					return three, ret
				}
			}
		}
	}

	m5, v51, v52 := cb.check2Couple(ghost)
	if m5 {
		ret := []int{v51, v52, 0}
		for i := 0; i < len(cb.sOrder); i++ {
			if cb.sOrder[i] != v51 && cb.sOrder[i] != v52 {
				ret[2] = cb.sOrder[i]
				return couple2, ret
			}
		}
	}

	m6, v6 := cb.checkCouple(ghost)
	if m6 {
		ret := []int{v6, 0, 0, 0}
		j := 1
		for i := 0; i < len(cb.sOrder); i++ {
			if cb.sOrder[i] != v6 {
				ret[j] = cb.sOrder[i]
				j++
				if j == 4 {
					return couple, ret
				}
			}
		}
	}

	ret := []int{0, 0, 0, 0, 0}
	for i := 0; i < 5; i++ {
		ret[i] = cb.sOrder[i]
	}

	if ghost {
		ret[0] = 14
	}
	return alone, ret
}

func TestAdd(cb1, cb2 *cardBuf, cards1, cards2 string) {
	length := len(cards1)
	cb1.addCard(cards1, length)
	cb2.addCard(cards2, length)

	cb1.clear()
	cb2.clear()
}

func Process(cb1, cb2 *cardBuf, cards1, cards2 string) (int, []int, int, []int) {
	length := len(cards1)
	ghost1 := cb1.addCard(cards1, length)
	ghost2 := cb2.addCard(cards2, length)
	sortCard(cb1.dOrder, cb1.dCnt)
	sortCard(cb1.sOrder, cb1.sCcnt)
	sortCard(cb2.dOrder, cb2.dCnt)
	sortCard(cb2.sOrder, cb2.sCcnt)

	mode1, v1 := cb1.checkType(ghost1)
	cb1.clear()
	mode2, v2 := cb2.checkType(ghost2)
	cb2.clear()
	return mode1, v1, mode2, v2
}

func (cb *cardBuf) process(cards string) (int, []int) {
	length := len(cards)
	ghost := cb.addCard(cards, length)
	sortCard(cb.dOrder, cb.dCnt)
	sortCard(cb.sOrder, cb.sCcnt)
	mode, v := cb.checkType(ghost)
	cb.clear()
	return mode, v
}

func (cb *cardBuf) check5SingleCardsOnlyWithoutGhost() (mode int, ret int) {
	dat := cb.sOrder
	length := cb.sCcnt
	if length != 5 {
		// 无赖子得满5张
		return alone, dat[0]
	}
	if iSame, _, candidate := cb.check5ColorSame(); iSame {
		v, ok := cb.tabFullNoGhost[combineKey(candidate[0], candidate[1], candidate[2], candidate[3], candidate[4])]
		if ok {
			return v.mode, v.max[0]
		} else {
			return suit, candidate[0]
		}
	}
	v, ok := cb.tabFullNoGhost[combineKey(dat[0], dat[1], dat[2], dat[3], dat[4])]
	if ok {
		return sequence, v.max[0]
	}
	return alone, dat[0]
}

type SimpleCards struct {
	cards  seqCards
	buf    []int
	color  []int
	table5 map[string]seqCards
}

func (cb *SimpleCards) add5Card(cards string, se []int) {
	cb.color = se

	if v, ok := cb.table5[cards]; ok {
		cb.cards = v
	} else {
		for i := 0; i < 5; i++ {
			card := table[cards[i]]
			cb.buf[i] = card
		}
		sortCard(cb.buf, 5)
		cb.cards = seqCards{alone, cb.buf}
	}

}

func (cb *SimpleCards) checkColor() bool {
	if cb.color[spades] == 5 || cb.color[hearts] == 5 || cb.color[diamonds] == 5 || cb.color[clubs] == 5 {
		return true
	}
	return false
}

func (cb *SimpleCards) checkType() (int, []int) {
	dat := make([]int, 0)
	if cb.checkColor() {
		if cb.cards.mode == royal || cb.cards.mode == flush {
			return cb.cards.mode, append(dat, cb.cards.max...)
		}
		return suit, append(dat, cb.cards.max...)
	}

	if cb.cards.mode == royal || cb.cards.mode == flush {
		return sequence, append(dat, cb.cards.max...)
	}

	return cb.cards.mode, append(dat, cb.cards.max...)
}

func (cb *SimpleCards) Process(hand string, se []int) (int, []int) {
	cb.add5Card(hand, se)
	return cb.checkType()
}
