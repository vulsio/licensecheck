module github.com/vulsio/licensecheck

go 1.17

require (
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.7
	github.com/google/licenseclassifier v0.0.0-20210722185704-3043a050f148
	github.com/urfave/cli/v2 v2.3.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
)

replace github.com/vulsio/licensecheck => ./
