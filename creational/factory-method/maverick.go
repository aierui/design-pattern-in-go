package factorymethod

type maverick struct {
	gun // Go中不存在继承 所以使用匿名组合来实现
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 4,
		},
	}
}
