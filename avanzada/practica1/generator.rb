# Define un método para generar un vector random de una longitud dada
def generate_random_array(length)
  Array.new(length) { rand(1..1000000) } # Cambia el rango como sea necesario
end

# Especifica el nombre del fichero de salida
file_name = "random_arrays.txt"

# Vector de diferentes longitudes
big_array_lengths = [5, 50, 500, 5000, 10000, 100000, 1000000, 10000000, 150000000]
mediumsmall_array_lengths = [5, 50, 500, 5000, 10000, 100000]
mediumbig_array_lengths = [5, 50, 500, 5000, 10000, 100000, 1000000, 5000000]
small_array_lengths = [5, 10, 50, 100, 500]
# Apertura de fichero para escribir
File.open("./datasets/big.txt", "w") do |file|
  big_array_lengths.each do |length|
    # Generar vector aleatorio
    random_array = generate_random_array(length)
    
    # Escribir el vector en el fichero en lineas separadas
    file.puts(random_array.join(' '))
  end
end

File.open("./datasets/mediumsmall.txt", "w") do |file|
  mediumsmall_array_lengths.each do |length|
    # Generar vector aleatorio
    random_array = generate_random_array(length)
    
    # Escribir el vector al fichero en lineas separadas
    file.puts(random_array.join(' '))
  end
end

File.open("./datasets/mediumbig.txt", "w") do |file|
  mediumbig_array_lengths.each do |length|
    # Generar vector aleatorio
    random_array = generate_random_array(length)
    
    # Escribir el vector al fichero en lineas separadas
    file.puts(random_array.join(' '))
  end
end

File.open("./datasets/small.txt", "w") do |file|
  small_array_lengths.each do |length|
   # Generar vector aleatorio
    random_array = generate_random_array(length)
    
    # Escribir el vector al fichero en lineas separadas
    file.puts(random_array.join(' '))
  end
end

puts "Random arrays have been written to #{file_name}"
