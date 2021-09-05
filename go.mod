module example.com/sorts

go 1.17

replace example.com/bubble => ./bubble
replace example.com/merge => ./merge
replace example.com/quick => ./quick

require example.com/bubble v0.0.0-00010101000000-000000000000
require example.com/merge v0.0.0-00010101000000-000000000000
require example.com/quick v0.0.0-00010101000000-000000000000
