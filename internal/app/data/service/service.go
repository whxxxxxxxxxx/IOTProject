package service

import (
	"fmt"
	"os/exec"
	"strconv"
)

func ConvertTimeToSeconds(timeStr string) (string, error) {
	lastChar := timeStr[len(timeStr)-1]
	valueStr := timeStr[:len(timeStr)-1]
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return "", err
	}
	var secondData int64

	switch lastChar {
	case 'm':
		if value < 2 {
			return "", fmt.Errorf("时间间隔太短，请至少一次获取2分钟以上的数据", value)
		}
		secondData = value * 60
	case 'h':
		secondData = value * 3600
	case 'd':
		secondData = value * 86400
	case 'w':
		secondData = value * 604800
	case 'n':
		secondData = value * 2592000 // Approximation: average seconds in a month
	case 'y':
		secondData = value * 31536000 // Approximation: average seconds in a year
	}
	//整除100
	splitData := secondData / 100
	return fmt.Sprintf("%ds", splitData), nil
}

func SaveDataToDB() error {
	cmd := "mqttx"
	num := "2"
	args := []string{"simulate", "--file", "industrial.js", "-c", num, "-h", "127.0.0.1", "-t", "mqttx/iot"}
	command := exec.Command(cmd, args...)
	// 启动命令
	err := command.Start()
	if err != nil {
		return err
	}
	fmt.Println("Command started. Waiting for 5 seconds before stopping...")
	return nil
}
