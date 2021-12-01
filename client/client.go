package main

import (
	"context"
	"encoding/json"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
	"dubbo3-demo/api"
)

var greeterProvider = &api.UserProviderClientImpl{}


func init() {
	// validate consumer greeterProvider ptr
	config.SetConsumerService(greeterProvider)
}

func main() {
	// init rootConfig with config api
	rc := config.NewRootConfigBuilder().
		SetConsumer(config.NewConsumerConfigBuilder().
			AddReference("UserProviderClientImpl", config.NewReferenceConfigBuilder().
				SetProtocol("tri").
				Build()).
			Build()).
		AddRegistry("zookeeper", config.NewRegistryConfigWithProtocolDefaultPort("zookeeper")).
		Build()

	// start dubbo-go framework with configuration
	if err := config.Load(config.WithRootConfig(rc)); err != nil{
		panic(err)
	}

	// run rpc invocation
	testGetUser()
}

func testSayHello() {
	ctx := context.Background()

	req := api.HelloRequest{
		Name: "laurence",
	}
	user, err := greeterProvider.SayHello(ctx, &req)
	if err != nil {
		panic(err)
	}

	logger.Infof("Receive user = %+v\n", user)
}

func testGetUser() {
	ctx := context.Background()

	req := api.GetUserReq{
		Name: "laurence",
	}
	user, err := greeterProvider.GetUser(ctx, &req)
	if err != nil {
		panic(err)
	}

	data, _ := json.Marshal(user)
	logger.Infof("Receive user = %+v\n", string(data))
}