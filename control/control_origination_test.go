package control


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/fedramp-templater/fixtures"
)

var _ = Describe("controlOrignation", func() {
	Describe("newControlOrignation", func() {
		It("should return", func() {
			doc := fixtures.LoadSSP("FedRAMP_ac-2-1_v2.1.docx")
			defer doc.Close()
			tables, err := doc.SummaryTables()
			Expect(err).NotTo(HaveOccurred())
			st := NewSummaryTable(tables[0])
			co, err := newControlOrigination(st)
			// Check error
			Expect(err).ToNot(HaveOccurred())
			// Check number of control origination.
			Expect(len(co.origins)).To(Equal(7))

			// Find the unchecked service provided corporate origination.
			Expect(co.origins[serviceProviderCorporateOrigination].getTextValue()).
				To(ContainSubstring(serviceProviderCorporateOrigination))
			Expect(co.origins[serviceProviderCorporateOrigination].isChecked()).To(Equal(false))

			// Find the unchecked service provided system specific origination.
			Expect(co.origins[serviceProviderSystemSpecificOrigination].getTextValue()).
				To(ContainSubstring(serviceProviderSystemSpecificOrigination))
			Expect(co.origins[serviceProviderSystemSpecificOrigination].isChecked()).To(Equal(false))

			// Find the unchecked service provided hybrid origination.
			Expect(co.origins[serviceProviderHybridOrigination].getTextValue()).
				To(ContainSubstring(serviceProviderHybridOrigination))
			Expect(co.origins[serviceProviderHybridOrigination].isChecked()).To(Equal(false))

			// Find the unchecked configured by customer origination.
			Expect(co.origins[configuredByCustomerOrigination].getTextValue()).
				To(ContainSubstring(configuredByCustomerOrigination))
			Expect(co.origins[configuredByCustomerOrigination].isChecked()).To(Equal(false))

			// Find the unchecked provided by customer origination.
			Expect(co.origins[providedByCustomerOrigination].getTextValue()).
				To(ContainSubstring(providedByCustomerOrigination))
			Expect(co.origins[providedByCustomerOrigination].isChecked()).To(Equal(false))

			// Find the unchecked shared origination.
			Expect(co.origins[sharedOrigination].getTextValue()).
				To(ContainSubstring(sharedOrigination))
			Expect(co.origins[sharedOrigination].isChecked()).To(Equal(false))

			// Find the unchecked inherited origination.
			Expect(co.origins[inheritedOrigination].getTextValue()).
				To(ContainSubstring(inheritedOrigination))
			Expect(co.origins[inheritedOrigination].isChecked()).To(Equal(false))
		})
	})
})
