<h1> Order Management Service using Go</h1>

<hr> 

<h2> Getting Started  </h2>

<h3> Installation </h3>
<code>docker-compose up -d</code> <br> <br>
Alternatively, run the following commands: <br>
<code>docker build --tag ordermanagement .</code> <br>
<code> docker run ordermanagement </code>

<hr> 
<h2> API Docs  </h2>
Base URL: <code> http://localhost:8000/ </code>

<h3>Get all orders </h3>
GET <code>http://localhost:8000/orders </code>

<h3> Get all orders sorted by field </h3>
GET <code>http://localhost:8000/orders/{field} </code>

<h3> Get all orders containing key-value pair </h3>
GET <code>http://localhost:8000/orders/{key}/{value} </code>

<h3> Add a new order </h3>
POST <code>http://localhost:8000/add</code> <br>
Body <code> {
"id": "abcdef-123456",
"status": "PENDING_INVOICE",
"items": [{
"id": "123456 ",
"description": "a product description",
"price": 12.40,
"quantity": 1
}],
"total": 12.40,
"currencyUnit": "USD"
} </code>

<h3> Update status of existing order </h3>
PUT <code>http://localhost:8000/update</code> <br>
Body <code> {
"id": "abcdef-123456",
"status": "PENDING_INVOICE"} </code>


<hr> 
<h2> Relational Database Schema </h2>

`orders`  <br>
`orderID` varchar(13) PRIMARY KEY NOT NULL <br>
`status` text NOT NULL <br>
`total` double NOT NULL <br>
`currencyUnit` varchar(3) NOT NULL <br> <br>

`items`  <br>
`itemID` varchar(6) NOT NULL PRIMARY KEY<br>
`orderID` varchar(13) DEFAULT NULL FOREIGN KEY REFERENCES `orders` (`orderID`) <br>
`description` text NOT NULL <br>
`price` double NOT NULL <br>
`quantity` int NOT NULL <br>


