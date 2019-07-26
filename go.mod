module ServerUtils

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190426145343-a29dc8fdc734
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190502175342-a43fa875dd82
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190503185657-3b6f9c0030f7

)

require (
	github.com/astaxie/beego v1.11.1
	github.com/go-sql-driver/mysql v1.4.1
	github.com/kr/fs v0.1.0 // indirect
	github.com/pkg/sftp v1.10.0
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)
