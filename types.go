package main

import (
//	"fmt"
)

type ApiKeyDescriptor struct {
	UUID		string
	Name		string
	Description	string
	AccessKey	string
	SecretKey	string
	Host		string
	GenAccessKey	string
	GenSecretKey	string
	ProjectId	string
}

type LocalAuthConfigDescriptor struct {
	UUID		string
	Host		string
	AccessKey	string
	SecretKey	string
	ProjectId	string
	Realname	string
	Username	string
	Password	string
	Enabled		bool
	Accessmode	string
}
