# Define a method to generate a random array of given length
def generate_random_array(length)
  Array.new(length) { rand(1..1000000) } # Change the range as needed
end

# Specify the file name
file_name = "random_arrays.txt"

# Array of different lengths
array_lengths = [5, 50, 500, 5000, 10000, 100000, 1000000, 10000000, 100000000]

# Open a file for writing
File.open(file_name, "w") do |file|
  array_lengths.each do |length|
    # Generate a random array
    random_array = generate_random_array(length)
    
    # Write the array to the file on a separate line
    file.puts(random_array.join(' '))
  end
end

puts "Random arrays have been written to #{file_name}"
