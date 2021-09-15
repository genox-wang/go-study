package proxy

type Seller struct {
	tickets int
}

type Station struct {
	seller *Seller
}

func (s *Station) sell() {
	s.seller.tickets--
}

type ProxyStation struct {
	station *Station
}

func (p *ProxyStation) sell() {
	p.station.sell()
}
