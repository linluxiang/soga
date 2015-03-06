package network

import (
	"io"
)

type CodecInterface interface {
	Encode(v interface{}) (err error)
	Decode() (msg interface{}, err error)
	Close() error
}

type Codec struct {
	encoder *Encoder
	decoder *Decoder
}

type EncoderInterface interface {
	Encode(v interface{}) (err error)
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v interface{}) (err error) {
	return
}

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) Decode() (msg interface{}, err error) {
	return
}
