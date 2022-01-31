package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV5NoAuthBigPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 0, t)
}

func TestV5NoAuthMiddlePictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 1, t)
}

func TestV5NoAuthSmallPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 2, t)
}

func TestV5NoAuthBigPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 3, t)
}

func TestV5NoAuthMiddlePictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 4, t)
}

func TestV5NoAuthSmallPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 5, t)
}

func TestV5NoAuthBigPictureAndIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 6, t)
}

func TestV5NoAuthMiddlePictureAndIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 7, t)
}

func TestV5NoAuthSmallPictureAndIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 8, t)
}
func TestV5PasswordAuthBigPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 9, t)
}

func TestV5PasswordAuthMiddlePictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 10, t)
}

func TestV5PasswordAuthSmallPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 11, t)
}

func TestV5PasswordAuthBigPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 12, t)
}

func TestV5PasswordAuthMiddlePictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 13, t)
}

func TestV5PasswordAuthSmallPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 14, t)
}

func TestV5PasswordAuthBigPictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 15, t)
}

func TestV5PasswordAuthMiddlePictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 16, t)
}

func TestV5PasswordAuthSmallPictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 17, t)
}
