#!/usr/bin/ruby
def generate_input_file(filename,size)
    return unless size
    File.open("#{filename}#{size}.txt", "w") do |file|
        ntest = 1#se podría usar otro valor pero de esta forma podemos comparar los tiempos
        ntest.times do
            n = rand(size..(size * 2))
            m = rand(1..size)
            p = rand(1..(size * 2))
            file.puts "#{n} #{m} #{p}"
            return unless p
            p.times do
                s = rand(1..m-1)
                return unless s
                e = rand(s+1..m)
                j = rand(1..n)
                file.puts "#{s} #{e} #{j}"
            end
        end
        file.puts "0 0 0"
    end
end
  
generate_input_file("tests/input",10)
generate_input_file("tests/input",25)
generate_input_file("tests/input",50)
generate_input_file("tests/input",75)
generate_input_file("tests/input",85)