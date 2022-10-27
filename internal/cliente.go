package WhatCook

const (
	masculino = iota + 1
	femenino
	otro
)

type Cliente struct {
	nombre    string
	apellidos string
	correo    string
	genero    int
}
