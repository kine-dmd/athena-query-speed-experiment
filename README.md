# Athena Query Speed Test
Code used to facilitate an investigation into how file size affects querying speed on AWS Athena.

## Running instructions
Run inside a `tmux` session with the following command, substituting the number of rows each time.
```
rm -rf /data/rp1615 && mkdir /data/rp1615 && cd ~/Documents && ./main 100000000 && rm -rf /data/rp1615 && exit
```

## Method
All tables had 100000000 rows. Query tested was `SELECT count(ax) FROM row100 WHERE ax > 0`.

## Results
Results:

| Rows per file | Time taken /s  | 
| ------------- |:--------------:| 
| 100           | 201.15         | 
| 1000          | 19.53          | 
| 10000         | 4.07           | 
| 100000        | 2.79           | 
| 1000000       | 2.35           | 
| 10000000      | 2.27           | 
| 100000000     | 2.83           | 


## Results analysis
![queryTimes](https://user-images.githubusercontent.com/26333869/60196348-d3c2ef80-9834-11e9-8956-53487cbfd2d8.png)
As shown in the graph, the query time is significantly longer when small files are used due to the additional overhead of creating new connections to S3 and reading additional metadata. The single file performance (largest entry) is also slower than multiple files as the query cannot be parallelised. 

## Further reading
Additional information can be found on the [AWS Big Data Blog](https://aws.amazon.com/blogs/big-data/top-10-performance-tuning-tips-for-amazon-athena/) and [Upsolver](https://www.upsolver.com/blog/aws-athena-performance-best-practices-performance-tuning-tips).
