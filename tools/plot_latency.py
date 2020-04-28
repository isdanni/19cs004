import os
import sys
import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import numpy as np

filename = sys.argv[1]
os.system("cat " + filename  + " | sed -n 's/\[\([0-9]*\)\].*latency=\([0-9]*\).*/\\1 \\2/p' > out.txt")

times, latency = [], []
with open("out.txt") as f:
    for line in f:
        data = line.split()
        times.append(float(data[0]))
        latency.append(float(data[1])/1000000.0)

min_time = min(times)
time_diff = [(t - min_time)/1000.0 for t in times]

print "average latency", np.average(np.array(latency))

plt.plot(time_diff, latency, 'ro', markersize=2)
plt.xlabel('Time (s)')
plt.ylabel('Latency (ms)')
plt.savefig("latency.pdf")

os.system("rm out.txt")
