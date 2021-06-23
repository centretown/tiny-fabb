package docs

import "testing"

func TestRead(t *testing.T) {
	docs := make(Docs)
	err := docs.ReadFile("../assets/doc.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range docs {
		t.Log(d.Sprint())
	}

	t.Log(docs.Sprint())

	d, err := docs.Find("#24---homing-feed-mmmin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(d.Sprint())

	d, err = docs.Find("#x---kill-alarm-lock")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(d.Sprint())

	_, err = docs.Find(" #x---kill-alarm-lock")
	if err == nil {
		t.Fatal("Found: ' #x---kill-alarm-lock'")
	}

}
