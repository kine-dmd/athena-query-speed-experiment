# athena-query-speed-test
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
