module server

go 1.13

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.43.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190731235908-ec7cb31e5a56
	golang.org/x/image => github.com/golang/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190719004257-d2bd2a29d028
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190801041406-cbf593c0f2f3
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20180917221912-90fa682c2a6e
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190801165951-fa694d86fc64
	google.golang.org/grpc => github.com/google/grpc v1.19.0
)

require (
	github.com/astaxie/beego v1.12.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/goconvey v1.6.4
	google.golang.org/appengine v1.6.5 // indirect
)
