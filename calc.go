package main

import (
	"bitbucket.org/copyninja/go9p/p"
	"bitbucket.org/copyninja/go9p/p/srv"
	"github.com/golang/glog"
)

type calcdir struct {
	srv.File
	user p.User
}

type Output struct {
	d      *dataFile
	result string
}

var op Output

type ctlFile struct {
	srv.File
}

type dataFile struct {
	srv.File
}

func mkCtlFile(root *srv.File, user p.User) error {
	glog.V(4).Infof("Entering mkCtlFile(%v, %v)", root, user)
	defer glog.V(4).Infof("Leaving mkCtlFile(%v, %v)", root, user)

	glog.V(3).Infoln("Creating ctl file")
	k := new(ctlFile)
	if err := k.Add(root, "calculate", user, nil, 0666, k); err != nil {
		glog.Errorln("Can't create ctl file: ", err)
		return err
	}

	return nil
}

func mkDataFile(root *srv.File, user p.User) error {
	glog.V(4).Infof("Entering mkDataFile(%v, %v)", root, user)
	defer glog.V(4).Infof("Leaving mkDataFile(%v, %v)", root, user)

	glog.V(3).Infoln("Creating data file")
	k := new(dataFile)

	if err := k.Add(root, "data", user, nil, 0555, k); err != nil {
		glog.Errorln("Can't create data file: ", err)
		return err
	}

	op.d = k
	return nil
}

func (c *ctlFile) Write(fid *srv.FFid, data []byte, offset uint64) (int, error) {
	glog.V(4).Infof("Entering ctlFile.Write(%v, %v, %v)", fid, data, offset)
	defer glog.V(4).Info("Leaving ctlFile.Write(%v, %v, %v)", fid, data, offset)

	c.Lock()
	defer c.Unlock()

	glog.V(3).Infof("Compute the expression: %s", string(data))
	// TODO: Compute the expression and handle error

	op.result = string(data)

	return len(data), nil
}

func (d *dataFile) Read(fid *srv.FFid, buf []byte, offset uint64) (int, error) {
	glog.V(4).Infof("Entering into dataFile.Read(%v, %v, %v)", fid, buf, offset)
	defer glog.V(4).Infof("Leaving dataFile.Read(%v, %v, %v)", fid, buf, offset)

	if offset > uint64(len(op.result.(string))) {
		return 0, nil
	}

	copy(buf, op.result.(string))

	return len(op.result.(string)), nil
}

func (c *ctlFile) Wstat(fid *srv.FFid, dir *p.Dir) error {
	return nil
}
