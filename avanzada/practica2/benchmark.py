##############################################################################
#                                                                            #
#     Archivo: benchmark.py                                                  #
#     Fecha de última revisión: 16/11/2023                                   #
#     Autores: Francisco Javier Pizarro 821259                               #
#              Jorge Solán Morote   	816259                               #
#     Comms:                                                                 #
#           Este script se emplea de para realizar los benchmark             #
#																			 #
##############################################################################
from datasetGen import *
import subprocess
import json

def execute_program():
    output = subprocess.check_output(['./testDataset', ''], text=True)
    last_line = output.strip().split('\n')[-1]
    return last_line

sizes = []
tiemposXsize = []
densities = []
tiemposXdensity = []
NSQUARES = 10

for size in range(5, 10):
    generate_dataset(size, 0.5, NSQUARES)
    last_line = execute_program()
    sizes.append(size)
    tiemposXsize.append(float(last_line) / NSQUARES)
    print(size)

density = 0
max_density = 1
step = 0.05
print("primer test finalizado")
while density < (max_density - step):
    density += step
    generate_dataset(9, density, NSQUARES)
    last_line = execute_program()
    densities.append(density)
    tiemposXdensity.append(float(last_line) / NSQUARES)

# Saving data to files
data_sizes = {'sizes': sizes, 'tiemposXsize': tiemposXsize}
data_densities = {'densities': densities, 'tiemposXdensity': tiemposXdensity}

with open('sizes_data.json', 'w') as sizes_file:
    json.dump(data_sizes, sizes_file)

with open('densities_data.json', 'w') as densities_file:
    json.dump(data_densities, densities_file)
