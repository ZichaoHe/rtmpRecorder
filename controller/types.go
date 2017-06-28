package controller

type baseResp struct {
	Code   int
	Result interface{}
}

type startRecordReq struct {
	Type    string
	SrcUrl  string
	DstPath string
}

type stopRecordReq struct {
	SrcUrl string
}
