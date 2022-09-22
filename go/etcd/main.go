package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	listName := flag.Bool("list-name", false, "查看所有存储键名")
	prefix := flag.String("prefix", "", "列出给定字符开头的内容")
	endPoint := flag.String("endpoint", "https://192.168.49.2:2379", "etcd服务地址")
	flag.Parse()

	// 使用etcd根证书认证
	caCert, err := ioutil.ReadFile("ca.crt")
	checkErr(err)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	serverCert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	checkErr(err)

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{*endPoint},
		DialTimeout: 5 * time.Second,
		TLS: &tls.Config{
			RootCAs:      caCertPool,
			Certificates: []tls.Certificate{serverCert},
		},
	})

	checkErr(err)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := client.Get(ctx, "", clientv3.WithPrefix())
	checkErr(err)
	cancel()

	for _, r := range resp.Kvs {
		if *listName {
			fmt.Println(string(r.Key))
			continue
		}
		if strings.Contains(string(r.Key), *prefix) {
			fmt.Println(string(r.Key))
			fmt.Println("-------------------------------------------------------------------------------")
			fmt.Println(string(r.Value))
			fmt.Println("-------------------------------------------------------------------------------")
			fmt.Println()
		}
	}
	checkErr(err)
}
