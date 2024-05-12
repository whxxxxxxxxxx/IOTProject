package service

import (
	"IOTProject/internal/app/device/dao"
	"IOTProject/internal/app/device/model"
	"IOTProject/kernel"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func UpdateDevicesList() {
	var deviceIds []string
	err := dao.Device.Model(&model.Device{}).
		Select("id").
		Find(&deviceIds).Error
	if err != nil {
		return
	}
	// 将设备ID列表转换为逗号分隔的字符串
	joinedIds := strings.Join(deviceIds, ",")

	// 打开或创建文件用于写入
	file, err := os.OpenFile("device_ids.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		// 如果文件操作出错，打印错误并退出函数
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// 写入数据到文件
	_, err = file.WriteString(joinedIds)
	if err != nil {
		// 如果写入失败，打印错误
		fmt.Println("Error writing to file:", err)
	}

}

func RestartCmd() error {
	var err error
	cmd := kernel.Kernel.JSCmd.Path
	args := kernel.Kernel.JSCmd.Args[1:]
	num, _ := strconv.Atoi(args[4])
	num = num + 1
	args[4] = strconv.Itoa(num)

	if kernel.Kernel.JSCmd != nil && kernel.Kernel.JSCmd.Process != nil {
		err = kernel.Kernel.JSCmd.Process.Kill()
	}
	if err != nil {
		return err
	}
	kernel.Kernel.JSCmd = nil
	kernel.Kernel.JSCmd = exec.Command(cmd, args...)
	err = kernel.Kernel.JSCmd.Start()
	if err != nil {
		return err
	}
	return nil
}
