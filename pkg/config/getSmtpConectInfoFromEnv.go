package config

import (
	"os"
	"strconv"
)

type SmtpConnectInfo struct {
	Server   string
	Port     int
	User     string
	Password string
}

// GetSmtpConnectInfoFromEnv returns an SMTP connection info struct based on the values of environment variables :
//	SMTP_SERVER : string containing the SMTP server name
//	SMTP_PORT : int value between 1 and 65535 containing the smtp port to use
//	SMTP_USER : string containing the smtp username
//	SMTP_PASSWORD : string containing the smtp user password
func GetSmtpConnectInfoFromEnv(defaultServer string, defaultPort int, defaultUser string, defaultPassword string) (*SmtpConnectInfo, error) {
	server := defaultServer
	srvPort := defaultPort
	user := defaultUser
	password := defaultPassword
	var mySmtpConnection SmtpConnectInfo

	var err error
	val, exist := os.LookupEnv("SMTP_PORT")
	if exist {
		srvPort, err = strconv.Atoi(val)
		if err != nil {
			return &mySmtpConnection, &ErrorConfig{
				err: err,
				msg: "ERROR: CONFIG ENV SMTP_PORT should contain a valid integer.",
			}
		}
		if srvPort < 1 || srvPort > 65535 {
			return &mySmtpConnection, &ErrorConfig{
				err: err,
				msg: "ERROR: CONFIG ENV SMTP_PORT should contain an integer between 1 and 65535",
			}
		}
	}
	val, exist = os.LookupEnv("SMTP_SERVER")
	if exist {
		server = val
	}
	val, exist = os.LookupEnv("SMTP_USER")
	if exist {
		user = val
	}
	val, exist = os.LookupEnv("SMTP_PASSWORD")
	if exist {
		password = val
	}
	mySmtpConnection = SmtpConnectInfo{
		Server:   server,
		Port:     srvPort,
		User:     user,
		Password: password,
	}
	return &mySmtpConnection, nil
}
