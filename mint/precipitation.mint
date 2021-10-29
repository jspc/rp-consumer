type Precipitation {
    +mint:doc:"Timestamp the reading was taken at"
    +mint:validate:date_in_past
    +mint:transform:date_in_utc
    datetime Timestamp = 0;

    +mint:doc:"Location contains a reference to the specified"
    +mint:doc:"location of this forecast"
    Location Location = 1;

    +mint:doc:"Specific sensor used for this reading"
    uuid Sensor = 2;

    +mint:doc:"Precipitation holds the returned precipitaton level"
    +mint:doc:"in mm"
    float32 Precipitation = 3;
}
