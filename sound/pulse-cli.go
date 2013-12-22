// +build ignore

package pulse

import (
	"bufio"
	//"io"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

type Pulse struct {
	Path   string
	Uid    int
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	eol    string
}

func (p *Pulse) GetRunTimePath() string {
	p.Uid = os.Getuid()
	p.Path = "/run/user/" + strconv.Itoa(p.Uid) + "/pulse/cli"
	return p.Path
}

func PulseCliShell(serverAddr string) {
	path := serverAddr
	if path == "" {
		path = "/run/user/1000/pulse/cli" //default socket path
	}

	conn, err := net.Dial("unix", path)

	if err != nil {
		panic(err)
	}
	lineRead := make([]byte, 65536)
	n, err := conn.Read(lineRead)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf(string(lineRead[0 : n-2]))

	i := 0
	for {
		bio := bufio.NewReader(os.Stdin)
		lineWrite, _, err := bio.ReadLine()

		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}
		}

		i = i + 1

		_, err = conn.Write(lineWrite)
		conn.Write([]byte("\r\n"))

		if err != nil {
			println(err.Error())
			break
		}

		n, err := conn.Read(lineRead)
		//println(n, ": \n")
		if err != nil {
			println(err.Error())
		}
		fmt.Printf(string(lineRead[0 : n-2]))
		time.Sleep(0)
	}
}

func (p *Pulse) PulseInit() int {
	p.GetRunTimePath()
	p.eol = "\r\n"
	return 0
}

func (p *Pulse) PulseConnect() int {
	if p.Path == "" {
		p.Path = "/run/user/1000/pulse/cli"
	}
	conn, err := net.Dial("unix", p.Path)
	p.conn = conn
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	p.reader = bufio.NewReader(p.conn)
	p.writer = bufio.NewWriter(p.conn)
	line, err := p.PulseRead()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(line)
	return 0
}

func (p *Pulse) PulseRead() (string, error) {
	line := make([]byte, 65536)
	n, err := p.conn.Read(line)
	return string(line[0 : n-2]), err
}

func (p *Pulse) PulseWrite(command string) (int, error) {
	n, err := p.conn.Write([]byte(command))
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	n, err = p.conn.Write([]byte(p.eol))
	return n, err
}

func (p *Pulse) PulseCtl(command string) (string, error) {
	_, err := p.PulseWrite(command)
	if err != nil {
		return "", err
	}
	line, err := p.PulseRead()
	return line, err
}
