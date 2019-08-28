package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mahendrakariya/mainsvc/add"
	"github.com/mahendrakariya/mainsvc/divide"
	"google.golang.org/grpc"
)

func main() {
	http_server()
}

type Output struct {
	Sum      string `json:"sum"`
	Diff     string `json:"diff"`
	Product  string `json:"product"`
	Quotient string `json:"quotient"`
}

type Numbers struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}

func http_server() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var n Numbers
		err := decoder.Decode(&n)
		if err != nil {
			panic(err)
		}

		o := &Output{}

		sum, err := add_grpc(n.A, n.B)
		if err != nil {
			o.Sum = err.Error()
		} else {
			o.Sum = strconv.Itoa(int(sum))
		}

		diff, err := subtract_http(n.A, n.B)
		if err != nil {
			o.Diff = err.Error()
		} else {
			o.Diff = strconv.Itoa(int(diff))
		}

		product, err := multiply_http(n.A, n.B)
		if err != nil {
			o.Product = err.Error()
		} else {
			o.Product = strconv.Itoa(int(product))
		}

		quotient, err := divide_grpc(n.A, n.B)
		if err != nil {
			o.Quotient = err.Error()
		} else {
			o.Quotient = strconv.Itoa(int(quotient))
		}

		json.NewEncoder(w).Encode(o)
	}).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func add_grpc(a, b int32) (int32, error) {
	add_grpc_address := "add-svc:50051"

	// Set up a connection to the server.
	conn, err := grpc.Dial(add_grpc_address, grpc.WithInsecure())
	if err != nil {
		return 0, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := add.NewAdderClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DoSum(ctx, &add.Numbers{A: a, B: b})
	if err != nil {
		return 0, fmt.Errorf("Could not add numbers: %v", err)
	}
	return r.GetSum(), nil
}

func divide_grpc(a, b int32) (int32, error) {
	divide_grpc_address := "divide-svc:50052"
	// Set up a connection to the server.
	conn, err := grpc.Dial(divide_grpc_address, grpc.WithInsecure())
	if err != nil {
		return 0, fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := divide.NewDivideClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DoDivide(ctx, &divide.Numbers{A: a, B: b})
	if err != nil {
		return 0, fmt.Errorf("Could not divide numbers: %v", err)
	}
	return r.GetQ(), nil
}

type ProductResponse struct {
	Product int32 `json:"product"`
}

func multiply_http(a, b int32) (int32, error) {
	multiply_http_address := fmt.Sprintf("http://%s:8001/multiply", "multiply-svc.default.svc.cluster.local")
	url := multiply_http_address
	fmt.Println(url)
	var jsonStr = []byte(fmt.Sprintf("{\"a\": %v, \"b\": %v}", a, b))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	fmt.Printf("Req: %v\n", req)
	fmt.Printf("Req Header: %v\n", req.Header)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var p ProductResponse
	json.Unmarshal(body, &p)
	return p.Product, nil
}

type SubtractResponse struct {
	Diff int32 `json:"diff"`
}

func subtract_http(a, b int32) (int32, error) {
	multiply_http_address := fmt.Sprintf("http://%s:4567/subtract", "subtract-svc.default.svc.cluster.local")

	var jsonStr = []byte(fmt.Sprintf("{'a': %v, 'b': %v}", a, b))

	req, err := http.NewRequest("POST", multiply_http_address, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var p SubtractResponse
	json.Unmarshal(body, &p)
	return p.Diff, nil
}
