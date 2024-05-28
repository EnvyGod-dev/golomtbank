package utils

import "errors"

var (
	ErrExpiredToken = errors.New("Токений хугацаа дууссан байна")
	ErrInvaildToken = errors.New("Токен хүчингүй байна")
)
