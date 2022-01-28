package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV5NoAuthWithBigPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 1, t)
}

func TestV5NoAuthWithMiddlePictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 4, t)
}

func TestV5NoAuthWithSmallPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 7, t)
}

func TestV5NoAuthWithBigPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 2, t)
}

func TestV5NoAuthWithMiddlePictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 5, t)
}

func TestV5NoAuthWithSmallPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 8, t)
}

func TestV5NoAuthWithBigPictureAndIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 3, t)
}

func TestV5NoAuthWithMiddlePictureAndIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 6, t)
}

func TestV5NoAuthWithSmallPictureAndIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 9, t)
}
func TestV5PasswordAuthWithBigPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 10, t)
}

func TestV5PasswordAuthWithMiddlePictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 13, t)
}

func TestV5PasswordAuthWithSmallPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "auth", 16, t)
}

func TestV5PasswordAuthWithBigPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 10, t)
}

func TestV5PasswordAuthWithMiddlePictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 13, t)
}

func TestV5PasswordAuthWithSmallPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 16, t)
}

func TestV5PasswordAuthWithBigPictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 12, t)
}

func TestV5PasswordAuthWithMiddlePictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 15, t)
}

func TestV5PasswordAuthSmallPictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "auth", 18, t)
}
