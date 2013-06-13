#!/usr/bin/env ruby
# encoding: utf-8
require "xml"

def generate_values_statement(reader)
  statement = "(";
  arr = %w(houseid houseguid aoguid postalcode housenum buildnum strucnum okato eststatus strstatus startdate updatedate enddate counter).map(&:upcase)
  arr.each do |item|
    statement += "'#{reader.get_attribute(item)}',"
  end
  statement.chomp!(",")
  statement += ")"
  statement
end

def print_insert_statement(values)
  $insert_statement ||= "INSERT INTO houses(`houseid`, `houseguid`, `aoguid`, `postalcode`, `housenum`, `buildnum`, `strucnum`, `okato`, `eststatus`, `strstatus`, `startdate`, `updatedate`, `enddate`, `counter`) \n VALUES ";
  puts($insert_statement)
  puts(values.join(","))
  puts(";")
end

# reader = XML::Reader.file("./houses.xml")
reader = XML::Reader.file("/media/DATA/fias_xml/AS_HOUSE_20130608_bc00dafd-fcd0-40a6-b6ec-47dbf5a95be0.XML")

count = 0
values = []

while (reader.node_type != XML::Reader::TYPE_END_ELEMENT)
  reader.read
  if reader.name == "House"
    values << generate_values_statement(reader)
    count += 1
    if count % 1000 == 0
      print_insert_statement(values)
      values = []
    end
  end
end
print_insert_statement(values) if values.any?