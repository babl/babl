package bablnaming

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Context("#RequestPathToTopic", func() {
		It("converts a request path to a Kafka topic", func() {
			Expect(RequestPathToTopic("/babl.larskluge.Telegram/IO")).To(Equal("babl.larskluge.Telegram.IO"))
		})
	})
	Context("#TopicToModuleName", func() {
		It("converts a Kafka topic to module name", func() {
			Expect(TopicToModule("babl.larskluge.Telegram.IO")).To(Equal("larskluge/telegram"))
		})
		It("supports two word module names", func() {
			Expect(TopicToModule("babl.larskluge.RenderWebsite.IO")).To(Equal("larskluge/render-website"))
		})
		It("also works with a meta topic", func() {
			Expect(TopicToModule("babl.babl.Events.Meta")).To(Equal("babl/events"))
		})
	})
	Context("#ModuleToTopic", func() {
		It("converts a module name to a kafka topic", func() {
			Expect(ModuleToTopic("larskluge/hi", false)).To(Equal("babl.larskluge.Hi.IO"))
		})
		It("converts a module name to a meta kafka topic", func() {
			Expect(ModuleToTopic("larskluge/hi", true)).To(Equal("babl.larskluge.Hi.meta"))
		})
		It("two word module", func() {
			Expect(ModuleToTopic("larskluge/image-resize", false)).To(Equal("babl.larskluge.ImageResize.IO"))
		})
	})
})
