package network

import ()

type CodecInterface interface {
	Encode(v interface{}) (err error)
	Decode() (msg interface{}, err error)
	Close() error
}
