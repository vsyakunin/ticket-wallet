# Ticket Wallet

This is a simple ticket-wallet API. It is fully dockerized.

## Table of contents

**[1 Response Data format](#1)**      
**[2 API](#2)**   
&emsp;**[2.1 Get layout](#2.1)**  
&emsp;**[2.2 Start seating arrangement](#2.2)**  
&emsp;**[2.3 Get seating arrangement](#2.3)**  
**[3 Requirements](#3)**  
**[4 Deploy](#4)**

****
<a name="1">1. Response data format</a>
-  
**Response data format:** JSON                

**Sample error response:**
```js
{
    "title": "invalid parameter", 
    "description": "group #1: size can't be smaller than 1"
}
```

****
## <a name="2">2. API</a>

### <a name="2.1">2.1 Get layout</a>

This endpoint is used to see what unfilled layout looks like.

**URL:**

&emsp;`.../api/v1/layout`  

**METHOD: GET**

**Status codes:**  
&emsp;`200` - OK,  
&emsp;`500` - server error,  

**Response Body**
```js
{
    "name": "test hall",                        // hall name, string
	"sections": [                           // array of sections
        {
            "name": "section 1",                // section name, string
            "is_curved": false,                 // is section curved, bool
            "rows": [                           // array of rows
                {
                    "num": 1,                    // row number, int                 
                    "seats": [                   // array of seats
                        {
                            "num": 1,            // seat number (as it numbered in a hall), int
                            "actual_num": 1,     // seat actual number (in order, from left to right), int
                            "rank": "1st",       // seat rank, string
                            "is_free": true,     // is seat free, bool
                            "is_blocked": false, // is seat blocked (e.g. for technical purpose), bool
                            "taken_by": ""       // name of person who bought this seat (empty if the seat is free or blocked), string
                        }, ...
                    ]
                }, ...
            ]
        }, ...
}
```

### <a name="2.2">2.2 Start seating arrangement</a>

This endpoint triggers the seating algorithm in non-blocking way.

**URL:**

&emsp;`.../api/v1/seating/start`

**METHOD: POST**

**POST Body**
```js
{
    "groups"
:
    [                          // array of people groups
        {
            "size": 1,         // size of a group, int, must be bigger than 1
            "name": "group1"   // name of a person/group, string, must be non-empty
        }, ...
    ]
}
```

**Status codes:**  
&emsp;`200` - OK,  
&emsp;`400` - bad request (invalid request body)  
&emsp;`500` - server error,

**Response Body**

```js
{
	"task_id": "cd021327-56ab-4565-abad-754f9bbb5bb4", // task ID assigned to task filling the layout with groups from request, UUID, 
	"status": "created",                               // task status, string
	"payload": {}                                      // payload will always be empty for this endpoint
}
```

### <a name="2.3">2.3 Get seating arrangement</a>

This endpoint check the result of filling the layout with seating algorithm and returns the filled layout when it's completed.

**URL:**

&emsp;`.../api/v1/seating/result?taskId={taskId}`  

**METHOD: GET**

**URL parameters:**  
&emsp;`taskId` - *// uuid, task id received in response body of /seating/start endpoint*

**Status codes:**  
&emsp;`200` - OK,  
&emsp;`400` - bad request (invalid request body)  
&emsp;`500` - server error,

**Response Body**

```js
{
	"task_id": "cd021327-56ab-4565-abad-754f9bbb5bb4", // requested task ID, UUID, 
	"status": "completed",                             // task status, string
	"payload": {                                       // payload will always be empty for this endpoint
            "name": "test hall",                           // hall name, string
            "sections": [                                  // array of sections
            {
                "name": "section 1",                       // section name, string
                "is_curved": false,                        // is section curved, bool
                "rows": [                                  // array of rows
                    {
                        "num": 1,                          // row number, int                 
                        "seats": [                         // array of seats
                            {
                                "num": 1,                  // seat number (as it numbered in a hall), int
                                "actual_num": 1,           // seat actual number (in order, from left to right), int
                                "rank": "1st",             // seat rank, string
                                "is_free": false,          // is seat free, bool
                                "is_blocked": false,       // is seat blocked (e.g. for technical purpose), bool
                                "taken_by": "John Doe"     // name of person who bought this seat (empty if the seat is free or blocked), string
                            }, ...
                        ]
                    }, ...
                ]
            }, ...
        }
}
```

Possible task statuses:

&emsp;`created` - task has just been created, processing hasn't started yet (payload field is empty)  
&emsp;`processing` - task is processing (payload field is empty)  
&emsp;`completed` - task is completed (payload field contains filled layout)  
&emsp;`error` - something went wrong during processing (payload field is empty)

## <a name="3">3. Requirements</a>

&emsp; `docker`  
&emsp; `docker-compose`

## <a name="4">4. Deploy</a>

1. Clone this repository.
2. Create container with `docker-compose -f docker-compose.yml up -d`
3. Server will start running. 