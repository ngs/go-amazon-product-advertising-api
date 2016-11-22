package amazon

import "testing"

func TestRegionEndpoint(t *testing.T) {
	for _, test := range []Test{
		{"webservices.amazon.com.br", RegionBrazil.Endpoint()},
		{"webservices.amazon.ca", RegionCanada.Endpoint()},
		{"webservices.amazon.cn", RegionChina.Endpoint()},
		{"webservices.amazon.de", RegionGermany.Endpoint()},
		{"webservices.amazon.es", RegionSpain.Endpoint()},
		{"webservices.amazon.fr", RegionFrance.Endpoint()},
		{"webservices.amazon.in", RegionIndia.Endpoint()},
		{"webservices.amazon.it", RegionItaly.Endpoint()},
		{"webservices.amazon.co.jp", RegionJapan.Endpoint()},
		{"webservices.amazon.com.mx", RegionMexico.Endpoint()},
		{"webservices.amazon.co.uk", RegionUK.Endpoint()},
		{"webservices.amazon.com", RegionUS.Endpoint()},
	} {
		test.Compare(t)
	}
}

func TestRegionHTTPSEndpoint(t *testing.T) {
	for _, test := range []Test{
		{"https://webservices.amazon.com.br/onca/xml", RegionBrazil.HTTPSEndpoint()},
		{"https://webservices.amazon.ca/onca/xml", RegionCanada.HTTPSEndpoint()},
		{"https://webservices.amazon.cn/onca/xml", RegionChina.HTTPSEndpoint()},
		{"https://webservices.amazon.de/onca/xml", RegionGermany.HTTPSEndpoint()},
		{"https://webservices.amazon.es/onca/xml", RegionSpain.HTTPSEndpoint()},
		{"https://webservices.amazon.fr/onca/xml", RegionFrance.HTTPSEndpoint()},
		{"https://webservices.amazon.in/onca/xml", RegionIndia.HTTPSEndpoint()},
		{"https://webservices.amazon.it/onca/xml", RegionItaly.HTTPSEndpoint()},
		{"https://webservices.amazon.co.jp/onca/xml", RegionJapan.HTTPSEndpoint()},
		{"https://webservices.amazon.com.mx/onca/xml", RegionMexico.HTTPSEndpoint()},
		{"https://webservices.amazon.co.uk/onca/xml", RegionUK.HTTPSEndpoint()},
		{"https://webservices.amazon.com/onca/xml", RegionUS.HTTPSEndpoint()},
		{"", Region("foo").HTTPSEndpoint()},
	} {
		test.Compare(t)
	}
}

func TestRegionHTTPEndpoint(t *testing.T) {
	for _, test := range []Test{
		{"http://webservices.amazon.com.br/onca/xml", RegionBrazil.HTTPEndpoint()},
		{"http://webservices.amazon.ca/onca/xml", RegionCanada.HTTPEndpoint()},
		{"http://webservices.amazon.cn/onca/xml", RegionChina.HTTPEndpoint()},
		{"http://webservices.amazon.de/onca/xml", RegionGermany.HTTPEndpoint()},
		{"http://webservices.amazon.es/onca/xml", RegionSpain.HTTPEndpoint()},
		{"http://webservices.amazon.fr/onca/xml", RegionFrance.HTTPEndpoint()},
		{"http://webservices.amazon.in/onca/xml", RegionIndia.HTTPEndpoint()},
		{"http://webservices.amazon.it/onca/xml", RegionItaly.HTTPEndpoint()},
		{"http://webservices.amazon.co.jp/onca/xml", RegionJapan.HTTPEndpoint()},
		{"http://webservices.amazon.com.mx/onca/xml", RegionMexico.HTTPEndpoint()},
		{"http://webservices.amazon.co.uk/onca/xml", RegionUK.HTTPEndpoint()},
		{"http://webservices.amazon.com/onca/xml", RegionUS.HTTPEndpoint()},
		{"", Region("foo").HTTPEndpoint()},
	} {
		test.Compare(t)
	}
}
