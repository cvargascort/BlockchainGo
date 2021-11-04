podcastsCollection 
:=
 client
.
Database
(
"quickstart"
)
.
Collection
(
"podcasts"
)

id
,
 
_
 
:=
 primitive
.
ObjectIDFromHex
(
"5d9e0173c1305d2a54eb431a"
)

result
,
 err 
:=
 podcastsCollection
.
UpdateOne
(

    ctx
,

    bson
.
M
{
"_id"
:
 id
}
,

    bson
.
D
{

        
{
"$set"
,
 bson
.
D
{
{
"author"
,
 
"Nic Raboy"
}
}
}
,

    
}
,

)

if
 err 
!=
 
nil
 
{

    log
.
Fatal
(
err
)

}

fmt
.
Printf
(
"Updated %v Documents!\n"
,
 result
.
ModifiedCount
)
