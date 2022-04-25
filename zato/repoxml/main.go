package main

import (
	"encoding/xml"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/xmltojson", Zatoresponse).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":20000", router))
}

func Zatoresponse(w http.ResponseWriter, r *http.Request) {
	var response = []byte(`
<repomd xmlns="http://linux.duke.edu/metadata/repo" xmlns:rpm="http://linux.duke.edu/metadata/rpm">
<revision>1611239850</revision>
<data type="filelists">
<checksum type="sha256">8425f30e1f1f821f8264764b2a6996c6f7f8e599d41c1d65654742d8176feaeb</checksum>
<open-checksum type="sha256">7784e16bddbf4cc990f4b78d4f6f7de68b5d931926f05acd0252f47d4f323fb5</open-checksum>
<location href="repodata/8425f30e1f1f821f8264764b2a6996c6f7f8e599d41c1d65654742d8176feaeb-filelists.xml.gz"/>
<timestamp>1611239852</timestamp>
<size>115043</size>
<open-size>2298912</open-size>
</data>
<data type="primary">
<checksum type="sha256">29cab02afe78bd95895c4ce7a3aa9e6f1b578577c8f063859e10a3de69b8c042</checksum>
<open-checksum type="sha256">57f23f2305a8433f25683a494732d9c3801287a1c4bf2837d839bd4e044d3351</open-checksum>
<location href="repodata/29cab02afe78bd95895c4ce7a3aa9e6f1b578577c8f063859e10a3de69b8c042-primary.xml.gz"/>
<timestamp>1611239852</timestamp>
<size>1710</size>
<open-size>8860</open-size>
</data>
<data type="primary_db">
<checksum type="sha256">9ea9dbe00dcd20909733213f364d1876b401cfa8178fe59030dc029aefd97d8b</checksum>
<open-checksum type="sha256">32f92f346cfb55d6eb4321f0f7d02db40f44c5e6a1b525560b0d6033aae11a95</open-checksum>
<location href="repodata/9ea9dbe00dcd20909733213f364d1876b401cfa8178fe59030dc029aefd97d8b-primary.sqlite.bz2"/>
<timestamp>1611239855</timestamp>
<database_version>10</database_version>
<size>5507</size>
<open-size>90112</open-size>
</data>
<data type="other_db">
<checksum type="sha256">96d11778f68f5a969e81ff226bcf9d594cba7eecfbec8958dfeab3ad4e501bcb</checksum>
<open-checksum type="sha256">3b58ba2ebceada49a0e080e426ca2a6b507d01ea5d83f232bbea55482785eca2</open-checksum>
<location href="repodata/96d11778f68f5a969e81ff226bcf9d594cba7eecfbec8958dfeab3ad4e501bcb-other.sqlite.bz2"/>
<timestamp>1611239853</timestamp>
<database_version>10</database_version>
<size>764</size>
<open-size>24576</open-size>
</data>
<data type="other">
<checksum type="sha256">75c98da0b1257c20548867ee07a51fb1b42e995d68d6f18149f0eceb095fd019</checksum>
<open-checksum type="sha256">f4f5cb6a817fde141c42839fd22fd9c17eecb7127f5de2a9eda03ee922aecc7a</open-checksum>
<location href="repodata/75c98da0b1257c20548867ee07a51fb1b42e995d68d6f18149f0eceb095fd019-other.xml.gz"/>
<timestamp>1611239852</timestamp>
<size>330</size>
<open-size>454</open-size>
</data>
<data type="filelists_db">
<checksum type="sha256">65fdbdc4b4eb39d8a1f1ab1ecae1898a1df04bd06129e59cfbc1d672a4b336dd</checksum>
<open-checksum type="sha256">65c15b06e6dc3f480385fff93c7cbf1cbe72b01046164da17a32adf903d7ab44</open-checksum>
<location href="repodata/65fdbdc4b4eb39d8a1f1ab1ecae1898a1df04bd06129e59cfbc1d672a4b336dd-filelists.sqlite.bz2"/>
<timestamp>1611239853</timestamp>
<database_version>10</database_version>
<size>154843</size>
<open-size>1204224</open-size>
</data>
</repomd>`)
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(200)
	if err := xml.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

type Response struct {
	Repomd Repomd `json:"repomd"`
}

type Repomd struct {
	Revision string  `json:"revision"`
	Data     []Datum `json:"data"`
	Xmlns    string  `json:"_xmlns"`
	XmlnsRPM string  `json:"_xmlns:rpm"`
}

type Datum struct {
	Checksum        Checksum `json:"checksum"`
	OpenChecksum    Checksum `json:"open-checksum"`
	Location        Location `json:"location"`
	Timestamp       string   `json:"timestamp"`
	Size            string   `json:"size"`
	OpenSize        string   `json:"open-size"`
	Type            string   `json:"_type"`
	DatabaseVersion *string  `json:"database_version,omitempty"`
}

type Checksum struct {
	Type Type   `json:"_type"`
	Text string `json:"__text"`
}

type Location struct {
	Href string `json:"_href"`
}

type Type string

const (
	Sha256 Type = "sha256"
)
