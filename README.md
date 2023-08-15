# practice
Golang server with these endpoints:
<ul>
<li>Get People</li>  
  <li>Add People</li> 
  <li>Search People(by address,name,surname)</li> 
  <li>Get People by name</li> 
  <li>Filter People by age</li> 
</ul>

Routes

    GET /people/:name:
    Retrieve a person's details by their name.

    POST /people:
    Add a new person's details.

    GET /people:
    Retrieve a list of people's names.

    GET /person:
    Search for people by name and retrieve matching results.

    GET /person_by_surname:
    Search for people by surname and retrieve matching results.
    
    GET /person_by_address:
    Search for people by address and retrieve matching results.
    
    GET /filter:
    Filter people by age within a specified range.
---
Example Request for "/filter"

Route: GET /filter

URL: http://localhost:8080/filter?min_age=20&max_age=40

Query Parameters:

    min_age: Minimum age for filtering.
    max_age: Maximum age for filtering.

---
Example Request for "/person"

Route: GET /person

URL: http://localhost:8080/person?name=John

Query Parameters:

    name: name for search
---
Example Request for "/person_by_address"

Route: GET /person_by_address

URL: http://localhost:8080/person_by_address?address=Nevskaya

Query Parameters:

    address: address for search
---
Example Request for "/person_by_surname"

Route: GET /person_by_surname

URL: http://localhost:8080/person_by_surname?address=Smith

Query Parameters:

    surname: surname for search

