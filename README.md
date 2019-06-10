# athena-query-speed-test
Code used to facilitate an investigation into how file size affects querying speed on AWS Athena.


Run inside a `tmux` session with the following command, substituting the number of rows each time.
```
rm -rf /data/rp1615 && mkdir /data/rp1615 && cd ~/Documents && ./main 100000000 && rm -rf /data/rp1615 && exit
```