package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/tonkeeper/tongo/liteapi"
)

func main() {
	godotenv.Load()
	client, err := liteapi.NewClient(liteapi.FromEnvsOrMainnet())
	if err != nil {
		log.Fatalln(err)
	}
	it := NewElectorRoundsIterator(context.Background(), client)
	for v := range it.Run {
		fmt.Printf("Election ID: %v\n", v.StartTime().Unix())
		fmt.Printf("Round Duration: %v\n", v.Duration())
		fmt.Printf("Total Stake: %.2f TON\n", float64(v.TotalStake)/1e9)
		if time.Now().Before(v.EndTime()) {
			fmt.Printf("Projected Bonuses: %.2f TON\n", v.ProjectedBonuses()/1e9)
			fmt.Printf("Projected APY: %.2f%%\n", v.ProjectedAPY()*100)
		} else {
			fmt.Printf("Bonuses: %.2f TON\n", float64(v.Bonuses)/1e9)
			fmt.Printf("APY: %.2f%%\n", v.APY()*100)
		}
		fmt.Println()
	}
	if it.Err() != nil {
		log.Fatalln(it.Err())
	}
}
