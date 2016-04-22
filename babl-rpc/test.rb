require 'quartz'
require 'base64'

client = Quartz::Client.new(bin_path: './babl-rpc')

puts "Structs: #{client.structs}"
babl = client[:babl]
puts "Struct methods for #{babl.struct_name}: #{babl.struct_methods}"
puts "Response:"
res = babl.call('Module', 'Name' => 'larskluge/string-upcase', 'Stdin' => Base64.encode64('Send reinforcements'))
puts Base64.decode64(res["Stdout"])
