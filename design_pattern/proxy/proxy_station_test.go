package proxy

import (
	"testing"
)

func TestProxyStation(t *testing.T) {
	seller := &Seller{tickets: 10}

	station := &Station{seller: seller}

	proxy := &ProxyStation{
		station: station,
	}
	proxy.sell()

	if seller.tickets != 9 {
		t.Errorf("111")
	}

}
