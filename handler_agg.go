package main

import (
	"context"
	"fmt"
)

const url = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	data, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(data)

	return nil
}