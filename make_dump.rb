#!/usr/bin/env ruby
# encoding: utf-8
require "xml"

reader = XML::Reader.file("./houses.xml")

reader.read
10.times do
  reader.read
  reader.read
  p reader.name
  attributes = []
  reader.attribute_count.times do |index|
    attributes << reader.get_attribute_no(index)
  end
  p attributes.join ','
end
