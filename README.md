# bot-for-food

To start MongoDB run:
Run docker-compose up


Access MongoDB shell: 
sudo docker ps (copy Container ID)
sudo docker exec -it {contained ID} mongo


Shell commands:

show dbs	

use {DB name}

Get documents by time:
db.orderbook.find({_id: {$gt: ObjectId.fromDate( new Date('2019-11-13') ) } })

Clean collection:
db.orderbook.remove({})


### run Client:

1) Enter to `/public/client_src` and run `npm install`
2) open Cherome and run `http://localhost:9090/app/#!/sidemenu/main`