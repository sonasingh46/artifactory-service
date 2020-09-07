package aql

/*

ToDo: We can do better here by utilising some creational/builder/templating patterns
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

But above can be improved by writing go based AQL query that abstracts on how
the result is produce.

Example :

aqlQuery:= NewAQLQuery.Get().Sort(sortOptions).Limit(n)

result:=aqlQuery.Exec()

*/

// For now just limiting to requirement

const PayLoad = "items.find(\n{\n\"repo\":\"backend-maven\"\n}\n).include(\"stat\").sort(\n{\n\"$desc\" : [\"stat.downloads\"]\n}\n).limit(2)"
