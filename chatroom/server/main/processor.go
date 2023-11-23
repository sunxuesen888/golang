package main

import (
	"fmt"
	"gotest/chatroom/common/message"
	process2 "gotest/chatroom/server/process"
	"gotest/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes===========", mes)
	switch mes.Type {
	case message.LoginMesType:
		{
			up := &process2.UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessLogin(mes)
			return
		}
	case message.RegisterMesType:
		{
			up := &process2.UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessRegister(mes)
			return
		}
	default:
		{
			fmt.Println("消息类型不存在")
		}

	}
	return
}
func (this *Processor) process2() (err error) {
	//循环客户端发送的消息
	for {
		//这里我们将读取数据包，直接封装成一个函数readpkg(),返回message error
		//闯将一个transfer 实例读包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务端也退出 ")
				return err
			} else {
				fmt.Println("readpkg err", err)
				return err
			}
		}

		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverprocessmess err", err)
		}

	}
}
