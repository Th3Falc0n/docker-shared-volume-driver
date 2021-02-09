package shared_volume_plugin

import (
	"github.com/docker/go-plugins-helpers/sdk"
	"github.com/docker/go-plugins-helpers/volume"
	//"os/user"
	//"strconv"
)

type SharedMountVolumeDriver struct {
	localMountedPath string
}

func (m SharedMountVolumeDriver) Create(request *volume.CreateRequest) error {
	panic("implement me")
}

func (m SharedMountVolumeDriver) List() (*volume.ListResponse, error) {
	listResp := new (volume.ListResponse)

	return listResp, nil
}

func (m SharedMountVolumeDriver) Get(request *volume.GetRequest) (*volume.GetResponse, error) {
	panic("implement me")
}

func (m SharedMountVolumeDriver) Remove(request *volume.RemoveRequest) error {
	panic("implement me")
}

func (m SharedMountVolumeDriver) Path(request *volume.PathRequest) (*volume.PathResponse, error) {
	panic("implement me")
}

func (m SharedMountVolumeDriver) Mount(request *volume.MountRequest) (*volume.MountResponse, error) {
	panic("implement me")
}

func (m SharedMountVolumeDriver) Unmount(request *volume.UnmountRequest) error {
	panic("implement me")
}

func (m SharedMountVolumeDriver) Capabilities() *volume.CapabilitiesResponse {
	capResp := new(volume.CapabilitiesResponse)
	capResp.Capabilities.Scope = "global"

	return capResp
}

func main() {
	d := SharedMountVolumeDriver{}
	h := volume.NewHandler(d)
	//u, _ := user.Lookup("root")
	//gid, _ := strconv.Atoi(u.Gid)
	//_ = h.ServeUnix("sharevol", gid)
	_ = h.ServeTCP("test", "0.0.0.0", sdk.WindowsDefaultDaemonRootDir(), nil)
}
