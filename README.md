## YnabReporter
Report generation for web YNAB users

###Why?
When YNAB relaunched with their web-based product, it didn't include many features that the desktop version had.

One of the big features that was lost (for now) were Reports.

Right now, this program will print your net worth from month to month. New, more detailed reports will be added later.

###Usage
Export your data from YNAB (Go to My Budget->Export Data). This will give you a .zip file with two CSV files.

Extract the CSV files.

Build this program, and run it supplying the filenames of the CSV files as arguments (budget first, then register)

###Example
`go build`

`./YnabReporter ~/budget.csv ~/register.csv`
