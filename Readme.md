# MongoDB driver speed test
This repo contains code to compare MongoDB performances using GO and PYTHON. Basically I wanted to see the difference in driver speed between go and python.
I was expecting GO to be much faster. There are also a table with the results in second that i got.
NB: MongoDB was lanched using DOCKER compose and exposed on localhost. 
The Python code and the GO code was run (one at the time) with clean DB direclty on the host machine.


### What it is
There are two folder:
- The python_speed_test contains the python code used to test the python mongodb driver.
- The go_speed_test contains the GO script used for testing the GO performance with MongoDB.
- Docker-compose file to start the DB


### How to use it.
	1- docker-compose up and you have the MongoDB exposed on your localhost
	2- run the Python_speed_test and see the performance. 
	3- clean the DB 
	4- run the GO_speed_test (you need to build it) and see the performance 
	5- Compare the results. 


### FAQ
General:
-  Q.is this code optimized?
- - A. No, i made it just for curiosity
- Q.What did you use for parallelize? 
- - A. I used go routines for GO and threads for python (I know, is not really fair.)
- Q. What results did you get? 
- - A. You can see from results file. there are two tables there with 100k and 500k entries. 
- Q. What tests did you run?
- - A. Look at the code (really is just 1 file!) or at the result file.

##### GLORY to the Great Evron Empire
