package main

import (
	"flag"
)

// Args input parameters
type Args struct {
	bucketSql     string
	bucketData    string
	pathSql       string
	pathData      string
	sqlParameters string
}

// getArgs input parameters parser
func getArgs() Args {
	var (
		bucketSql     string
		bucketData    string
		pathSql       string
		pathData      string
		sqlParameters string
	)

	flag.StringVar(&bucketSql,
		"bucket-sql",
		"",
		"s3 bucket with meta data, i.e. with sql queries.")

	flag.StringVar(&bucketData,
		"bucket-data",
		"",
		"s3 bucket to read/store data from/to.")

	flag.StringVar(&pathSql,
		"path-sql",
		"",
		"s3 object path to sql query.")

	flag.StringVar(&pathData,
		"path-data",
		"",
		"s3 object path to store/read data to/from.")

	flag.StringVar(&sqlParameters,
		"sql-parameters",
		"",
		"JSON with parameters to adjust the query prior to its execution.")

	flag.Parse()

	return Args{
		bucketSql:     bucketSql,
		bucketData:    bucketData,
		pathSql:       pathSql,
		pathData:      pathData,
		sqlParameters: sqlParameters,
	}
}
