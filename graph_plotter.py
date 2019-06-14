import matplotlib.pyplot as plt


x = [100, 1000, 10000, 100000, 1000000, 10000000, 100000000]
y = [201.15, 19.53, 4.07, 2.79, 2.35, 2.27, 2.83]


fig, ax = plt.subplots()
ax.semilogx(x, y, 'x-')
plt.xlabel("Number of rows per file")
plt.xlim((800, 120000000))
plt.ylim((1.5, 21))
plt.ylabel("Time taken to execute query (s)")
plt.title("Query time over files containing a total of 100,000,000 data rows")

#plt.show()
plt.savefig('queryTimes', dpi=500)
