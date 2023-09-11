#!/usr/bin/ruby
# Function to generate a random integer array of a given length
def generate_random_array(length)
    Array.new(length) { rand(1..100) } # Change the range as needed
  end
  
  # Create three random integer arrays
  array1 = generate_random_array(5)
  array2 = generate_random_array(50)
  array3 = generate_random_array(500)
  
  # Open a file for writing
  file_name = "random_arrays.txt"
  File.open(file_name, "w") do |file|
    # Write each array to the file on a separate line
    file.puts(array1.join(' '))
    file.puts(array2.join(' '))
    file.puts(array3.join(' '))
  end
  
  puts "Random arrays have been written to #{file_name}"
  