package signature

type Signature struct {
	Id          int64
	Sign        []byte
	Sha         string
	OffsetBegin int64 // смещение в байтах от начала
	OffsetEnd   int64
	Dtype       string
}
