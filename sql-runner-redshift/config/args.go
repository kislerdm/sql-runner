package config

import (
	"flag"
	"fmt"
)

// CmdArgs command line input parameters
type CmdArgs struct {
	BucketSQL     string
	PathSQL       string
	SQLParameters string
	SQL           string
}

// ReadArgs command line input parameters parser
func ReadArgs() (*CmdArgs, error) {
	var args CmdArgs

	flag.StringVar(&args.BucketSQL,
		"bucket-sql",
		"",
		"s3 bucket with meta data, i.e. with sql queries.")

	flag.StringVar(&args.PathSQL,
		"path-sql",
		"",
		"s3 object path to sql query.")

	flag.StringVar(&args.SQL,
		"sql",
		"",
		"Sql queries as text.")

	flag.StringVar(&args.SQLParameters,
		"sql-parameters",
		"",
		"JSON with parameters to adjust the query prior to its execution.")

	flag.Parse()

	if !((args.BucketSQL != "" && args.PathSQL != "") || args.SQL != "") {
		return &args,
			fmt.Errorf("'bucket-sql' and 'path-sql', or 'sql' command line parameters %s",
				"must be provided, type '-h' for details")
	}

	return &args, nil
}
