package scrapper_test

import (
	"github.com/vds/amazon_scrapper/pkg/scrapper"
	"testing"
)

const link=`https://www.amazon.in/gp/product/B07HGMLBW1/ref=s9_acss_bw_cg_Topbann_2a1_w?pf_rd_m=A1K21FY43GMZF8&pf_rd_s=merchandised-search-4&pf_rd_r=81HES4P01Y18V9XEE2P6&pf_rd_t=101&pf_rd_p=6c0e5a1a-a9c2-441c-968c-513e0354b7a3&pf_rd_i=16613114031`

func TestScrapper(t *testing.T){
	title,company,price,status:=scrapper.ScrapeLink(link)
	assertTitle(t,title,`OnePlus 7 (Mirror Blue, 6GB RAM, 128GB Storage)`)
	assertCompany(t,company,`OnePlus`)
	assertPrice(t,price,32999.00)
	assertStatus(t,status,0)
}

func assertTitle(t *testing.T,got string,want string){
	t.Helper()
	if got!=want{
		t.Errorf("Got title %s want %s",got,want)
	}
}
func assertCompany(t *testing.T,got string,want string){
	t.Helper()
	if got!=want{
		t.Errorf("Got company %s want %s",got,want)
	}
}
func assertPrice(t *testing.T,got float64,want float64){
	t.Helper()
	if got!=want{
		t.Errorf("Got title %f want %f",got,want)
	}
}
func assertStatus(t *testing.T,got int,want int){
	t.Helper()
	if got!=want{
		t.Errorf("Got title %d want %d",got,want)
	}
}