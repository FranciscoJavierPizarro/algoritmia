# Define a method to generate a random array of given length
def generate_random_array(length)
  Array.new(length) { rand(1..1000000) } # Change the range as needed
end

# Specify the file name
file_name = "random_arrays.txt"

# Array of different lengths
big_array_lengths = [5, 50, 500, 5000, 10000, 100000, 1000000, 10000000, 150000000]
mediumsmall_array_lengths = [5, 50, 500, 5000, 10000, 100000]
mediumbig_array_lengths = [5, 50, 500, 5000, 10000, 100000, 1000000, 5000000]
small_array_lengths = [5, 10, 50, 100, 500]
# Open a file for writing
File.open("./datasets/big.txt", "w") do |file|
  big_array_lengths.each do |length|
    # Generate a random array
    random_array = generate_random_array(length)
    
    # Write the array to the file on a separate line
    file.puts(random_array.join(' '))
  end
end

File.open("./datasets/mediumsmall.txt", "w") do |file|
  mediumsmall_array_lengths.each do |length|
    # Generate a random array
    random_array = generate_random_array(length)
    
    # Write the array to the file on a separate line
    file.puts(random_array.join(' '))
  end
end

File.open("./datasets/mediumbig.txt", "w") do |file|
  mediumbig_array_lengths.each do |length|
    # Generate a random array
    random_array = generate_random_array(length)
    
    # Write the array to the file on a separate line
    file.puts(random_array.join(' '))
  end
end

File.open("./datasets/small.txt", "w") do |file|
  small_array_lengths.each do |length|
    # Generate a random array
    random_array = generate_random_array(length)
    
    # Write the array to the file on a separate line
    file.puts(random_array.join(' '))
  end
end

puts "Random arrays have been written to #{file_name}"
