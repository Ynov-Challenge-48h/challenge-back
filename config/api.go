package config

import "api_test/data"

// ConfigDataApi initializes and returns a DataApi instance with predefined album data.
func ConfigDataApi() *data.DataApi {
	return &data.DataApi{
		// Predefined list of albums with their respective details.
		Albums: []data.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}
