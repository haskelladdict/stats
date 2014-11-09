stats
=====

stats is a simple command utility for computing basics statistics on plain text 
data files. 

The input to stats can be any plain text single column data file. Stats will interpret
each line in the input file as a 64bit floating point number and then compute 
statistics on the parsed data. Currently stats evaluates the following quantities:

* mean
* variance
* median
* min
* max

Please note that the computation of the median currently requires stat to store the data
internally and thus consumes memory on the order of the data set itself.
