# CSV Statistics

Calculate average, median and person with medina from a set of files comming for a network location.

## Input

Rules applied to this version implementation:

- Local file location for this iteration
- Just one file
- csv file contains headers

## Output

Average age

## Example data

`./data` folder will contain several examples

## Usage

`go run main.go ./data/10-records.csv`

# Backlog

- [x] Read miltiple files
- [x] Read files from remote location (file://, http://, etc)
- [ ] Performance report
- [ ] Read files in parallel
- [ ] Streaming processing approach


## Hypothesis

- Read files in chunks vs process an entire file


Json
    FileLocator - Reads the json and obtains the CSV files locations
    Mapper - Reads a file and maps every recod
    Analise
    Report
Csv
    Mapper
    Analise
    Report
