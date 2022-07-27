"use strict";

db;
show dbs;

use home;

db.createCollection("contacts");

db.contacts.insert({
  "name": "Jon Wexler",
  "email": "jon@jonwexler.com",
  "note": "Decent guy.",
});

show collections;


var coll = db.getCollection("contacts");

coll.findOne();
coll.find().pretty();

db.contacts.find({"_id": ObjectId("604afd43241818278390169b")}).limit(1);

db.contacts.updateOne({"name": "Jon Wexler"}, {"name": "Jon_Wexler"});
