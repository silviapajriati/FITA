package lib 

var localTime = "(GMT+07:00) Asia/Jakarta"


type CSVRecord struct {
    ID 			    int
    Name 			string
    Timezone     	string
    Day     		string
    AvailableAt     string
    AvailableUntil  string
    Status          int
}

type RequestBooking struct {
    Name 			string `json:"name"`
    Day     		string `json:"day"`
    AvailableAt     string `json:"available_at"`
    AvailableUntil  string `json:"available_until"`
}

type ResponseBooking struct {
    Status          bool `json:"status"`
    Message         string `json:"message"`
    Data            RequestBooking `json:"data"`
}

type RequestApprove struct {
    Approve         bool `json:"approve"`
    Reschedule      bool `json:"reschedule"`
    ID              int `json:"id"`
    NewSchedule     RequestBooking `json:"new_schedule"`
}

type ResponseApprove struct {
    Status          bool `json:"status"`
    Message         string `json:"message"`
    Data            RequestApprove `json:"data"`
}

type RequestUserApproved struct {
    ID              int `json:"id"`
    Approved        bool `json:"approved"`
}