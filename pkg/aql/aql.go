package aql

/*

ToDo: We can do better here by utilising some creational/builder patterns
to create different AQL as per need.

For example :

aql:= NewAQLItems.
WithInludeOptions().
WithSorting().
WithLimit().
Encode()

Functions like WithInludeOptions, WithSorting and WithLimit can add capabilities to
the AQL and Encode should finally parse it into a Payload string that can be POSTed
to the REST API

But above can be improved by writing AQLClient that abstracts on how
the result is produce.

Example :

aqlClient:= NewAQLClient.Get().Sort(sortOptions).Limit(n)

*/

// For now just limiting to requirement

const PayLoad = "items.find(\n{\n\"repo\":\"backend-maven\"\n}\n).include(\"stat\").sort(\n{\n\"$asc\" : [\"stat.downloads\"]\n}\n).limit(2)"
