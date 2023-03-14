set title "Medida de rendimiento de los algoritmos haskell para una N creciente"
set ylabel "Tiempo en ms"
set xlabel "N=10**"
set grid
#unset key #esconde leyenda con nombre del fichero de datos

# set the X range to include all the data points
#set xrange [0:10000000]
set multiplot

plot "./rendimientoBallsHaskSimulado.txt" with lines title 'Haskell Simulado' linestyle 1 linecolor rgb "blue", \
"./rendimientoBallsHaskDirecto.txt" with lines title 'Haskell Directo' linestyle 2 linecolor rgb "green"
pause -1