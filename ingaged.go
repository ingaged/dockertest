package dockertest

import (
	"fmt"
	"time"
)

func SetupSFTPContainer(sourceVolume, destVolume string, credentials string) (c ContainerID, ip string, port int, err error) {
	port = randInt(1024, 49150)
	forward := fmt.Sprintf("%d:%d", port, 22)
	if BindDockerToLocalhost != "" {
		forward = "127.0.0.1:" + forward
	}

	volume := ""

	if sourceVolume != "" && destVolume != "" {
		volume = fmt.Sprintf("-v %s:%s", sourceVolume, destVolume)
	}

	c, ip, err = setupContainer("atmoz/sftp", port, 10*time.Second, func() (string, error) {
		res, err := run("--name", "ingaged-sftp-test", volume, "-p", forward, "-d", "atmoz/sftp", credentials)
		return res, err
	})
	return
}
