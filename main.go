package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	// Connection to database
	dbUri := "neo4j://192.168.18.191:7687"
	dbUser := "neo4j"
	dbPassword := "ddi123"
	dbName := "kemenag"
	dbFullUri := dbUri + "/" + dbName

	driver, err := neo4j.NewDriver(
		dbFullUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""),
	)
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	err = driver.VerifyConnectivity()
	if err != nil {
		fmt.Println("Fail to connect to Neo4j database")
		return
	} else {
		fmt.Println("Successfully connected to Neo4j database")
	}

	// Menjalankan query
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead, DatabaseName: dbName})
	defer session.Close()
	query := "MATCH (n:Jemaah) RETURN n.ktp"
	result, err := session.Run(query, nil)
	if err != nil {
		panic(err)
	}

	// Mengambil hasil query
	for result.Next() {
		record := result.Record()
		ktp, found := record.Get("n.ktp")
		if !found {
			fmt.Println("KTP property not found in record")
			continue
		}

		if ktp, ok := ktp.(string); ok {
			fmt.Printf("KTP: %s\n", ktp)
		} else {
			fmt.Println("KTP property is not a string")
		}
	}

	if err := result.Err(); err != nil {
		panic(err)
	}
}
