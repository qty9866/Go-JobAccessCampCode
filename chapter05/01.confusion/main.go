package main

func main() {
	d := downloadFromNetDisk{
		secret:   &mobileTokenDynamic{mobileNum: "17805101835"},
		filepath: "/src/Grand Theft Auto VI",
	}
	d.DownloadFile()
}

type DynamicSecret interface {
	GetSecret() string
}

type mobileTokenDynamic struct {
	mobileNum string
}

func (m *mobileTokenDynamic) GetSecret() string {
	// todo 获取密码
	return "something"
}

// 通常开发的时候，第一个版本都叫做happy path
// 剩下的是痛点：无法应对变更，简单的变更会带来更加痛苦的维护
