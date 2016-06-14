def preguntar()
  puts "Eres humano?"
  humano=gets.chomp
  if humano == "si" || humano == "no"
    if humano=="si"
      puts "Eres humano"
    else
      puts "no eres humano"
    end
  else
    return preguntar
  end


end

preguntar
