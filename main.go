package main

import (
	"errors"
	"github.com/docker/go-plugins-helpers/volume"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
)

type SharedMountVolumeDriver struct {
	localMountedPath string
}

func (m SharedMountVolumeDriver) Create(request *volume.CreateRequest) error {
	return os.Mkdir("/opt/sharevol/"+request.Name, 750)
}

func (m SharedMountVolumeDriver) Remove(request *volume.RemoveRequest) error {
	return os.RemoveAll("/opt/sharevol/" + request.Name)
}

func (m SharedMountVolumeDriver) Get(request *volume.GetRequest) (*volume.GetResponse, error) {

	files, err := ioutil.ReadDir("/opt/sharevol/")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() && f.Name() == request.Name {
			getResp := new(volume.GetResponse)

			getResp.Volume = new(volume.Volume)

			getResp.Volume.Name = f.Name()
			getResp.Volume.Mountpoint = "/opt/sharevol/" + f.Name()

			return getResp, nil
		}
	}

	return nil, errors.New("volume not found")
}

func (m SharedMountVolumeDriver) List() (*volume.ListResponse, error) {
	log.Print("listing volumes")

	listResp := new(volume.ListResponse)

	files, err := ioutil.ReadDir("/opt/sharevol/")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			vol := new(volume.Volume)

			log.Print("found " + f.Name())

			vol.Name = f.Name()
			vol.Mountpoint = "/opt/sharevol/" + f.Name()

			listResp.Volumes = append(listResp.Volumes, vol)
		}
	}

	return listResp, nil
}

func (m SharedMountVolumeDriver) Path(request *volume.PathRequest) (*volume.PathResponse, error) {
	pathResponse := new(volume.PathResponse)

	pathResponse.Mountpoint = "/opt/sharevol/" + request.Name

	return pathResponse, nil
}

func (m SharedMountVolumeDriver) Mount(request *volume.MountRequest) (*volume.MountResponse, error) {
	lockFile := "/opt/sharevol/" + request.Name + ".mounted"

	if _, err := os.Stat(lockFile); err == nil {
		return nil, errors.New("already mounted")
	}

	if _, err := os.Create(lockFile); err != nil {
		return nil, errors.New("can't lock volume")
	}

	mountResp := new(volume.MountResponse)

	mountResp.Mountpoint = "/opt/sharevol/" + request.Name

	return mountResp, nil
}

func (m SharedMountVolumeDriver) Unmount(request *volume.UnmountRequest) error {
	lockFile := "/opt/sharevol/" + request.Name + ".mounted"

	if _, err := os.Stat(lockFile); os.IsNotExist(err) {
		return errors.New("not mounted")
	}

	return os.Remove(lockFile)
}

func (m SharedMountVolumeDriver) Capabilities() *volume.CapabilitiesResponse {
	capResp := new(volume.CapabilitiesResponse)
	capResp.Capabilities.Scope = "global"

	return capResp
}

func main() {
	d := SharedMountVolumeDriver{}
	h := volume.NewHandler(d)
	u, _ := user.Lookup("root")
	gid, _ := strconv.Atoi(u.Gid)

	println("start")

	_ = h.ServeUnix("sharevol", gid)
}
