# cmpe273-lab-3
The purpose of this lab is to understand how consistent hashing works and to implement a simple RESTful key-value cache data store.


You have to run the server first by typing 

go run server_lab3.go 

Now,the server is running at three ports 3000,3001 and 3002

You can run the client by 

go run ConsistentHashing_client.go

The client implements consistent hashing and it shards following data set into three server instances that you created.

    {Key => Value}
    1 => a
    2 => b
    3 => c
    4 => d
    5 => e
    6 => f
    7 => g
    8 => h
    9 => i
    10 => j
    
    
    
    
