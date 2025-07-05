package main

var Configuration *Config;

type Config struct {
	KosmostServerAddress 		string		`yaml:"kosmos_server_port"`
	UserServiceTarget 			string		`yaml:"user_service_target"`
	LanguageServiceTarget 		string		`yaml:"language_service_target"`
	SecurityServiceTarget 		string		`yaml:"security_service_target"`
	CourseServiceTarget 		string		`yaml:"course_service_target"`	
}
