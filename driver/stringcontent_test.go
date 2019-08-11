package digestfs_driver_test

import (
	"github.com/reiver/go-digestfs/driver"

	"unicode/utf8"

	"testing"
)

func TestStringContentAsContent(t *testing.T) {

	var content digestfs_driver.Content = digestfs_driver.StringContent("Hello world!") // THIS IS WHAT ACTUALLY MATTERS!

	if nil == content {
		t.Error("This should never happen.")
	}
}

func TestStringContentClose(t *testing.T) {

	tests := []struct{
		Content  string
	}{
		{
			Content: "",
		},



		{
			Content: "A",
		},
		{
			Content: "Ab",
		},
		{
			Content: "AbC",
		},
		{
			Content: "AbCd",
		},
		{
			Content: "AbCdE",
		},
		{
			Content: "AbCdEf",
		},
		{
			Content: "AbCdEfG",
		},
		{
			Content: "AbCdEfGh",
		},
		{
			Content: "AbCdEfGhI",
		},
		{
			Content: "AbCdEfGhIj",
		},
		{
			Content: "AbCdEfGhIjK",
		},
		{
			Content: "AbCdEfGhIjKl",
		},
		{
			Content: "AbCdEfGhIjKlM",
		},
		{
			Content: "AbCdEfGhIjKlMn",
		},
		{
			Content: "AbCdEfGhIjKlMnO",
		},
		{
			Content: "AbCdEfGhIjKlMnOp",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQ",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQr",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrS",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrSt",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStU",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUv",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvW",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWx",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWxY",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWxYz",
		},



		{
			Content: "apple",
		},
		{
			Content: "BANANA",
		},
		{
			Content: "Cherry",
		},
		{
			Content: "dATE",
		},



		{
			Content: "Hello world!",
		},



		{
			Content: "😏😐👾🤖😈",
		},



		{
			Content: "ا ب پ ت ث ج چ ح خ د ذ ر ز ژ س ش ص ض ط ظ ع غ ف ق ک گ ل م ن و ه ی",
		},
	}

	for testNumber, test := range tests {

		var content digestfs_driver.Content = digestfs_driver.StringContent(test.Content)

		err := content.Close()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Logf("CONTENT: %q", test.Content)
			continue
		}

		if expected, actual := len(test.Content), content.Len(); expected != actual {
			t.Errorf("For test #%d, the actual len was not what was exepcted.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("CONTENT: %q", test.Content)
			continue
		}

		{
			var buffer []byte = make([]byte, len(test.Content))

			n, err := content.ReadAt(buffer, 0)
			if nil == err {
				t.Errorf("For test #%d, expected an error, but did not actually get one: %#v", testNumber, err)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			if expected, actual := 0, n; expected != actual {
				t.Errorf("For test #%d, did not actually get the expected number of bytes read.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
		}
	}
}

func TestStringContentLen(t *testing.T) {

	tests := []struct{
		Content  string
	}{
		{
			Content: "",
		},



		{
			Content: "A",
		},
		{
			Content: "Ab",
		},
		{
			Content: "AbC",
		},
		{
			Content: "AbCd",
		},
		{
			Content: "AbCdE",
		},
		{
			Content: "AbCdEf",
		},
		{
			Content: "AbCdEfG",
		},
		{
			Content: "AbCdEfGh",
		},
		{
			Content: "AbCdEfGhI",
		},
		{
			Content: "AbCdEfGhIj",
		},
		{
			Content: "AbCdEfGhIjK",
		},
		{
			Content: "AbCdEfGhIjKl",
		},
		{
			Content: "AbCdEfGhIjKlM",
		},
		{
			Content: "AbCdEfGhIjKlMn",
		},
		{
			Content: "AbCdEfGhIjKlMnO",
		},
		{
			Content: "AbCdEfGhIjKlMnOp",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQ",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQr",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrS",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrSt",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStU",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUv",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvW",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWx",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWxY",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWxYz",
		},



		{
			Content: "apple",
		},
		{
			Content: "BANANA",
		},
		{
			Content: "Cherry",
		},
		{
			Content: "dATE",
		},



		{
			Content: "Hello world!",
		},



		{
			Content: "😏😐👾🤖😈",
		},



		{
			Content: "ا ب پ ت ث ج چ ح خ د ذ ر ز ژ س ش ص ض ط ظ ع غ ف ق ک گ ل م ن و ه ی",
		},
	}

	for testNumber, test := range tests {

		var content digestfs_driver.Content = digestfs_driver.StringContent(test.Content)

		if expected, actual := len(test.Content), content.Len(); expected != actual {
			t.Errorf("For test #%d, the actual len was not what was exepcted.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("CONTENT: %q", test.Content)
			continue
		}
	}
}

func TestStringContent(t *testing.T) {

	tests := []struct{
		Content string
	}{
		{
			Content: "",
		},



		{
			Content: "A",
		},
		{
			Content: "Ab",
		},
		{
			Content: "AbC",
		},
		{
			Content: "AbCd",
		},
		{
			Content: "AbCdE",
		},
		{
			Content: "AbCdEf",
		},
		{
			Content: "AbCdEfG",
		},
		{
			Content: "AbCdEfGh",
		},
		{
			Content: "AbCdEfGhI",
		},
		{
			Content: "AbCdEfGhIj",
		},
		{
			Content: "AbCdEfGhIjK",
		},
		{
			Content: "AbCdEfGhIjKl",
		},
		{
			Content: "AbCdEfGhIjKlM",
		},
		{
			Content: "AbCdEfGhIjKlMn",
		},
		{
			Content: "AbCdEfGhIjKlMnO",
		},
		{
			Content: "AbCdEfGhIjKlMnOp",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQ",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQr",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrS",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrSt",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStU",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUv",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvW",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWx",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWxY",
		},
		{
			Content: "AbCdEfGhIjKlMnOpQrStUvWxYz",
		},



		{
			Content: "apple",
		},
		{
			Content: "BANANA",
		},
		{
			Content: "Cherry",
		},
		{
			Content: "dATE",
		},



		{
			Content: "Hello world!",
		},



		{
			Content: "😏😐👾🤖😈",
		},



		{
			Content: "ا ب پ ت ث ج چ ح خ د ذ ر ز ژ س ش ص ض ط ظ ع غ ف ق ک گ ل م ن و ه ی",
		},
	}

	for testNumber, test := range tests {

		var content digestfs_driver.Content = digestfs_driver.StringContent(test.Content)
		if nil == content {
			t.Errorf("For test #%d, expected non-nil content, but actually got: %#v", testNumber, content)
			continue
		}

		{
			var buffer []byte = make([]byte, len(test.Content))

			n, err := content.ReadAt(buffer, 0)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			if expected, actual := len(test.Content), n; expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected for number of byte read.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("CONTENT: %q", test.Content)
				t.Logf("BUFFER:  %q", buffer)
				continue
			}

			if expected, actual := test.Content, string(buffer); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected for what was read.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
			}

			buffer = nil
		}

		if 4 <= utf8.RuneCountInString(test.Content) {
			r0, size0 := utf8.DecodeRuneInString(test.Content)
			if utf8.RuneError == r0 {
				t.Errorf("For test #%d, did not expect a rune error, but actually got one,", testNumber)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}


			var length int
			var errored bool
			func(){
				s := test.Content[size0:]
				for i:=0; i<3; i++ {
					r, size := utf8.DecodeRuneInString(test.Content)
					if utf8.RuneError == r {
						t.Errorf("For test #%d, did not expect a rune error, but actually got one,", testNumber)
						t.Logf("CONTENT: %q", test.Content)
						errored = true
						return
					}
					s = s[size:]

					length += size
				}
			}()
			if errored {
				t.Errorf("For test #%d, something went wrong when was trying to calculated the buffer length.", testNumber)
				continue
			}

			var buffer []byte = make([]byte, length)


			n, err := content.ReadAt(buffer[:], int64(size0))
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			if expected, actual := length, n; expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected for number of byte read.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("CONTENT: %q", test.Content)
				t.Logf("BUFFER:  %q", buffer)
				continue
			}

			if expected, actual := test.Content[size0:size0+length], string(buffer); expected != actual {
				t.Errorf("For test #%d, did not actually get what was expected for what was read.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
			}

			buffer = nil
		}
	}
}
