package avs

//здесь все, что касается антивируса, сигнатур и тд

type AVS struct {
	signatures map[string]Signature
	VirusStats map[string][]string
}

type Signature struct {
	id          int64
	sign        []byte
	sha         string
	offsetBegin string
	offsetEnd   string
	dtype       string
}
