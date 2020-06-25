package factorymethod

type ak47 struct {
	gun // Go中不存在继承 所以使用匿名组合来实现
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
