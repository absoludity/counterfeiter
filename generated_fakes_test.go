package main_test

import (
	"errors"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
	"github.com/maxbrunsfeld/counterfeiter/fixtures/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generated fakes", func() {
	var fake *fakes.FakeSomeInterface

	BeforeEach(func() {
		fake = new(fakes.FakeSomeInterface)
	})

	It("implements the interface", func() {
		var someInterface fixtures.SomeInterface = fake
		Expect(someInterface).NotTo(BeNil())
	})

	It("can have its behavior configured using stub functions", func() {
		fake.DoThingsStub = func(arg1 string, arg2 uint64) (int, error) {
			Expect(arg1).To(Equal("stuff"))
			Expect(arg2).To(Equal(uint64(5)))
			return 3, errors.New("the-error")
		}

		num, err := fake.DoThings("stuff", 5)

		Expect(num).To(Equal(3))
		Expect(err).To(Equal(errors.New("the-error")))
	})

	It("can have its return values configured", func() {
		fake.DoThingsReturns(3, errors.New("the-error"))

		num, err := fake.DoThings("stuff", 5)
		Expect(num).To(Equal(3))
		Expect(err).To(Equal(errors.New("the-error")))
	})

	It("doesn't mind when no stub is provided", func() {
		fake.DoThings("stuff", 5)
		fake.DoNothing()
	})

	It("records the arguments it was called with", func() {
		Expect(fake.DoThingsCalls()).To(HaveLen(0))

		fake.DoThings("stuff", 5)

		Expect(fake.DoThingsCalls()).To(HaveLen(1))
		Expect(fake.DoThingsCalls()[0].Arg1).To(Equal("stuff"))
		Expect(fake.DoThingsCalls()[0].Arg2).To(Equal(uint64(5)))
	})

	It("records its calls without race conditions", func() {
		go fake.DoNothing()

		Eventually(fake.DoNothingCalls, 1.0).Should(HaveLen(1))
	})
})