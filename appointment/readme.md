======================================
            GET ALL DATA
======================================
- user can get all data from hit the API
- method : GET
- endpoint : localhost:8090/data/all

======================================
        BOOKING APPOINTMENT
======================================
- user booking appointment with the selected coach
- there is validation with equal date, coach, and day is forbidden
- format time is convert to time local (GMT+7)
- method : POST
- endpoint : localhost:8090/data/appointment
- request : {
    "name" : "Elyssa O'Kon",
	"day" : "Saturday",
	"available_at" : "7:00PM",
	"available_until": "10:30PM"
}
- if success then data will saved in csv file with status 0

======================================
    COACH CONFIRMATION & RESCHEDULE
======================================
- Coach can decide if want approve or reject the appointment
- If approved then data in CSV will update the status 1
- If rejected then data in CSV will update the status 3
- If reschedule then data in CSV will update the status 2
- reschedule have same validation with booking appointment
- method : POST
- endpoint : localhost:8090/data/confirmation
- request : {
	"approve": true,
	"reschedule": true,
	"id" : 27,
	"new_schedule": {
			"day" : "Saturday",
			"available_at" : "6:00AM",
			"available_until": "9:00AM"
	}
}

======================================
        USER CONFIRMATION
======================================
- User can approve or rejec data reschedule wich have status = 2
- If user approved then data in CSV will update the status 1
- If user rejected then data in CSV will update the status 3
- method : POST
- endpoint : localhost:8090/data/user-confirmation
- request : {
    "approved": false,
	"id" : 25
}
