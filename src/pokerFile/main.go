package main

import (
	"encoding/json"
	"fmt"
	"pokerv2/src/pokerFile/model"
	"pokerv2/src/pokerFile/poker_server"
	"pokerv2/src/pokerFile/readFile"
	"time"
)

func main() {
	var raw map[string][]model.Round
	inputData := readFile.ReadFile("./src/pokerFile/match.json")
	err := json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions := raw["matches"]
	number := len(questions)
	alices := []string{}
	aliceColor := [][]int{}
	bobs := []string{}
	bobColor := [][]int{}
	for i := 0; i < number; i++ {
		a, ac := poker_server.Preprocess(questions[i].Alice)
		alices = append(alices, a)
		aliceColor = append(aliceColor, ac)

		b, bc := poker_server.Preprocess(questions[i].Bob)
		bobs = append(bobs, b)
		bobColor = append(bobColor, bc)
	}
	mgr := poker_server.NewSimpleCards()

	start := time.Now()
	for i := 0; i < number; i++ {
		aliceMode, aliceV := mgr.Process(alices[i], aliceColor[i])
		bobMode, bobV := mgr.Process(bobs[i], bobColor[i])
		questions[i].Result = poker_server.CompareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("5 张牌无赖子: ", time.Now().Sub(start))

	//...------------------------------------------------------------- 7cards===========================
	inputData = readFile.ReadFile("./src/pokerFile/seven_cards.json")
	err = json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions = raw["matches"]
	mgr1 := poker_server.NewCardBuf()
	mgr2 := poker_server.NewCardBuf()

	number = len(questions)
	start = time.Now()
	for i := 0; i < number; i++ {
		aliceMode, aliceV, bobMode, bobV := poker_server.Process(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		questions[i].Result = poker_server.CompareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("7 张牌无赖子: ", time.Now().Sub(start))

	//====================================================================5 with ghost=============
	inputData = readFile.ReadFile("./src/pokerFile/seven_cards_with_ghost.json")
	err = json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions = raw["matches"]
	mgr1 = poker_server.NewCardBuf()
	mgr2 = poker_server.NewCardBuf()

	number = len(questions)
	start = time.Now()
	for i := 0; i < number; i++ {
		aliceMode, aliceV, bobMode, bobV := poker_server.Process(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		questions[i].Result = poker_server.CompareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("5 张牌有赖子: ", time.Now().Sub(start))
	//====================================================================7 with ghost====================
	inputData = readFile.ReadFile("./src/pokerFile/seven_cards_with_ghost.json")
	err = json.Unmarshal(inputData, &raw)
	if err != nil {
		panic(err)
	}
	questions = raw["matches"]
	mgr1 = poker_server.NewCardBuf()
	mgr2 = poker_server.NewCardBuf()

	number = len(questions)
	start = time.Now()
	for i := 0; i < number; i++ {
		aliceMode, aliceV, bobMode, bobV := poker_server.Process(mgr1, mgr2, questions[i].Alice, questions[i].Bob)
		questions[i].Result = poker_server.CompareResult(aliceMode, aliceV, bobMode, bobV)
	}
	fmt.Println("7 张牌有赖子: ", time.Now().Sub(start))
}
