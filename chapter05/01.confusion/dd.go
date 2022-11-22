package main

type downloadFromNetDisk struct {
	secret   DynamicSecret
	filepath string
}

func (d *downloadFromNetDisk) DownloadFile() (file string) {
	if err := d.loginCheck(); err != nil {
		// todo 重新登录
	}
	d.downloadFromAliyunNetDisk(file)
	return ""
}

func (d *downloadFromNetDisk) loginCheck() error {
	d.checkSecret(d.secret.GetSecret())
	return nil
}

func (d *downloadFromNetDisk) downloadFromAliyunNetDisk(file string) {

}

func (d *downloadFromNetDisk) checkSecret(secret string) {
	// todo 调用阿里云的接口去验证密码是否有效
}
