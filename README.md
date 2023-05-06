# Money_Manager_App

Only Transaction Service should be able to access json file or database -Why , why can't other services access it? 


Reading and writing to a file should be a separate function package altogether - Why? --> For now we are using json, Lets say I expand it to excel .... This will come in handy ... 
Then what about database ? --> Since this tool is for actual using purpose as well, as of now there is no idea to use a database

## Main is the executable package: Entry point where the command executes from

##### Write code in such a way that data can be read from any source(json file, xcel file, database)

##### Extract all the configurations to a single point, so it can be changed according to different envs (cloud, local etc)

##### Check If we can use any mutual Authentication method(jwt etc)(Will be only required if this is going to be a API)

##### How many separate api's should be defined?(Transactions API, Savings API etc)

##### Should we make everything so complex?

