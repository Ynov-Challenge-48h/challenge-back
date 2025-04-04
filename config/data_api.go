package config

import (
	"api_test/data"
)

// ConfigDataApi initializes and returns a new instance of DataApi with default configuration.
//
// This function sets up the DataApi with predefined folder and database name.
//
// Returns:
//
//	*data.DataApi: A pointer to a DataApi instance configured with default folder and database name.
func ConfigDataApi() *data.DataApi {
	return &data.DataApi{
		Folder:       "./DB",        // Host on which the application will run.
		NameDatabase: "database.db", // Port on which the application will run.
	}
}
