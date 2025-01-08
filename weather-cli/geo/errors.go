package geo

import "errors"


func ErrorNot200Status(statusCode int) error {
	return errors.New("Not 200 status while GET geo by IP. StatusCode == " + string(statusCode))
}
