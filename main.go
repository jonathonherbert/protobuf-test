package main

import (
	"bufio"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	var data []byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}

	recipeHashes := &RecipeHashes{}
	err := protojson.Unmarshal(data, recipeHashes)

	if err != nil {
		print(string(data))
		panic(err)
	}

	log.Printf("Size of msg: %d", proto.Size(recipeHashes))
}
