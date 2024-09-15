package main

// import (
// 	"net/http"
// )

func main() {
	// Dependency Injection example
	rc := newRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc.climbRock()
	}

	// Decorator Pattern Example
	// s := &Store{}
	// http.HandleFunc("/", makeHTTPFunc(s, handler))
}

// func eg() {

// 	builder := Aero.builder()

// 	builder.inject(DBContext)
// 	builder.inject(Logger)

// 	builder.AddControllers(Controllers)

// 	app := builder.Build()

// 	app.useCors()

// 	app.serve()

// }
