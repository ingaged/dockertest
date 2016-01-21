package dockertest

import (
	"fmt"
	"time"

	"github.com/pborman/uuid"
)

type SFTPCredentials struct {
	Username string
	Password string
	UID      string
}

type SFTPConfig struct {
	Credentials      SFTPCredentials
	SourceVolumePath string
	DestPath         string
}

func SetupSFTPContainer(config SFTPConfig) (c ContainerID, ip string, port int, err error) {
	port = randInt(1024, 49150)
	forward := fmt.Sprintf("%d:%d", port, 22)
	if BindDockerToLocalhost != "" {
		forward = "127.0.0.1:" + forward
	}

	volume := ""

	if config.SourceVolumePath != "" && config.DestPath != "" {
		volume = fmt.Sprintf("-v %s:/home/%s/%s", config.SourceVolumePath, config.Credentials.Username, config.DestPath)
	}

	credentials := fmt.Sprintf("%s:%s:%s", config.Credentials.Username, config.Credentials.Password, config.Credentials.UID)

	c, ip, err = setupContainer("atmoz/sftp", port, 10*time.Second, func() (string, error) {
		if volume != "" {
			return run("--name", uuid.New(), volume, "-p", forward, "-d", "atmoz/sftp", credentials)
		}

		return run("--name", uuid.New(), "-p", forward, "-d", "atmoz/sftp", credentials)
	})
	return
}

func SetupSSHContainer(password string) (c ContainerID, ip string, port int, err error) {
	port = randInt(1024, 49150)
	forward := fmt.Sprintf("%d:%d", port, 22)
	if BindDockerToLocalhost != "" {
		forward = "127.0.0.1:" + forward
	}

	credentials := fmt.Sprintf(`--env="ROOT_PASS=%s"`, password)

	c, ip, err = setupContainer("million12/ssh", port, 10*time.Second, func() (string, error) {
		return run("--name", uuid.New(), "-p", forward, credentials, "-d", "million12/ssh")
	})

	return
}
