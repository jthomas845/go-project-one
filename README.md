# Welcome to Go Project 1!

Hi! This is my first project in **Go**. The application will read a file from the filesys and write it out to the terminal; as well as keep track of how many times the application is invoked.

# Running

Run `go run main.go`

# go.mod

The go module configuration file. This project is build on Go v1.24.1

# data.txt

The invocation count is stored here, in the specified format:

    count=x
where x is simply an integer value; no extra decorators or quotes necessary.

# signal.txt

The quote to read out is here.

# main.go

The R/W function is here.
> **Helper Functions**: There are two helper functions, WriteToFile and WriteCountToFile. Why the difference? **Count needs to be stored in a specific format,** so it felt appropriate to decouple this WriteCountToFile function from WriteToFile, which writes the string more directly.

|                		|args|dependencies used                         |
|----------------|-------------------------------|-----------------------------|
|WriteToFile|(content *string*, filePath *string*)            |strings.Trim(), os.WriteFile(), fmt.Sprint()            |
|WriteCountToFile |(count *int*, filePath *string*)            |bytes.Trim(), os.WriteFile(), fmt.Sprintf()            |



# main_test.go

The strain test function is here.

## Testing

In version 1 of this application, the test is to run 50 invocations of the app and change the "signal" string twice.

In future versions of the application, the app could be containerized and the n=50 could be increased to a value that truly strains the container.

Another idea for a strain test could be to utilize a longer signal string: perhaps an entire Shakespeare play? 

## CI/CD
This repo's test fucntions are integrated with GitHub Actions for CI/CD.
Visit [the Actions Tab](https://github.com/jthomas845/go-project-one/actions) for a look at the results of the strain test.
