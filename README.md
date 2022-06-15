# Assignment 2

## Repo

- github url: https://github.com/ajatdarojat45/ptp-glng-023-a2

## GET /orders

response:

```json
{
	"count": 2,
	"result": [
		{
			"ID": 0,
			"CreatedAt": "0001-01-01T07:07:12+07:07",
			"UpdatedAt": "0001-01-01T00:00:00Z",
			"DeletedAt": null,
			"id": 3,
			"customerName": "Kosasih edited",
			"updated_at": "0001-01-01T07:07:12+07:07",
			"items": [
				{
					"ID": 0,
					"CreatedAt": "2022-06-15T15:31:01.615362+07:00",
					"UpdatedAt": "2022-06-15T15:58:00.500465+07:00",
					"DeletedAt": null,
					"lineItemId": 1,
					"itemCode": "001",
					"description": "iPhone 12 Pro Max edited",
					"quantity": 1,
					"orderId": 3
				}
			]
		}
	]
}
```

## POST /orders

body:

```json
{
	"orderAt": "2019-11-09T21:21:46+00:00",
	"customerName": "Udin edited",
	"items": [
		{
			"itemCode": "001",
			"description": "iPhone 12 Pro Max edited",
			"quantity": 1
		},
		{
			"itemCode": "002",
			"description": "MacBook Pro edited",
			"quantity": 1
		}
	]
}
```

response

```json
{
	"result": "Data successfully created"
}
```

## PUT /orders/:orderId

body:

```json
{
	"ID": 1,
	"orderAt": "2019-11-09T21:21:46+00:00",
	"customerName": "Udin edited",
	"items": [
		{
			"lineItemId": 1,
			"itemCode": "001",
			"description": "iPhone 12 Pro Max edited",
			"quantity": 1
		},
		{
			"lineItemId": 2,
			"itemCode": "002",
			"description": "MacBook Pro edited",
			"quantity": 1
		}
	]
}
```

response

```json
{
	"result": "Data successfully updated"
}
```

## DELETE /orders/:orderId

response

```json
{
	"result": "Data successfully deleted"
}
```
