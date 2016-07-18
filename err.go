//校验IP的合法性
package main

import (
	"errors"
)

/****************************************************************
*
*  FUNCTION : This function call used by  StartHTTP().
*               检查IP是否在白名单中
*  DESCRIPTION :
*          checkSrcIP(ip string) error
*	return :
*          errors.new("invalid src IP")
*  DESIGNER :
*
*****************************************************************/
func checkSrcIP(ip string) error {
	for _, v := range srcIP {
		if v == ip {
			return nil
		}
	}
	return errors.New("invalid src ip")
}

const (
	errSuccess = 1
	errSrcIP   = 2
	errJson    = 3
	errUnkown  = 99
)
