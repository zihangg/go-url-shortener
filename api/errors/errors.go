package errors

import "errors"

var (
	ErrInvalidUrl = errors.New("invalid url")
	ErrInvalidEncodedString = errors.New("unable to decode symbol")
	ErrShortUrlAlreadyExists = errors.New("short url already exists")
	ErrRedisSetError = errors.New("redis set failed")
	ErrRedisEmptyResult = errors.New("key not found in redis")
	ErrRedisGetError = errors.New("redis get failed")
)