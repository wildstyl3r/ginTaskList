package main

type Config struct {
	Database struct {
		Username string
		Password string
		Host     string
		Port     string
		DBname   string
	}
}
