package main

/*
	封装NewXXX函数
*/

type IDownload interface {
	Download()
}

type SmbDownloader struct{}

func (s *SmbDownloader) Download() {
	println("smb download")
}

type NfsDownloader struct{}

func (n *NfsDownloader) Download() {
	println("nfs download")
}

func NewDownloader(t string) IDownload {
	switch t {
	case "smb":
		return new(SmbDownloader)
	case "nfs":
		return new(NfsDownloader)
	}
	return nil
}

func main() {
	//测试：根据协议类型，创建不同类型的下载器
	smbDownloader := NewDownloader("smb")
	smbDownloader.Download()
}
