======================================
            GET ALL DATA
======================================
- user can get all data from hit the API
- method : GET
- endpoint : localhost:8090/product/data

======================================
            CHECKOUT
======================================
- method : POST
- endpoint : localhost:8090/product/checkout
- request : [
	{
		"sku": "A304SD",
		"quantity": 3
	}
]
- Calculate total price when checkout
- User can checkout more than 1 type SKU
- Have fitur bundling in some item
- Have fitur discount in some item