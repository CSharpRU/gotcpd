package gotcpd

type Config struct {
	Server          struct {
				Host string
				Port uint16
			}
	ReplacePatterns struct {
				Login    string
				Password string
			}
	Epp             struct {
				Login    string
				Password string
				Host     string
				Port     uint16
			}
}

var AppConfig Config
