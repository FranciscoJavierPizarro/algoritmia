import matplotlib.pyplot as plt
import pandas as pd
import matplotlib.ticker as ticker
# Load data from CSV file
df = pd.read_csv('medidas.txt', delim_whitespace=True)
# Create a basic plot
plt.figure(figsize=(10, 6))
algoritms = [
# "HeapSort", "TreeSort", "RadixSort", "MergeSort", "QuickSort", "BubbleSort", "PancakeSort"
# "HeapSort", "TreeSort", "RadixSort", "MergeSort", "QuickSort"
"HeapSort", "TreeSort", "RadixSort", "MergeSort", "QuickSort", "ConcurrentMergeSort", "ConcurrentQuickSort"
]
for alg in algoritms:
    plt.plot(df['Size'], df[alg], label=alg)
plt.xlabel('Tamaño del vector')
plt.ylabel('ms')
plt.title('Sorting algoritms')
plt.legend()
plt.grid(True)

# min_y = 0  # Minimum Y-axis value
# max_y = 1  # Maximum Y-axis value
# plt.ylim(min_y, max_y)

# ax = plt.gca()
# ax.xaxis.set_major_locator(ticker.MaxNLocator(integer=True))
plt.show()