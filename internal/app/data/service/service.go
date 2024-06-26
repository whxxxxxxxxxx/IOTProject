package service

import (
	"IOTProject/internal/app/data/dao"
	"IOTProject/internal/app/data/dto"
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/exec"
	"time"
)

type Cmd struct {
	Cmd      *exec.Cmd
	JSCancel context.CancelFunc
	Ctx      context.Context
}

var CmdStruct *Cmd

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	var Data dto.Data
	err := json.Unmarshal(msg.Payload(), &Data)
	if err != nil {
		fmt.Println("Unmarshal error")
		return
	}

	Data.Status.Ts = time.Unix(Data.TimeStamp, 0)
	Data.PerformanceMetrics.Ts = time.Unix(Data.TimeStamp, 0)

	go dao.Data.InsertDataById(Data)
	if err != nil {
		fmt.Println("InsertDataById error")
		return
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func SaveDataToDB() error {
	var broker = "tcp://127.0.0.1:1883"
	var topic = "mqttx/iot"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("user")
	opts.SetPassword("password")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	client.Subscribe(topic, 1, nil)
	return nil
}
