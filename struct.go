package main

import "time"

type User struct {
	User    string `gorm:"primaryKey"`
	Pwd     string
	Session string `gorm:"index"`
}

type Recycle struct {
	Id    int `gorm:"primaryKey;autoIncrement"`
	CUser string
	Place string
	RUser string
	Time  time.Time
	Info  string
}

type Rubbish struct {
	Name string `gorm:"primaryKey;index"`
	Type string
}
