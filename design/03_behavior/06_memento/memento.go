package memento

// 备忘录模式
// 三部分组成备忘录对象，包含状态的对象，备忘录管理器
// 备忘录管理器负责创建备忘录和从备忘录中取出状态对象，备忘录负责存储状态对象

// Stater 状态对象接口
type Stater interface {
	SetState(string)
	GetState() string
}

// State 状态对象
type State struct {
	Text string
}

func (s *State) SetState(text string) {
	s.Text = text
}

func (s *State) GetState() string {
	return s.Text
}

// Memo 备忘录对象
type Memo struct {
	Stater
}

// MemoManager 备忘录管理器
type MemoManager struct {
}

// Create 创建备忘录
func (m *MemoManager) Create(s Stater) *Memo {
	return &Memo{s}
}

// GetStater 从备忘录获取记录的状态对象
func (m *MemoManager) GetStater(memo *Memo) Stater {
	return memo.Stater
}
