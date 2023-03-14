set title "Medida de rendimiento de los 3 algoritmos"
set ylabel "Tiempo en ms"
set xlabel "Profunidad del árbol"
set grid
unset key #esconde leyenda con nombre del fichero de datos

set term qt 0
# set the X range to include all the data points
#set xrange [0:10000000]
set multiplot
plot "./rendimientoPyth.txt" with lines title 'Python' linestyle 1 linecolor rgb "red"
replot "./rendimientoHaskSimulado.txt" with lines title 'Haskell Simulado' linestyle 2 linecolor rgb "blue"
replot "./rendimientoHaskDirecto.txt" with lines title 'Haskell Directo' linestyle 3 linecolor rgb "green"
key outside right top
pause -1