package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"

	"google.golang.org/grpc"
)

var (
	dgraph = flag.String("d", "127.0.0.1:9080", "Dgraph Alpha address")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	_uid := "0x1, 0x22, 0xbb1"
	_var := `{
		genre(func: uid(` + _uid + `)) {
		  uid 
		  genre.name
		  genre.id
		}
	  }`

	resp, err := dg.NewTxn().Query(context.Background(), _var)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)
}
