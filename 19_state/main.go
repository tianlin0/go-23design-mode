package main

import "fmt"

// ActionState 定义状态接口：每个状态可以对应那些动作
type ActionState interface {
	View()   //查看
	Comment() //评论
	Post()   //提交
}

type NormalState struct {
}
func (n *NormalState) Post() {
	fmt.Println("正常发帖")
}
func (n *NormalState) View() {
	fmt.Println("正常看帖")
}
func (n *NormalState) Comment() {
	fmt.Println("正常评论")
}

type CloseState struct {
}
func (n *CloseState) Post() {
	fmt.Println("发帖")
}
func (n *CloseState) View() {
	fmt.Println("看帖")
}
func (n *CloseState) Comment() {
	fmt.Println("评论")
}



// Account 定义账户结构体：包含当前账户状态State、HealthValue账号健康值
type Account struct {
	State       ActionState
	HealthValue int
}

func NewAccount(health int) *Account {
	a := new(Account)
	a.SetHealth(health)
	return a
}

func (a *Account) SetHealth(value int) {
	a.HealthValue = value
	a.changeState()
}

func (a *Account) changeState() {
	if a.HealthValue > 0 {
		a.State = &NormalState{}
	}else{
		a.State = &CloseState{}
	}
}

func (a *Account) Post() {
	a.State.Post()
}
func (a *Account) View() {
	a.State.View()
}
func (a *Account) Comment() {
	a.State.Comment()
}



func main() {
	fmt.Println("===========正常账户===========")
	//正常账户：可发帖、评论、查看
	account := NewAccount(10)
	account.Post()
	account.View()
	account.Comment()
	fmt.Println("===========受限账户===========")
	//受限账户：不能发帖、可以评论和查看
	account.SetHealth(-5)
	account.Post()
	account.View()
	account.Comment()
	fmt.Println("===========被封号账户===========")
	//被封号账户：不能发帖、不能评论、不能查看
	account.SetHealth(-11)
	account.Post()
	account.View()
	account.Comment()
}
