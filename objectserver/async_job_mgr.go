package objectserver

import (
	"flag"

	"github.com/iqiyi/auklet/common/conf"
	"github.com/iqiyi/auklet/common/fs"
)

const (
	BLOOMFILTER_ENTRIES  = 8096.0
	BLOOMFILTER_FP_RATIO = 0.01
)

type AsyncJob interface {
	GetMethod() string
	GetHeaders() map[string]string
	GetAccount() string
	GetContainer() string
	GetObject() string
}

type AsyncJobMgr interface {
	New(vars, headers map[string]string) AsyncJob

	Save(job AsyncJob) error

	Next(device string, policy int) AsyncJob

	Finish(job AsyncJob) error
}

func startKVRpcService(cnf conf.Config, flags *flag.FlagSet) {
	driveRoot := cnf.GetDefault("app:object-server", "devices", "/srv/node")

	ringPort := int(cnf.GetInt("DEFAULT", "bind_port", 6000))
	kv := NewKVStore(driveRoot, ringPort)
	test := cnf.GetBool("app:object-server", "test_mode", false)
	if !test {
		m := fs.NewMountMonitor()
		m.RegisterCallback("async-job-mgr", kv.mountListener)
		go m.Start()
	} else {
		kv.setTestMode(true)
	}

	var rpcSvc *KVService
	rpcPort := int(cnf.GetInt("app:object-server", "async_kv_service_port", 60001))
	if cnf.GetBool("app:object-server", "async_kv_fs_compatible", false) {
		fs := NewFSStore(driveRoot)
		rpcSvc = NewKVFSService(fs, kv, rpcPort)
	} else {
		rpcSvc = NewKVService(kv, rpcPort)
	}

	go rpcSvc.start()
}

func initKVAsyncJobMgr(
	cnf conf.Config, flags *flag.FlagSet) (*KVAsyncJobMgr, error) {
	startKVRpcService(cnf, flags)
	port := int(cnf.GetInt("app:object-server", "async_kv_service_port", 60001))
	return NewKVAsyncJobMgr(port)
}

func initFSAsyncJobMgr(
	cnf conf.Config, flags *flag.FlagSet) (*FSAsyncJobMgr, error) {
	driveRoot := cnf.GetDefault("app:object-server", "devices", "/srv/node")
	return NewFSAsyncJobMgr(driveRoot)
}

func NewAsyncJobMgr(cnf conf.Config, flags *flag.FlagSet) (AsyncJobMgr, error) {
	mgr := cnf.GetDefault("app:object-server", "async_job_manager", "fs")
	switch mgr {
	case "fs":
		return initFSAsyncJobMgr(cnf, flags)
	case "kv":
		return initKVAsyncJobMgr(cnf, flags)
	default:
		return nil, ErrUnknownAsyncJobMgr
	}
}
