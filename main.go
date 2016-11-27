package main

import (
	"bitbucket.org/copyninja/go9p/p"
	"bitbucket.org/copyninja/go9p/p/srv"
	"github.com/golang/glog"
	"os"
)

func main() {
	root, err := mkCalcfs()
	if err != nil {
		glog.Errorln("Failed to create root file: ", err)
		os.Exit(2)
	}

	s := srv.NewFileSrv(root)
	s.Dotu = true
	s.Debuglevel = 1
	s.Start(s)

	if err = s.StartNetListener("tcp", "0.0.0.0:5640"); err != nil {
		glog.Errorf("Listener failed to start (%v)", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func mkCalcfs() (*srv.File, error) {
	var err error

	user := p.OsUsers.Uid2User(os.Geteuid())
	root := new(srv.File)

	err = root.Add(nil, "/", user, nil, p.DMDIR|0755, nil)
	if err != nil {
		return nil, err
	}

	err = mkCtlFile(root, user)
	if err != nil {
		return nil, err
	}

	err = mkDataFile(root, user)
	if err != nil {
		return nil, err
	}

	return root, nil
}
