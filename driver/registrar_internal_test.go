package digestfs_driver

import (
	"testing"
)

func TestInternalRegistrarAsRegistrar(t *testing.T) {

	var registrar Registrar = new(internalRegistrar) // THIS IS WHAT ACTUALLY MATTERS!

	if nil == registrar {
		t.Error("This should never happen.")
	}
}

func TestInternalRegistrar(t *testing.T) {

	var names []string = []string{
		"APPLE",
		"banana",
		"Cherry",
		"dATE",
	}

	var registrar internalRegistrar

	for testNumber, name := range names {
		mounter, err := registrar.Fetch(name)
		if nil == err {
			t.Errorf("For test #%d, expected to get an error, but actually did not: %#v", testNumber, err)
			t.Logf("Name: %q", name)
			continue
		}
		if nil != mounter {
			t.Errorf("For test #%d, did not expect a mounter, but actually got one: %#v", testNumber, mounter)
			t.Logf("Name: %q", name)
			continue
		}

		if _, casted := err.(MounterNotFound); !casted {
			t.Errorf("For test #%d, expected ‘not found’ error, but did not actually get it: (%T) %q", testNumber, err, err)
			t.Logf("Name: %q", name)
			continue
		}
	}

	for testNumber, name := range names {
		var mounter Mounter = MounterFunc(func(params ...interface{}) (MountPoint, error){
			return nil, nil
		})

		if err := registrar.Register(mounter, name); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}
	}

	for testNumber, name := range names {
		mounter, err := registrar.Fetch(name)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Logf("Name: %q", name)
			continue
		}
		if nil == mounter {
			t.Errorf("For test #%d, expected a mounter, but did not actually get one: %#v", testNumber, mounter)
			t.Logf("Name: %q", name)
			continue
		}
	}

	{
		var otherNames []string = []string{
			"StrawBerry",
			"Watermelon",
		}

		for testNumber, name := range otherNames {
			mounter, err := registrar.Fetch(name)
			if nil == err {
				t.Errorf("For test #%d, expected to get an error, but actually did not: %#v", testNumber, err)
				t.Logf("Name: %q", name)
				continue
			}
			if nil != mounter {
				t.Errorf("For test #%d, did not expect a mounter, but actually got one: %#v", testNumber, mounter)
				t.Logf("Name: %q", name)
				continue
			}

			if _, casted := err.(MounterNotFound); !casted {
				t.Errorf("For test #%d, expected ‘not found’ error, but did not actually get it: (%T) %q", testNumber, err, err)
				t.Logf("Name: %q", name)
				continue
			}
		}
	}
}
