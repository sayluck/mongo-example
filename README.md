# mongo-example
some examples for mongodb used by mongo-go-driver

## Basic Usage
1.start a mongo container user docker-compose,this's not necessary if you have one. 
```
    cd $projectpath/mongo-example/assenbly/deploy/mongodb/
    docker-compose up -d
```
View the database information through the web page after completion.(http://127.0.0.1:8081)

2.the follow functions is supported and it will be more in the after.
```
    Connect()           // connect a mongo db,get a db client
    InseartOne()
    findByID()
    Find()
    FindWithRelate()   // Associate records of two documents
```

## TODO
The function is not perfect now and the following functions will be added.


```
inseart many
delete 
find
    sort
    limit
    skip
```
