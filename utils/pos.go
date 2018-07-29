package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tarm/serial"
)

type POS struct {
	// destination
	serialPort *serial.Port

	// font metrics
	width, height uint8

	// state toggles ESC[char]
	underline  uint8
	emphasize  uint8
	upsidedown uint8
	rotate     uint8

	// state toggles GS[char]
	reverse, smooth uint8

	data []byte
}

func (pos *POS) textReplace(data string) string {
	// text replacement map
	var posTextReplaceMap = map[string]string{
		// horizontal tab
		"&#9;":  "\x09",
		"&#x9;": "\x09",

		// linefeed
		"&#10;": "\n",
		"&#xA;": "\n",

		// xml stuff
		"&apos;": "'",
		"&quot;": `"`,
		"&gt;":   ">",
		"&lt;":   "<",

		// ampersand must be last to avoid double decoding
		"&amp;": "&",
	}

	for k, v := range posTextReplaceMap {
		data = strings.Replace(data, k, v, -1)
	}
	return data
}

func (pos *POS) reset() {
	pos.width = 1
	pos.height = 1

	pos.underline = 0
	pos.emphasize = 0
	pos.upsidedown = 0
	pos.rotate = 0

	pos.reverse = 0
	pos.smooth = 0
}

func NewPOS(portName string) (pos *POS) {
	pos = &POS{}
	serialConfig := &serial.Config{Name: portName, Baud: 9600}
	serialPort, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	pos.serialPort = serialPort
	pos.reset()
	return
}

func (pos *POS) WriteRaw(data []byte) {
	pos.data = append(pos.data, data...)
}

func (pos *POS) Write(data string) {
	pos.WriteRaw([]byte(" " + data + " "))
}

func (pos *POS) Init() {
	pos.reset()
	pos.Write("\x1B@")
}

func (pos *POS) Print() {
	pos.Linefeed(2)
	pos.Write("\xFA")
	if pos.serialPort != nil {
		pos.serialPort.Write(pos.data)
	}
}

func (pos *POS) Group() {
	pos.Write("\x1DV")
}

func (pos *POS) Cash() {
	pos.Write("\x1B\x70\x00\x0A\xFF")
}

func (pos *POS) Linefeed(n int) {
	pos.Write(strings.Repeat("\x0A", n))
}

func (pos *POS) FormfeedN(n int) {
	pos.Write(fmt.Sprintf("\x1Bd%c", n))
}

func (pos *POS) Formfeed() {
	pos.FormfeedN(1)
}

func (pos *POS) SetFont(font string) {
	f := 0

	switch font {
	case "A":
		f = 0
	case "B":
		f = 1
	case "C":
		f = 2
	default:
		log.Printf("Invalid font: '%s', defaulting to 'A'", font)
		f = 0
	}

	pos.Write(fmt.Sprintf("\x1BM%c", f))
}

func (pos *POS) SendFontSize() {
	pos.Write(fmt.Sprintf("\x1D!%c", ((pos.width-1)<<4)|(pos.height-1)))
}

func (pos *POS) SetFontSize(width, height uint8) {
	if width > 0 && height > 0 && width <= 8 && height <= 8 {
		pos.width = width
		pos.height = height
		pos.SendFontSize()
	} else {
		log.Printf("Invalid font size passed: %d x %v", width, height)
	}
}

func (pos *POS) SendUnderline() {
	pos.Write(fmt.Sprintf("\x1B-%c", pos.underline))
}

func (pos *POS) SendEmphasize() {
	pos.Write(fmt.Sprintf("\x1BG%c", pos.emphasize))
}

func (pos *POS) SendUpsidedown() {
	pos.Write(fmt.Sprintf("\x1B{%c", pos.upsidedown))
}

func (pos *POS) SendRotate() {
	pos.Write(fmt.Sprintf("\x1BR%c", pos.rotate))
}

func (pos *POS) SendReverse() {
	pos.Write(fmt.Sprintf("\x1DB%c", pos.reverse))
}

func (pos *POS) SendSmooth() {
	pos.Write(fmt.Sprintf("\x1Db%c", pos.smooth))
}

func (pos *POS) SendMoveX(x uint16) {
	pos.Write(string([]byte{0x1b, 0x24, byte(x % 256), byte(x / 256)}))
}

func (pos *POS) SendMoveY(y uint16) {
	pos.Write(string([]byte{0x1d, 0x24, byte(y % 256), byte(y / 256)}))
}

func (pos *POS) SetUnderline(v uint8) {
	pos.underline = v
	pos.SendUnderline()
}

func (pos *POS) SetEmphasize(u uint8) {
	pos.emphasize = u
	pos.SendEmphasize()
}

func (pos *POS) SetUpsidedown(v uint8) {
	pos.upsidedown = v
	pos.SendUpsidedown()
}

func (pos *POS) SetRotate(v uint8) {
	pos.rotate = v
	pos.SendRotate()
}

func (pos *POS) SetReverse(v uint8) {
	pos.reverse = v
	pos.SendReverse()
}

func (pos *POS) SetSmooth(v uint8) {
	pos.smooth = v
	pos.SendSmooth()
}

func (pos *POS) Pulse() {
	// with t=2 -- meaning 2*2msec
	pos.Write("\x1Bp\x02")
}

func (pos *POS) SetAlign(align string) {
	a := 0
	switch align {
	case "left":
		a = 0
	case "center":
		a = 1
	case "right":
		a = 2
	default:
		log.Printf("Invalid alignment: %s", align)
	}
	pos.Write(fmt.Sprintf("\x1Ba%c", a))
}

func (pos *POS) SetLang(lang string) {
	l := 0

	switch lang {
	case "en":
		l = 0
	case "fr":
		l = 1
	case "de":
		l = 2
	case "uk":
		l = 3
	case "da":
		l = 4
	case "sv":
		l = 5
	case "it":
		l = 6
	case "es":
		l = 7
	case "ja":
		l = 8
	case "no":
		l = 9
	default:
		log.Printf("Invalid language: %s", lang)
	}
	pos.Write(fmt.Sprintf("\x1BR%c", l))
}

func (pos *POS) Text(params map[string]string, data string) {

	if align, ok := params["align"]; ok {
		pos.SetAlign(align)
	}

	if lang, ok := params["lang"]; ok {
		pos.SetLang(lang)
	}

	if smooth, ok := params["smooth"]; ok && (smooth == "true" || smooth == "1") {
		pos.SetSmooth(1)
	}

	if em, ok := params["em"]; ok && (em == "true" || em == "1") {
		pos.SetEmphasize(1)
	}

	if ul, ok := params["ul"]; ok && (ul == "true" || ul == "1") {
		pos.SetUnderline(1)
	}

	if reverse, ok := params["reverse"]; ok && (reverse == "true" || reverse == "1") {
		pos.SetReverse(1)
	}

	if rotate, ok := params["rotate"]; ok && (rotate == "true" || rotate == "1") {
		pos.SetRotate(1)
	}

	if font, ok := params["font"]; ok {
		pos.SetFont(strings.ToUpper(font[5:6]))
	}

	if dw, ok := params["dw"]; ok && (dw == "true" || dw == "1") {
		pos.SetFontSize(2, pos.height)
	}

	if dh, ok := params["dh"]; ok && (dh == "true" || dh == "1") {
		pos.SetFontSize(pos.width, 2)
	}

	if width, ok := params["width"]; ok {
		if i, err := strconv.Atoi(width); err == nil {
			pos.SetFontSize(uint8(i), pos.height)
		} else {
			log.Printf("Invalid font width: %s", width)
		}
	}

	if height, ok := params["height"]; ok {
		if i, err := strconv.Atoi(height); err == nil {
			pos.SetFontSize(pos.width, uint8(i))
		} else {
			log.Printf("Invalid font height: %s", height)
		}
	}

	if x, ok := params["x"]; ok {
		if i, err := strconv.Atoi(x); err == nil {
			pos.SendMoveX(uint16(i))
		} else {
			log.Printf("Invalid x param %v", x)
		}
	}

	if y, ok := params["y"]; ok {
		if i, err := strconv.Atoi(y); err == nil {
			pos.SendMoveY(uint16(i))
		} else {
			log.Printf("Invalid y param %v", y)
		}
	}

	data = pos.textReplace(data)
	if len(data) > 0 {
		pos.Write(data)
	}
}

func (pos *POS) Feed(params map[string]string) {
	// handle lines (form feed X lines)
	if l, ok := params["line"]; ok {
		if i, err := strconv.Atoi(l); err == nil {
			pos.FormfeedN(i)
		} else {
			log.Printf("Invalid line number %v", l)
		}
	}

	// handle units (dots)
	if u, ok := params["unit"]; ok {
		if i, err := strconv.Atoi(u); err == nil {
			pos.SendMoveY(uint16(i))
		} else {
			log.Printf("Invalid unit number %v", u)
		}
	}

	// send linefeed
	pos.Linefeed(1)

	// reset variables
	pos.reset()

	// reset printer
	pos.SendEmphasize()
	pos.SendRotate()
	pos.SendSmooth()
	pos.SendReverse()
	pos.SendUnderline()
	pos.SendUpsidedown()
	pos.SendFontSize()
	pos.SendUnderline()
}

func (pos *POS) FeedAndGroup(params map[string]string) {
	if t, ok := params["type"]; ok && t == "feed" {
		pos.Formfeed()
	}

	pos.Group()
}

func (pos *POS) Barcode(barcode string, format int) {
	code := ""
	switch format {
	case 0:
		code = "\x00"
	case 1:
		code = "\x01"
	case 2:
		code = "\x02"
	case 3:
		code = "\x03"
	case 4:
		code = "\x04"
	case 73:
		code = "\x49"
	}

	// reset settings
	pos.reset()

	// set align
	pos.SetAlign("center")

	// write barcode
	if format > 69 {
		pos.Write(fmt.Sprintf("\x1dk"+code+"%v%v", len(barcode), barcode))
	} else if format < 69 {
		pos.Write(fmt.Sprintf("\x1dk"+code+"%v\x00", barcode))
	}
	pos.Write(fmt.Sprintf("%v", barcode))
}

func (pos *POS) gSend(m byte, fn byte, data []byte) {
	l := len(data) + 2

	pos.Write("\x1b(L")
	pos.WriteRaw([]byte{byte(l % 256), byte(l / 256), m, fn})
	pos.WriteRaw(data)
}

func (pos *POS) Image(params map[string]string, data string) {
	// send alignment to printer
	if align, ok := params["align"]; ok {
		pos.SetAlign(align)
	}

	// get width
	wstr, ok := params["width"]
	if !ok {
		log.Printf("No width specified on image")
	}

	// get height
	hstr, ok := params["height"]
	if !ok {
		log.Printf("No height specified on image")
	}

	// convert width
	width, err := strconv.Atoi(wstr)
	if err != nil {
		log.Printf("Invalid image width %s", wstr)
	}

	// convert height
	height, err := strconv.Atoi(hstr)
	if err != nil {
		log.Printf("Invalid image height %s", hstr)
	}

	// decode data frome b64 string
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Printf(err.Error())
	}

	log.Printf("Image len:%d w: %d h: %d\n", len(dec), width, height)

	// $imgHeader = self::dataHeader(array($img -> getWidth(), $img -> getHeight()), true);
	// $tone = '0';
	// $colors = '1';
	// $xm = (($size & self::IMG_DOUBLE_WIDTH) == self::IMG_DOUBLE_WIDTH) ? chr(2) : chr(1);
	// $ym = (($size & self::IMG_DOUBLE_HEIGHT) == self::IMG_DOUBLE_HEIGHT) ? chr(2) : chr(1);
	//
	// $header = $tone . $xm . $ym . $colors . $imgHeader;
	// $this -> graphicsSendData('0', 'p', $header . $img -> toRasterFormat());
	// $this -> graphicsSendData('0', '2');

	header := []byte{
		byte('0'), 0x01, 0x01, byte('1'),
	}

	a := append(header, dec...)

	pos.gSend(byte('0'), byte('p'), a)
	pos.gSend(byte('0'), byte('2'), []byte{})
}

func (pos *POS) WriteNode(name string, params map[string]string, data string) {
	cstr := ""
	if data != "" {
		str := data[:]
		if len(data) > 40 {
			str = fmt.Sprintf("%s ...", data[0:40])
		}
		cstr = fmt.Sprintf(" => '%s'", str)
	}
	log.Printf("Write: %s => %+v%s\n", name, params, cstr)

	switch name {
	case "text":
		pos.Text(params, data)
	case "feed":
		pos.Feed(params)
	case "cut":
		pos.FeedAndGroup(params)
	case "pulse":
		pos.Pulse()
	case "image":
		pos.Image(params, data)
	}
}
