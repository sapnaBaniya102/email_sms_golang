package gsm7bit

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var Packed = gsm7Encoding{
	encoder: &gsm7Encoder{packed: true},
	decoder: &gsm7Decoder{packed: true},
}
var UnPacked = gsm7Encoding{
	encoder: &gsm7Encoder{packed: false},
	decoder: &gsm7Decoder{packed: false},
}

type gsm7Encoding struct{ encoder, decoder transform.Transformer }

func (e gsm7Encoding) NewDecoder() *encoding.Decoder {
	return &encoding.Decoder{Transformer: e.decoder}
}

func (e gsm7Encoding) NewEncoder() *encoding.Encoder {
	return &encoding.Encoder{Transformer: e.encoder}
}
