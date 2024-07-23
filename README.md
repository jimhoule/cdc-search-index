 Simple Change Data Capture (CDC) pipeline

 The Users Module is connected to an Arango database where all write operation happen. A user is created, 
 updated and deleted from this module (there are also a functions to fetch just for convenience). The Users Module 
 is also connected to a Kafka Producer which will stream the result of a write operation (create, update and delete)
 to the Search Module which will create, update or delete the same user in Elastic Search search index. The Search Module
 which is connecte to the search index exposes to functions to fetch users from a Rest API (http).

 Architecture and design patterns:
  * Domain Drvien design (DDD)
  * Hexagonal Architecture
  * 4 Layered Architecture
  * Event Driven Architecture (EDA)

 Steps:
   * Fill necessary env variables
   * Run docker-compose file to start necessary services
   * Run go run . in root folder to start application