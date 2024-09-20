package action

import "fmt"

// Machine 状态机
type Machine struct {
	state IState
}

// SetState 更新状态
func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// IState 状态
type IState interface {
	// Approval 审批通过
	Approval(m *Machine)
	// Reject 驳回
	Reject(m *Machine)
}

// leaderApproveState 直属领导审批
type leaderApproveState struct{}

// Approval 获取状态名字
func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}

// Reject 获取状态名字
func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

// financeApproveState 财务审批
type financeApproveState struct{}

// Approval 审批通过
func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}

// Reject 拒绝
func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

// GetFinanceApproveState GetFinanceApproveState
func GetFinanceApproveState() IState {
	return &financeApproveState{}
}
