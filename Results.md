# TEST RESULTS (In Seconds)

## With 100.000 Entries:

|  TASK    												   |      Python               |  GO            |
|----------------------------------------------------------|:-------------------------:|---------------:|
|"insert single" no parallelism -- Duration:  			   |    88.62129855155945      |  58.1979807	|
|"insert single" with parallelism.  -- Duration: 		   |	55.817471742630005	   |  0.5619999     |
|"insert bulk" -- Duration: 							   |    2.331458568572998      |  8.5179686     |
|"findingone(filter)" with print statements(x10)--Duration:| 	0.012998819351196289   |  0.0089954		|
|"Count Documents" in collection: 						   |    2960	(total filtered entries)			   |  2925	(total filtered entries)		|
|"count_documents(filter)",  --- Duration: 				   |    0.18000030517578125    |  0.1860008		|
|"find(filter)" with print(x10)--- Duration:			   |	0.007998943328857422   |  0.009004		|


## With 500.000 Entries:

|  TASK    												   |      Python                |  GO            |
|----------------------------------------------------------|:--------------------------:|---------------:|
|"insert single" no parallelism -- Duration:  			   |  	431.0530788898468  		|	284.461093	 |
|"insert single" with parallelism.  -- Duration: 		   |	275.8861155509949		|	5.1890711    |
|"insert bulk". -- Duration: 							   | 	11.506523132324219		|	58.5371364	 |						
|"findingone(filter)" with print (x10).Duration: 		   |	0.014000892639160156	|	0.0179986    |
|"Count Documents" in collection: 						   |	17830 (total filtered entries)					|	17979 (total filtered entries)   	 |
|"count_documents(filter)",  --- Duration: 				   |	1.0300014019012451		|	1.1010009	 |
|"Find(filter)" with print(x10)--- Duration:  			   |	0.007997751235961914	|	0.0109986    |


## Glory to the Great Evron Empire!

