package main

import "fmt"

// 命令模式

// 核心计算模块 doctor 命令的接收者
type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 抽象的命令
type Command interface {
	Treat()
}

type CommandTreatEye struct {
	Doctor *Doctor
}

// 病单 命令
func (cmd *CommandTreatEye) Treat() {
	cmd.Doctor.treatEye()
}

type CommandTreatNose struct {
	Doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.Doctor.treatNose()
}

// 护士 命令的发送者
type Nurse struct {
	CmdList []Command
}

// 统一发送命令
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}
func main() {
	// doctor := new(Doctor)
	// doctor.treatEye()
	// doctor.treatNose()
	doctor := new(Doctor)
	cmdEye := &CommandTreatEye{Doctor: doctor}
	//cmdEye.Treat()
	cmdNose := &CommandTreatNose{Doctor: doctor}
	//cmdNose.Treat()
	nurse := new(Nurse)
	//添加指令
	nurse.CmdList = append(nurse.CmdList, cmdNose)
	nurse.CmdList = append(nurse.CmdList, cmdEye)
	// 发送指令给Doctor核心模块
	nurse.Notify()
}
