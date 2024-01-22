package main

type MalUser struct {
	Id              int          `json:"id"`
	Name            string       `json:"name"`
	Picture         string       `json:"picture"`
	Gender          string       `json:"gender"`
	Birthday        string       `json:"birthday"`
	Location        string       `json:"location"`
	JoinedAt        string       `json:"joined_at"`
	AnimeStatistics MalUserStats `json:"anime_statistics"`
	TimeZone        string       `json:"time_zone"`
	IsSupporter     bool         `json:"is_supporter"`
}

type MalUserStats struct {
	NumItemsWatching    int     `json:"num_items_watching"`
	NumItemsCompleted   int     `json:"num_items_completed"`
	NumItemsOnHold      int     `json:"num_items_on_hold"`
	NumItemsDropped     int     `json:"num_items_dropped"`
	NumItemsPlanToWatch int     `json:"num_items_plan_to_watch"`
	NumItems            int     `json:"num_items"`
	NumDaysWatched      float64 `json:"num_days_watched"`
	NumDaysWatching     float64 `json:"num_days_watching"`
	NumDaysCompleted    float64 `json:"num_days_completed"`
	NumDaysOnHold       float64 `json:"num_days_on_hold"`
	NumDaysDropped      float64 `json:"num_days_dropped"`
	NumDays             float64 `json:"num_days"`
	NumEpisodes         int     `json:"num_episodes"`
	NumTimesRewatched   int     `json:"num_times_rewatched"`
	MeanScore           float64 `json:"mean_score"`
}
