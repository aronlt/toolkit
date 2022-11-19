package ttypes

import "errors"

var ErrorTimeout = errors.New("time out error")
var ErrorFullChan = errors.New("full chan error")
var ErrorInvalidParameter = errors.New("invalid parameter")
