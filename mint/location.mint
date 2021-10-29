type Location {
    +mint:doc:"Name contains the name of a location"
    +mint:validate:string_not_empty
    string Name = 0;

    +custom:validate:valid_lat
    +mint:doc:"Latitude relates the latitude of the"
    +mint:doc:"described location"
    float32 Latitude = 1;

    +mint:doc:"Longitude relates the longitude of the"
    +mint:doc:"described location"
    +custom:validate:valid_long
    float32 Longitude = 2
}
