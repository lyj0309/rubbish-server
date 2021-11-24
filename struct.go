package main

import "time"

type User struct {
	User    string `gorm:"primaryKey"`
	Pwd     string
	Session string `gorm:"index"`
	Type    string
	Phone   string
}

type Recycle struct {
	Id    int `gorm:"primaryKey;autoIncrement"`
	CUser string
	Place string
	RUser string
	Time  time.Time
	Info  string
	RPhone string
}

type Rubbish struct {
	Name string `gorm:"primaryKey;index"`
	Type string
}

type Garbage struct {
	ID      int
	Name    string
	Fname   string
	Content string
}
