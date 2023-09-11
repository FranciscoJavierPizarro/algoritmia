set title "Medida de rendimiento de los 2 algoritmos en los ficheros custom"
set ylabel "Tiempo en ms"
set xlabel "Tamaño del fichero"
set grid
#unset key #esconde leyenda con nombre del fichero de datos
set key font ',14'

# set the X range to include all the data points
#set xrange [0:10000000]
set multiplot
plot "./rramif.txt" with lines title 'Ramificación' linestyle 1 linecolor rgb "red", \
"./rlineal.txt" with lines title 'Lineal' linestyle 2 linecolor rgb "blue", \
#key outside right top
pause -1