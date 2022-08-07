package utils

import (
	"NFT-BASE-BACK/base"
	"regexp"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func Check(minLength, maxLength, minLevel int, pwd string) base.ErrCode {
	if len(pwd) < minLength {
		return base.PasswdLengthError
	}
	if len(pwd) > maxLength {
		return base.PasswdLengthError
	}

	var level int = levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	if level < minLevel {
		return base.PasswdFormatError
	}
	return base.Success
}
