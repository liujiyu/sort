package main

import (
	"fmt"
	"sort"
)

type Rounds []*Round

type Round struct {
	ante           int     //底注
	loseExp        int     //输的时候添加的经验值
	winExp         int     //赢牌经验值
	winCoefficient float64 //赢的时候 添加的经验值系数
}

func (r Rounds) Less(i, j int) bool {
	return r[i].ante < r[j].ante
}
func (r Rounds) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Rounds) Len() int {
	return len(r)
}

type Game struct {
	rounds Rounds
}

func main() {
	game := new(Game)
	round1 := new(Round)
	round1.ante = 10

	round2 := new(Round)
	round2.ante = 8

	round3 := new(Round)
	round3.ante = 15

	game.rounds = append(game.rounds, round1, round2, round3)

	sort.Sort(game.rounds)

	fmt.Println(game.rounds[0].ante, game.rounds[1].ante, game.rounds[2].ante)
}
