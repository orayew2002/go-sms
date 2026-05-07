package utils

import "github.com/fiorix/go-smpp/smpp/pdu/pdutext"

func GetTextCodec(textType string, text string) pdutext.Codec {
	switch textType {

	case "GSM7":
		return pdutext.GSM7(text)

	case "GSM7Packed":
		return pdutext.GSM7Packed(text)

	case "ISO88595":
		return pdutext.ISO88595(text)

	case "Latin1":
		return pdutext.Latin1(text)

	case "UCS2":
		return pdutext.UCS2(text)

	default:
		return pdutext.Raw(text)
	}
}

