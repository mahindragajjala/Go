package common

type LogConfig struct
{
	Path string
	MaxBackups int
	LogDebug bool
	LogInfo bool
}

type Configuration struct 
{
	WebLog LogConfig;
	DBLog  LogConfig;
}

var Config Configuration;

func InitConfig() {
	Config.WebLog.LogDebug = true;
	Config.WebLog.LogInfo = true;
	Config.DBLog.LogDebug = true;
	Config.DBLog.LogInfo = true;	
}
