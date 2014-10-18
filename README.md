# Sentimental Breakdown
---------

A Go package for measuring sentiment for a given passage of text. 
Uses the AFINN-11 

## Usage example

``` golang
	package main

	import(
		sentiment "github.com/k4orta/sentimental-breakdown"
	)

	fmt.Println(sentiment.Measure("Golly, this sure is swell!"))

	fmt.Println(sentiment.Measure("Come out to the farm and eat my ass, idiot."))
```

## Api 
	
	Measure(text string) int

	takes a string and returns a positive or negative number. 