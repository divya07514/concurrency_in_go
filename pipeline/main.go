package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func ScatterGather(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	ab := NewABProcessor()
	ab.start(ctx, data)
	inputC, err := ab.wait(ctx)
	if err != nil {
		return COut{}, err
	}
	c := newCProcessor()
	c.start(ctx, inputC)
	wait, err := c.wait(ctx)
	if err != nil {
		return COut{}, err
	}
	return wait, nil
}

func main() {

	cout, err := ScatterGather(context.Background(), Input{
		A: "Data for A",
		B: "Data for B",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cout)
}
