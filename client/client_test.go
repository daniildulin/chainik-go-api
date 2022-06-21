package client

import "testing"

func TestTrackedCoins(t *testing.T) {
	client := NewClient()
	list, err := client.TrackedCoins()
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(list) == 0 {
		t.Errorf("empty tracked coins list")
	}
}

func TestTrackedCoin(t *testing.T) {
	client := NewClient()
	_, err := client.TrackedCoin("bee")
	if err != nil {
		t.Errorf(err.Error())
	}
}
