set title "Medida de rendimiento de los 2 algoritmos en los ficheros originales"
set ylabel "Tiempo en us"
set xlabel "Nº Palabras"
set grid
#unset key #esconde leyenda con nombre del fichero de datos
set key font ',14'

# set the X range to include all the data points
#set xrange [0:10000000]
set multiplot
plot "./precalc.txt" with lines title 'Precalculado' linestyle 1 linecolor rgb "red", \
"./pdonfly.txt" with lines title 'Al vuelo' linestyle 2 linecolor rgb "blue", \
#key outside right top
pause -1