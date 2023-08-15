# modern slice methods for go
  This package allows the use of 3 methods that can be used on slices of strings or numbers:
    - Filter
    - Map
    - Reduce

  Each method is implemented with the constraint.Ordered implementation to be able to accept any numeric type (including signed, unsigned, ints, and floats of all flavors) as well as strings.

  **The methods accept a pointer to a slice** to avoid copying of the slice itself.

  ## Filter
    >Filter accepts a pointer to a slice of a single type as well as a callback function that accepts a single value of that type and returns a boolean value. It will return a new slice based on values that return true in the callback

  ## Map
    >Map accepts a poitner to a slice of a single type as well as a callback funciton that accepts a single value of that type and *returns any*. If the callback doesn't return any, the code won't compile. This could cause some bugs, so be aware of it. Map will return a new slice based on the type specified in the callback function.

  ## Reduce
    >Reduce accepts a pointer to a slice of a single type as well as a callback function that accepts a single value of that type *and a starting value*. Because go does not accept default arguments, you must specify the starting value of for a reducer as a 3rd argument. Reduce will return the reducer.

Feel free to make requests to make changes!