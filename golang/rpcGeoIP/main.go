package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/oschwald/geoip2-golang"
)

// Response .
type Response struct {
	Country   string
	Province  string
	City      string
	ISP       string
	TimeZone  string
	Latitude  float64
	Longitude float64
}

// IP2Addr .
type IP2Addr struct {
	db *geoip2.Reader
}

// Agrs .
type Agrs struct {
	IPString string
}

// Address .
func (t *IP2Addr) Address(agr *Agrs, res *Response) error {
	netIP := net.ParseIP(agr.IPString)
	record, err := t.db.City(netIP)
	res.City = record.City.Names["zh-CN"]
	if len(record.Subdivisions) != 0 {
		res.Province = record.Subdivisions[0].Names["zh-CN"]
	}
	res.Country = record.Country.Names["zh-CN"]
	res.Latitude = record.Location.Latitude
	res.Longitude = record.Location.Longitude
	return err
}

func main() {
	go Server()
	Client()
}

// Server .
func Server() {
	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	ip2addr := &IP2Addr{db}

	rpc.Register(ip2addr)

	address := ":3344"
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalln("ResolveIPAddr Error: ", address)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln("ListenTCP Error: ", tcpAddr)
	}
	log.Println("json rpc is listening", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

// Client .
func Client() {
	fmt.Println(`
	Enter ip to view ip info
	example:
		Enter IP-> 192.111.1.1
		{Country:美国 Province:加利福尼亚州 City:圣克拉拉 ISP: TimeZone: Latitude:37.3961 Longitude:-121.9617}
	q -- exit
		`)
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:3344")
	if err != nil {
		log.Fatalln("dialing: ", err)
	}

	r := bufio.NewReader(os.Stdin)

	var res Response

	for {
		fmt.Print("Enter IP-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" {
			break
		}

		if err := client.Call("IP2Addr.Address", Agrs{line}, &res); err != nil {
			log.Println("arith error: ", err)
			continue
		}
		fmt.Printf("%+v\n", res)
	}

	// var res Response
	// err = client.Call("IP2Addr.Address", Agrs{"219.140.227.235"}, &res)
	// if err != nil {
	// 	log.Fatalln("arith error: ", err)
	// }
	// fmt.Printf("%+v", res)
}
