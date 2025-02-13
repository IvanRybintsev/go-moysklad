package moysklad

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Uint is a helper routine that allocates a new uint64 value
// to store v and returns a pointer to it.
func Uint(v uint64) *uint64 { return &v }

// Float is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Clamp задаёт значение в диапазоне между указанными нижней и верхней границами
func Clamp(val, min, max int) int {
	switch {
	case val < min:
		return min
	case val > max:
		return max
	default:
		return val
	}
}

type FileTypes interface {
	File | Image
}

// NewFileFromFilepath возвращает *File, на основе переданного пути до файла
// и ошибку, если файл не удалось найти
func NewFileFromFilepath(filePath string) (*File, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(filePath)
	content := base64.StdEncoding.EncodeToString(b)
	f := &File{
		Title:    String(fileName),
		Filename: String(fileName),
		Content:  String(content),
	}
	return f, nil
}

// NewImageFromFilepath возвращает *Image, на основе переданного пути до файла
// и ошибку, если файл не удалось найти
func NewImageFromFilepath(filePath string) (*Image, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(filePath)
	content := base64.StdEncoding.EncodeToString(b)
	f := &Image{
		Title:    String(fileName),
		Filename: String(fileName),
		Content:  String(content),
	}
	return f, nil
}

func unmarshalAny[T any](data []byte) (*T, error) {
	v := new(T)
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}

const errWrongMetaTypeMessage = "meta type mismatch! have %s, need %s\n"

var ErrWrongMetaType = func(h, n MetaType) error { return fmt.Errorf(errWrongMetaTypeMessage, h, n) }

type DataMetaTyper interface {
	MetaTyper
	Data() json.RawMessage
}

// unmarshalTo принимает объект, удовлетворяющий интерфейсу DataMetaTyper
// и структурирует в тип T
func unmarshalTo[T MetaTyper, E DataMetaTyper](element E) (*T, error) {
	t := new(T)
	needType := (*t).MetaType()
	haveType := element.MetaType()

	if haveType != needType {
		return t, ErrWrongMetaType(haveType, needType)
	}
	data := element.Data()
	if data == nil {
		return nil, errors.New("data is empty")
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return t, err
	}
	return t, nil
}

type Interval string

func (i Interval) String() string {
	return string(i)
}

const (
	IntervalHour  Interval = "hour"
	IntervalDay   Interval = "day"
	IntervalMonth Interval = "month"
)

// RGBtoUint64 конвертирует код цвета из формата RRGGBB и RGB в uint64
// Example
// RGBtoUint64("#E6E6E6")
// RGBtoUint64("e3e3e3")
// RGBtoUint64("FFF")
func RGBtoUint64(str string) (uint64, error) {
	var errInvalid = errors.New("invalid format")
	var err error

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}

		err = errInvalid

		return 0
	}

	n := strings.Replace(str, "0x", "", -1)
	n = strings.Replace(n, "0X", "", -1)
	n = strings.Replace(n, "#", "", -1)

	c := color.RGBA{A: 0xff}

	switch len(n) {
	case 6: // RRGGBB
		c.R = hexToByte(n[0])<<4 + hexToByte(n[1])
		c.G = hexToByte(n[2])<<4 + hexToByte(n[3])
		c.B = hexToByte(n[4])<<4 + hexToByte(n[5])
	case 3: // RGB
		c.R = hexToByte(n[0]) * 17
		c.G = hexToByte(n[1]) * 17
		c.B = hexToByte(n[2]) * 17
	default: // invalid format
		err = errInvalid
	}

	if err != nil {
		return 0, err
	}

	hex := fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
	output, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return 0, err
	}
	return output, nil
}

var reContentDisposition = regexp.MustCompile(`filename="(.*)"`)

func GetFileFromResponse(resp *Response) (*PrintFile, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}

	var fileName string
	headerStr := resp.Header.Get(headerContentDisposition)
	if match := reContentDisposition.FindStringSubmatch(headerStr); len(match) > 1 {
		fileName = match[1]
	}
	file := &PrintFile{buf, fileName}
	return file, nil
}

// Deref разыменовывает указатель
func Deref[T any](ptr *T) T {
	if ptr == nil {
		var v T
		return v
	}
	return *ptr
}
