package gocodec

import (
	"unsafe"
)

type stringCodec struct {
	BaseCodec
}

func (codec *stringCodec) Encode(stream *Stream) {
	pStr := unsafe.Pointer(&stream.buf[stream.cursor])
	str := *(*string)(pStr)
	offset := uintptr(len(stream.buf)) - stream.cursor
	header := (*stringHeader)(pStr)
	header.Data = offset
	stream.buf = append(stream.buf, str...)
}

func (codec *stringCodec) Decode(iter *Iterator) {
	pStr := unsafe.Pointer(&iter.cursor[0])
	header := (*stringHeader)(pStr)
	offset := header.Data
	header.Data = uintptr(unsafe.Pointer(&iter.cursor[offset]))
}
