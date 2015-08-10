# GoCap

GoCap is a golang based library and CLI for REDCap. It provides helpful abstractions for making requests to REDCap's REST API.

The project is currently in development and not considered production-ready.

## Example

#### Quickstart

```go
project := redcap.NewRedcapProject("http://redcap.example.com", "<my_api_token>", true)

fmt.Println(project.Forms)
// returns: map[demographics:0xc2082b4500 testing:0xc2082b4640 imaging:0xc2082b4780]

```
#### Creating a Postgres schema for a project and persisting forms to that DB with goroutines

```go
// initialize the project
project := redcap.NewRedcapProject("http://redcap.example.com", "<my_api_token>", true)

// Create REDCap base tables
_, err = db.Query(string(project.ToSQL("postgres")))
if err != nil {
	log.Fatal("[redcap] error creating base redcap tables. ", err)
}
```

Pull forms concurrently and persist them to the database.

```go
// Wait group for concurrent form retrieval
var form_retrieval sync.WaitGroup
form_retrieval.Add(len(project.Forms))
// Export records for project forms
for _, form := range project.Forms {
    go func(f *redcap.RedcapForm) {
        // Inform WaitGroup after each successful form retrieval
        defer form_retrieval.Done()
        // Define Export Parameters
		params := redcap.ExportParameters{
			Fields:              []string{project.Unique_key.Field_name, "redcap_event_name"},
			Forms:               []string{f.Name},
			RawOrLabel:          "label",
			Format:              "csv",
			ExportCheckboxLabel: true,
		}
        // Perform the Export
		res := project.ExportRecords(params)
        // Read the resulting bytes
		r := bytes.NewReader(res)
        // Create a CSV reader from the bytes
        reader := csv.NewReader(r)
		csvData, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

        // Open DB transaction
		txn, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		// Prepare INSERT statements from CSV file
		for i, line := range csvData {
			// skip header line
			if i == 0 {
				continue
			}
            // Clean our values
			for j, value := range line {
				// Escape quoted characters and convert to NULL if empty
				line[j] = prepareValue(value)
			}

			values := strings.Join(line, ",")
            // Create our SQL statement
			stmt, err := txn.Prepare(fmt.Sprintf("insert into %s values(%s)", f.Name, values))
			if err != nil {
				log.Fatal(fmt.Sprintf("[redcap] error formatting insert statement for table %s. ", f.Name), err)
			}
            // Execute our SQL statement
			if stmt != nil {
				_, err = stmt.Exec()
				if err != nil {
					log.Fatal("[redcap] error executing statement ", err)
				}
				err = stmt.Close()
			}
		}
        // Commit the transaction
		err = txn.Commit()
		fmt.Printf("Loaded REDCap form \"%s\"\n", f.Name)
		if err != nil {
			log.Fatal(err)
		}
	}(form)
}

form_retrieval.Wait()

```


TODO:
=====

Api functionality
-----------------
* Import Records
* Export File
* Import File
* Delete File
* Export Users
* Export Form Event Mappings


CLI
---

Tests
-----
