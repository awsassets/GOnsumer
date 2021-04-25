package portchecker

import (
	"net"
)

type (
	PortCheckerService struct {
	}
)

func (pc *PortCheckerService) Check(ip, port string) bool {
	server, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		return false
	}

	server.Close()
	return true
}
