##############################################################################
#                                                                            #
#     Archivo: plots.py                                                      #
#     Fecha de última revisión: 16/11/2023                                   #
#     Autores: Francisco Javier Pizarro 821259                               #
#              Jorge Solán Morote   	816259                               #
#     Comms:                                                                 #
#           Este script se emplea de forma manual para generar las           #
#           gráficas empleadas en la memoria.               	  			 #
#																			 #
##############################################################################
import matplotlib.pyplot as plt
import json

# Reading data from files
with open('sizes_data.json', 'r') as sizes_file:
    sizes_data = json.load(sizes_file)

with open('densities_data.json', 'r') as densities_file:
    densities_data = json.load(densities_file)

# Plotting sizes data
sizes = sizes_data['sizes']
tiemposXsize = sizes_data['tiemposXsize']

plt.plot(sizes, tiemposXsize)
plt.xlabel('N')
plt.ylabel('Tiempo medio en segundos')
plt.title('Relación tamaño-coste temporal')
plt.savefig("benchmarkSizes.png")
plt.show()

# Plotting densities data
densities = densities_data['densities']
tiemposXdensity = densities_data['tiemposXdensity']

plt.plot(densities, tiemposXdensity)
plt.xlabel('Densidad')
plt.ylabel('Tiempo medio en segundos')
plt.title('Relación densidad-coste temporal')
plt.savefig("benchmarkDensidad.png")
plt.show()
