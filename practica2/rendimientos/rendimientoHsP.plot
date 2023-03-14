set title "Medida de rendimiento de los algoritmos haskell para una P creciente"
set ylabel "Tiempo en ms"
set xlabel "Profunidad del árbol"
set grid
unset key #esconde leyenda con nombre del fichero de datos

set term qt 0
# set the X range to include all the data points
#set xrange [0:10000000]
set multiplot

plot "./rendimientoProfHaskSimulado.txt" with lines title 'Haskell Directo' linestyle 1 linecolor rgb "blue"
replot "./rendimientoProfHaskDirecto.txt" with lines title 'Haskell Simulado' linestyle 2 linecolor rgb "green"
pause -1