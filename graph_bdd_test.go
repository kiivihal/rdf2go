package rdf2go_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kiivihal/rdf2go"
)

var testJSONLD = `[{"http://purl.org/dc/terms/extent": [{"@value": "Hoogte: 186 mm, diameter: 148-165 mm"}], "http://purl.org/dc/terms/createdEnd": [{"@value": "-0221-01-01T00:00:01"}], "http://purl.org/dc/terms/createdRaw": [{"@value": "-481 t/m -221"}], "@id": "http://data.brabantcloud.nl/resource/document/museum-klok-en-peel/2458", "http://purl.org/dc/elements/1.1/title": [{"@value": "Terracottabel type bo [Periode van de Strijdende Staten]"}], "http://purl.org/dc/terms/spatial": [{"@value": "China, Azie"}], "@type": ["http://www.europeana.eu/schemas/edm/ProvidedCHO"], "http://purl.org/dc/elements/1.1/identifier": [{"@value": "2458"}], "http://purl.org/dc/elements/1.1/description": [{"@value": "Bellen uit terracotta zoals deze waren grafgiften. Zij werden gemaakt ter vervanging van het originele object dat in voorgaande perioden de dode meegegeven werd. Het bekendste voorbeeld van dit gebruik is het terracotta-leger van de eerste keizer van China."}], "http://purl.org/dc/elements/1.1/date": [{"@value": "-481 t/m -221"}], "http://www.europeana.eu/schemas/edm/type": [{"@value": "IMAGE"}], "http://purl.org/dc/terms/medium": [{"@value": "keramiek"}], "http://purl.org/dc/terms/created": [{"@value": "-0481-01-01T00:00:01"}]}, {"@type": ["http://xmlns.com/foaf/0.1/Document"], "http://schemas.delving.eu/narthex/terms/saveTime": [{"@value": "2018-02-12T18:36:30Z"}], "http://schemas.delving.eu/narthex/terms/belongsTo": [{"@id": "http://data.brabantcloud.nl/resource/dataset/museum-klok-en-peel"}], "http://schemas.delving.eu/narthex/terms/synced": [{"@value": false}], "http://schemas.delving.eu/narthex/terms/contentHash": [{"@value": "de8bc9366bacd77ed1d3060f0ba2b73e124c74f0"}], "@id": "http://data.brabantcloud.nl/resource/aggregation/museum-klok-en-peel/2458/about", "http://creativecommons.org/ns#attributionName": [{"@value": "museum-klok-en-peel"}], "http://xmlns.com/foaf/0.1/primaryTopic": [{"@id": "http://data.brabantcloud.nl/resource/aggregation/museum-klok-en-peel/2458"}]}, {"http://schemas.delving.eu/nave/terms/allowSourceDownload": [{"@value": "false"}], "@type": ["http://schemas.delving.eu/nave/terms/DelvingResource"], "http://schemas.delving.eu/nave/terms/allowLinkedOpenData": [{"@value": "true"}], "http://schemas.delving.eu/nave/terms/featured": [{"@value": "false"}], "http://schemas.delving.eu/nave/terms/allowDeepZoom": [{"@value": "true"}], "@id": "_:Nd1aca3dce3c7451ab5ce6d0c0f7a3009", "http://schemas.delving.eu/nave/terms/public": [{"@value": "true"}], "http://schemas.delving.eu/nave/terms/deepZoomUrl": [{"@value": "https://media.delving.org/iip/deepzoom/mnt/tib/tiles/brabantcloud/museum-klok-en-peel/2458-Bel_type_bo_terracotta_China_strijdende_staten_voorkant.tif.dzi"}]}, {"http://www.europeana.eu/schemas/edm/isShownBy": [{"@id": "https://media.delving.org/thumbnail/brabantcloud/museum-klok-en-peel/2458-Bel_type_bo_terracotta_China_strijdende_staten_voorkant/500"}], "@type": ["http://www.openarchives.org/ore/terms/Aggregation", "http://schemas.delving.eu/narthex/terms/Record"], "http://www.europeana.eu/schemas/edm/rights": [{"@id": "http://creativecommons.org/publicdomain/zero/1.0/"}], "http://www.europeana.eu/schemas/edm/object": [{"@id": "https://media.delving.org/thumbnail/brabantcloud/museum-klok-en-peel/2458-Bel_type_bo_terracotta_China_strijdende_staten_voorkant/220"}], "http://www.europeana.eu/schemas/edm/provider": [{"@value": "Erfgoed Brabant"}], "http://www.europeana.eu/schemas/edm/dataProvider": [{"@value": "Museum Klok & Peel"}], "http://www.europeana.eu/schemas/edm/aggregatedCHO": [{"@id": "http://data.brabantcloud.nl/resource/document/museum-klok-en-peel/2458"}], "@id": "http://data.brabantcloud.nl/resource/aggregation/museum-klok-en-peel/2458", "http://www.openarchives.org/ore/terms/aggregates": [{"@id": "_:Nc7d29843d06541eca36bea1cf446e648"}, {"@id": "_:Nd1aca3dce3c7451ab5ce6d0c0f7a3009"}], "http://www.europeana.eu/schemas/edm/isShownAt": [{"@id": "http://data.brabantcloud.nl/resource/aggregation/museum-klok-en-peel/2458"}]}, {"http://schemas.delving.eu/nave/terms/creatorRole": [{"@value": "gieter"}], "@type": ["http://schemas.delving.eu/nave/terms/BrabantCloudResource"], "http://schemas.delving.eu/nave/terms/collection": [{"@value": "Museum Klok & Peel"}], "http://schemas.delving.eu/nave/terms/thumbLarge": [{"@value": "https://media.delving.org/thumbnail/brabantcloud/museum-klok-en-peel/2458-Bel_type_bo_terracotta_China_strijdende_staten_voorkant/500"}], "@id": "_:Nc7d29843d06541eca36bea1cf446e648", "http://schemas.delving.eu/nave/terms/material": [{"@value": "keramiek"}], "http://schemas.delving.eu/nave/terms/collectionPart": [{"@value": "opgravingen"}], "http://schemas.delving.eu/nave/terms/collectionType": [{"@value": "Algemeen"}], "http://schemas.delving.eu/nave/terms/thumbSmall": [{"@value": "https://media.delving.org/thumbnail/brabantcloud/museum-klok-en-peel/2458-Bel_type_bo_terracotta_China_strijdende_staten_voorkant/220"}], "http://schemas.delving.eu/nave/terms/dimension": [{"@value": "Hoogte: 186 mm, diameter: 148-165 mm"}], "http://schemas.delving.eu/nave/terms/objectNumber": [{"@value": "2458"}]}]`

var _ = Describe("Graph", func() {

	Describe("when serializing JSON-LD", func() {

		Context("in generic json-ld input", func() {

			It("should have an id", func() {
				g := NewGraph("")
				err := g.Parse(strings.NewReader(testJSONLD), "application/ld+json")
				Expect(err).ToNot(HaveOccurred())
				body, err := g.GenerateJSONLD()
				Expect(err).ToNot(HaveOccurred())
				Expect(body).ToNot(BeNil())
				Expect(body).To(HaveLen(4))

				var buf bytes.Buffer
				g.SerializeFlatJSONLD(&buf)
				Expect(buf.String()).ToNot(BeEmpty())
			})

		})

		Describe("when creating an RDF interface model", func() {

			Context("and extracting the subject ID", func() {

				It("should extract the subject ID of a Resource", func() {
					t := NewTriple(
						NewResource("http://example.com/id1"),
						NewResource("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"),
						NewResource("http://example.org/vocab#Chapter"),
					)
					id := t.GetSubjectID()
					Expect(id).To(Equal("http://example.com/id1"))
				})

				It("should extract the subject ID of a BlankNode", func() {
					t := NewTriple(
						NewBlankNode(10),
						NewResource("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"),
						NewResource("http://example.org/vocab#Chapter"),
					)
					id := t.GetSubjectID()
					Expect(id).To(Equal("_:n10"))
				})

			})

			Context("and extracting the RDF type of the triple", func() {

				It("should give back the type when the predicate is rdf:type", func() {
					s := "http://example.com/id1"
					o := "http://example.org/vocab#Chapter"
					t := NewTriple(
						NewResource(s),
						NewResource("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"),
						NewResource(o),
					)
					Expect(NewResource(o)).To(Equal(NewResource(o)))
					ttype, ok := t.GetRDFType()
					Expect(ok).To(BeTrue())
					Expect(ttype).ToNot(BeEmpty())
					Expect(ttype).To(Equal(o))

				})

				It("should give back nil when the predicate is not rdf:type", func() {
					s := "http://example.com/id1"
					t := NewTriple(
						NewResource(s),
						NewResource("http://example.org/vocab#contains"),
						NewResource("http://example.org/vocab#Chapter"),
					)
					ttype, ok := t.GetRDFType()
					Expect(ok).To(BeFalse())
					Expect(ttype).To(BeEmpty())

				})
			})

			Context("and creating as subject map", func() {

				It("should not duplicate subjects", func() {
					m := map[string]*LdEntry{}
					s := "http://example.com/id1"
					o := "http://example.org/vocab#Chapter"
					t := NewTriple(
						NewResource(s),
						NewResource("http://www.w3.org/1999/02/22-rdf-syntax-ns#type"),
						NewResource(o),
					)
					err := AppendTriple(m, t)
					Expect(err).ToNot(HaveOccurred())
					err = AppendTriple(m, t)
					Expect(err).ToNot(HaveOccurred())
					Expect(m).To(HaveLen(1))
					ld, ok := m[s]
					Expect(ok).To(BeTrue())
					Expect(ld.ID).To(Equal(s))
					Expect(ld.Types).To(ContainElement(o))
					Expect(ld.Types).To(HaveLen(1))

				})

				It("should append predicates", func() {
					m := map[string]*LdEntry{}

					s := NewBlankNode(10)
					p := "http://example.org/vocab#contains"
					o := "http://example.org/vocab#Chapter"
					t := NewTriple(
						s,
						NewResource(p),
						NewResource(o),
					)
					err := AppendTriple(m, t)
					Expect(err).ToNot(HaveOccurred())
					ld, ok := m[GetResourceID(s)]
					Expect(ok).To(BeTrue())
					Expect(ld.ID).To(Equal("_:n10"))
					Expect(ld.Types).To(HaveLen(0))
					Expect(ld.Predicates).To(HaveLen(1))
					ldo, ok := ld.Predicates[p]
					Expect(ok).To(BeTrue())
					Expect(ldo[0].ID).To(Equal(o))
					Expect(ldo[0].Value).To(BeEmpty())
					Expect(ldo[0].Language).To(BeEmpty())
					Expect(ldo[0].Datatype).To(BeEmpty())

					t2 := NewTriple(
						s,
						NewResource(p),
						NewLiteralWithLanguage("Chapter 1", "en"),
					)
					err = AppendTriple(m, t2)
					ld, ok = m[GetResourceID(s)]
					Expect(ok).To(BeTrue())
					Expect(err).ToNot(HaveOccurred())
					contains, ok := ld.Predicates[GetResourceID(NewResource(p))]
					Expect(ok).To(BeTrue())
					Expect(contains).To(HaveLen(2))

					//jsonld := ld.AsEntry()
					//Expect(jsonld).To(Equal(""))

				})

			})
		})

	})
})
