#!/usr/bin/ruby
# Function to generate a random integer array of a given length
def generate_random_array(length)
    Array.new(length) { rand(1..1000000) } # Change the range as needed
  end
  
  # Create three random integer arrays
  array1 = generate_random_array(5)
  array2 = generate_random_array(50)
  array3 = generate_random_array(500)
  array4 = generate_random_array(5000)
  array5 = generate_random_array(10000)
  array6 = generate_random_array(100000)
  array7 = generate_random_array(1000000)
  array8 = generate_random_array(10000000)
  array9 = generate_random_array(100000000)
  # Open a file for writing
  file_name = "random_arrays.txt"
  File.open(file_name, "w") do |file|
    # Write each array to the file on a separate line
    file.puts(array1.join(' '))
    file.puts(array2.join(' '))
    file.puts(array3.join(' '))
    file.puts(array4.join(' '))
    file.puts(array5.join(' '))
    file.puts(array6.join(' '))
    file.puts(array7.join(' '))
    file.puts(array8.join(' '))
    file.puts(array9.join(' '))
  end
  
  puts "Random arrays have been written to #{file_name}"
  