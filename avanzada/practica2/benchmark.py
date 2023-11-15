from datasetGen import *
import subprocess
import matplotlib.pyplot as plt

def execute_program():
    output = subprocess.check_output(['./testDataset', ''], text=True)
    last_line = output.strip().split('\n')[-1]
    return last_line

sizes = []
tiemposXsize = []
densities = []
tiemposXdensity = []
NSQUARES=10

for size in range(4, 9):
    generate_dataset(size,0.5,NSQUARES)
    last_line = execute_program()
    sizes.append(size)
    tiemposXsize.append(float(last_line)/NSQUARES)


density = 0
max_density = 1
step = 0.05

while density < (max_density - step):
    density += step
    generate_dataset(7,density,NSQUARES)
    last_line = execute_program()
    densities.append(density)
    tiemposXdensity.append(float(last_line)/NSQUARES)

plt.plot(sizes, tiemposXsize)
plt.xlabel('N')
plt.ylabel('Tiempo en segundos')
plt.title('Relación tamaño-coste temporal')
plt.show()

plt.plot(densities, tiemposXdensity)
plt.xlabel('Densidad del latinSquare')
plt.ylabel('Tiempo en segundos')
plt.title('Relación densidad-coste temporal')
plt.show()