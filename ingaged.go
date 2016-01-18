package dockertest

import (
	"fmt"
	"time"
)

func SetupSFTPContainer() (c ContainerID, ip string, port int, err error) {
	port = randInt(1024, 49150)
	forward := fmt.Sprintf("%d:%d", port, 22)
	if BindDockerToLocalhost != "" {
		forward = "127.0.0.1:" + forward
	}

	c, ip, err = setupContainer("atmoz/sftp", port, 10*time.Second, func() (string, error) {
		res, err := run("--name", "ingaged-sftp-test", "-p", forward, "-d", "atmoz/sftp", "foo:123:1001")
		return res, err
	})
	return
}
